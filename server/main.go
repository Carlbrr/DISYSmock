package main

import (
	"bufio"
	"context"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/Carlbrr/DISYSmock/proto"
	"google.golang.org/grpc"
)

const (
	MainPort = 5001
)

type Server struct {
	proto.UnimplementedChatServiceServer
	id                int
	Ports             []int32
	Port              int32
	value             int32
	ServerTimestamp   proto.LamportClock
	receivedNewLeader bool
	PrimePulse        time.Time
}

func (s *Server) Increment(ctx context.Context, incMsg *proto.IncMessage) (*proto.IncResponse, error) {
	s.ServerTimestamp.Tick()
	lastcount := s.value
	s.value += incMsg.IncCount
	return &proto.IncResponse{CurrentCount: lastcount}, nil
}

//------PASSIVE REPLICATION -----//

func (s *Server) Pulse() {
	for {
		time.Sleep(time.Second * 2)
		if s.Port == MainPort {
			s.ServerTimestamp.Tick()

			go func() {
				portIndex := s.FindIndex(s.Port)
				for i := 0; i < len(s.Ports); i++ {

					if i != int(portIndex) {
						conn, err := grpc.Dial(":"+strconv.Itoa(int(s.Ports[i])), grpc.WithInsecure())

						if err != nil {
							log.Fatalf("Failed to dial port: %v", err)
						}

						defer conn.Close()

						client := proto.NewChatServiceClient(conn)
						msg := proto.ReplicateMessage{
							TimeStamp: s.ServerTimestamp.GetTime(),
						}
						client.RegisterPulse(context.Background(), &msg)
					}
				}
			}()
		} else if s.receivedNewLeader {
			//This makes sure, that a port will not request to become leader right after it has received a new leader
			s.receivedNewLeader = false
		} else if time.Since(s.PrimePulse).Seconds() > time.Second.Seconds()*6 {

			if len(s.Ports) > 2 {
				s.ServerTimestamp.Tick()
				log.Printf("Missing pulse from prime replica. Last Pulse: %v seconds ago", time.Since(s.PrimePulse))

				log.Printf("Leader election called")

				s.CallRingElection(context.Background())

			} else {

				s.LastLeader()
			}
		}

	}
}

func (s *Server) LastLeader() { //pick the note itself as leader
	updatedPort := s.FindIndex(s.Port)
	s.Ports = removePort(s.Ports, int32(updatedPort))
	s.Port = MainPort

	grpcServer := grpc.NewServer()

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(int(s.Port)))
	if err != nil {
		log.Fatalf("Failed to listen port: %v", err)
	}

	log.Printf("Auction open at: %v", lis.Addr())

	proto.RegisterChatServiceServer(grpcServer, s)
	go func(listener net.Listener) {
		defer func() {
			lis.Close()
			log.Printf("Server stopped listening")
		}()

		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve server: %v", err)
		}
	}(lis)
}

func (s *Server) RegisterPulse(ctx context.Context, msg *proto.ReplicateMessage) (*proto.Void, error) {
	s.ServerTimestamp.SyncClocks(uint32(msg.GetTimeStamp()))
	s.PrimePulse = time.Now().UTC()
	log.Fatalf("Recieved pulse")
	return &proto.Void{}, nil
}

//actual replications
func (s *Server) Replicate(ctx context.Context, repMessage *proto.ReplicateMessage) (*proto.ReplicateResponse, error) {
	s.value = repMessage.IncSum
	s.ServerTimestamp.SyncClocks(uint32(repMessage.TimeStamp))
	return &proto.ReplicateResponse{Body: "oi we gucci"}, nil
}

func (s *Server) ReplicateBackups(ctx context.Context, currentAmount int32) {

	localWG := new(sync.WaitGroup)

	for i := 0; i < len(s.Ports); i++ {

		currentPort := s.Ports[i]

		if currentPort != s.Port {
			conn, err := grpc.Dial(":"+strconv.Itoa(int(s.Ports[i])), grpc.WithInsecure())

			if err != nil {
				log.Fatalf("Failed to dial this port(Message all): %v", err)
			}

			defer conn.Close()
			client := proto.NewChatServiceClient(conn)

			repMsg := &proto.ReplicateMessage{
				IncSum:    currentAmount,
				TimeStamp: s.ServerTimestamp.GetTime(),
			}

			var repReply *proto.ReplicateResponse
			go func() {
				var error error
				repReply, error = client.Replicate(ctx, repMsg)
				if error != nil {
					log.Printf("Could not replicate to port. %v", error)
				}
			}()
			localWG.Add(1)

			go func(port int32) {

				defer localWG.Done()
				time.Sleep(time.Second * 2)
				if repReply == nil {
					s.CallCutOffReplicate(ctx, port) //handle back-up replica failure
				}
			}(currentPort)
		}
	}
	localWG.Wait()
}

//--failure handle faults on back-up replicas--
func (s *Server) CallCutOffReplicate(ctx context.Context, brokenPort int32) {
	for i := 0; i < len(s.Ports); i++ {
		if s.Ports[i] != brokenPort {
			conn, err := grpc.Dial(":"+strconv.Itoa(int(s.Ports[i])), grpc.WithInsecure())

			if err != nil {
				log.Fatalf("Failed to dial port: %v. The server has been declared dead to remove inconsistencies", err)
			}

			defer conn.Close()
			client := proto.NewChatServiceClient(conn)

			msg := &proto.CutReplica{
				Port: brokenPort,
			}

			client.CutOfReplica(ctx, msg)
		}
	}
}
func (s *Server) CutOfReplica(ctx context.Context, msg *proto.CutReplica) (*proto.Void, error) {
	brokenPort := msg.Port
	index := s.FindIndex(brokenPort)

	s.Ports = removePort(s.Ports, int32(index))

	return &proto.Void{}, nil
}

//------LEADER ELECTION, RING BASED--------//

func (s *Server) CallRingElection(ctx context.Context) {

	listOfPorts := make([]int32, 0)
	listOfClocks := make([]uint32, 0)

	listOfPorts = append(listOfPorts, s.Port)
	listOfClocks = append(listOfClocks, s.ServerTimestamp.GetTime())

	index := s.FindIndex(s.Port)

	nextPort := s.FindNextPort(index)

	conn, err := grpc.Dial(nextPort, grpc.WithInsecure())
	if err != nil {
		log.Printf("Election: Failed to dial this port: %v", err)
	} else {
		defer conn.Close()
		client := proto.NewChatServiceClient(conn)

		protoListOfPorts := proto.PortsAndClocks{
			ListOfPorts:  listOfPorts,
			ListOfClocks: listOfClocks,
		}

		client.RingElection(ctx, &protoListOfPorts)
	}
}

func (s *Server) RingElection(ctx context.Context, msg *proto.PortsAndClocks) (*proto.Void, error) {

	listOfPorts := msg.ListOfPorts
	listOfClocks := msg.ListOfClocks

	if listOfPorts[0] == s.Port {
		var highestClock uint32
		var highestPort int32

		for i := 0; i < len(listOfPorts); i++ {
			if listOfClocks[i] > highestClock {
				highestPort = listOfPorts[i]
				highestClock = listOfClocks[i]
			} else if listOfClocks[i] == highestClock && listOfPorts[i] > highestPort { //det her else er lidt ligemeget lmao
				highestPort = listOfPorts[i]
				highestClock = listOfClocks[i]
			}
		}

		conn, err := grpc.Dial(":"+strconv.Itoa(int(highestPort)), grpc.WithInsecure())
		if err != nil {

			log.Printf("Election: Failed to dial port: %v", err)

		} else {

			defer conn.Close()
			client := proto.NewChatServiceClient(conn)

			newPortList, err := client.SelectAsNewLeader(ctx, &proto.Void{})

			if err != nil {
				log.Printf("Election: Failed to select new leader: %v", err)
			}

			for i := 0; i < len(newPortList.ListOfPorts); i++ {

				conn, err := grpc.Dial(":"+strconv.Itoa(int(newPortList.ListOfPorts[i])), grpc.WithInsecure())

				if err != nil {
					log.Printf("Election: Failed to dial this port: %v", err)
				} else {
					defer conn.Close()
					client := proto.NewChatServiceClient(conn)

					client.BroadcastNewLeader(ctx, newPortList)
				}

			}
		}
	} else {
		msg.ListOfPorts = append(msg.ListOfPorts, s.Port)

		index := s.FindIndex(s.Port)

		nextPort := s.FindNextPort(index)

		conn, err := grpc.Dial(nextPort, grpc.WithInsecure())
		if err != nil {
			log.Printf("Election: Failed to dial this port: %v", err)
		} else {
			defer conn.Close()
			client := proto.NewChatServiceClient(conn)

			protoListOfPorts := proto.PortsAndClocks{
				ListOfPorts:  listOfPorts,
				ListOfClocks: listOfClocks,
			}

			client.RingElection(ctx, &protoListOfPorts)
		}
	}
	return &proto.Void{}, nil
}

func (s *Server) BroadcastNewLeader(ctx context.Context, newPorts *proto.ElectionPorts) (*proto.Void, error) { //this is to set the boolean and update port list to port of the back-up replica that now listens on the main port
	log.Println("Received new leader")
	s.receivedNewLeader = true
	s.Ports = newPorts.ListOfPorts

	return &proto.Void{}, nil
}

func (s *Server) SelectAsNewLeader(ctx context.Context, void *proto.Void) (*proto.ElectionPorts, error) {

	oldPortIndex := s.FindIndex(s.Port)
	s.Ports = removePort(s.Ports, int32(oldPortIndex))
	s.Port = MainPort

	grpcServer := grpc.NewServer()

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(int(s.Port)))
	if err != nil {
		log.Fatalf("Failed to listen port: %v", err)
	}

	log.Printf("Auction open at: %v", lis.Addr())

	proto.RegisterChatServiceServer(grpcServer, s)
	go func(listener net.Listener) {
		defer func() {
			lis.Close()
			log.Printf("Server stopped listening")
		}()

		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve server: %v", err)
		}
	}(lis)

	newPortList := proto.ElectionPorts{
		ListOfPorts: s.Ports,
	}

	return &newPortList, nil
}

//-----MAIN-----///

func main() {
	SId := flag.Int("I", -1, "id")
	flag.Parse()

	portFile, err := os.Open("../../ports.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(portFile)
	var ports []int32

	for scanner.Scan() {
		nextPort, _ := strconv.Atoi(scanner.Text())
		ports = append(ports, int32(nextPort))
	}

	s := &Server{
		id:                *SId,
		Ports:             ports,
		Port:              ports[*SId],
		receivedNewLeader: false,
		value:             0,
		ServerTimestamp:   proto.LamportClock{},
		PrimePulse:        time.Now(),
	}
	go s.Pulse()
	//------ LOG-----////

	filename := "Server" + strconv.Itoa(s.id) + "logs.txt"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)

	//--------------////

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(int(s.Port)))
	if err != nil {
		log.Fatalf("Failed to listen on port:, %v", err)
	} else {
		log.Printf("Listening on port :" + strconv.Itoa(int(s.Port)))
	}

	grpcServer := grpc.NewServer()

	proto.RegisterChatServiceServer(grpcServer, s)

	defer func() {
		lis.Close()
		log.Printf("Server stopped listening")
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server, %v", err)
	}

}

////------ HELPER METHODS ------- ////

//helper method, to not find index of own port, as the list of ports will change

func (s *Server) FindIndex(port int32) int {
	index := -1

	for i := 0; i < len(s.Ports); i++ {
		if s.Ports[i] == port {
			index = i
			break
		}
	}

	return index
}
func removePort(s []int32, i int32) []int32 {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
func (s *Server) FindNextPort(index int) string {

	nextPort := s.Ports[(index+1)%len(s.Ports)]
	if nextPort == MainPort {
		nextPort = s.Ports[(index+2)%len(s.Ports)]
	}

	return ":" + strconv.Itoa(int(nextPort))
}

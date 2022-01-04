package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/Carlbrr/DISYSmock/proto"

	"google.golang.org/grpc"
)

const (
	MainPort = 5001
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":"+strconv.Itoa(int(MainPort)), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	done := make(chan int)

	c := proto.NewChatServiceClient(conn)
	log.Println("connected to port :", MainPort)
	wait := sync.WaitGroup{}
	//the loop for making requests from terminal

	wait.Add(1)

	go func() {

		defer wait.Done()

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			inputArray := strings.Fields(scanner.Text())

			input := strings.ToLower(strings.TrimSpace(inputArray[0]))

			if input == "exit" {
				os.Exit(1)
			}
			if input != "" {
				inc, err := strconv.Atoi(input)
				if err == nil {

					msg := proto.IncMessage{
						IncCount: int32(inc),
					}

					response, err := c.Increment(context.Background(), &msg)

					if err != nil {
						log.Printf("Error callin increment: %v", err)
						os.Exit(1)
					}

					log.Printf("Response from server: %d", response.CurrentCount)
				}
			}
		}
	}()
	go func() {
		wait.Wait()
		close(done)
	}()

	<-done
}

package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpcgps/proto"
	"log"
	"os"
	"strings"
	"time"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMyGpsClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please enter the Address to Get current GPS location")
		fmt.Println("----------------------------------------------------")
		inputString, _ := reader.ReadString('\n')
		// convert CRLF to LF
		inputString = strings.Replace(inputString, "\n", "", -1)
		if inputString == "q" {
			break
		}
		r, err := c.GetMyAddress(ctx, &pb.Address{Addr: inputString})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		location := r.GetLocation()

		log.Printf("Address Location: Latitude: %f, Longitude: %f\n", location.GetP1(), location.GetP2())
	}
}

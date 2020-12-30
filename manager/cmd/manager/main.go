package main

import (
	"context"
	"google.golang.org/grpc"
	"imager/manager/internal/client"
	"log"
)

func main() {
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	err = client.New(conn).Process(context.Background())
	if err != nil {
		log.Fatalf("did not get result: %s", err)
	}
}
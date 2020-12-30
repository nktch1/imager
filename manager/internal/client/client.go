package client

import (
	"context"
	"crypto/md5"
	"fmt"
	"google.golang.org/grpc"
	"imager/manager/api"
	"io"
	"log"
	"strconv"
)

type client struct {
	conn *grpc.ClientConn
}

func New(conn *grpc.ClientConn) *client {
	return &client{conn: conn}
}

func (c *client) Process(ctx context.Context) error  {
	client := api.NewProcessorClient(c.conn)
	stream, err := client.Process(ctx)
	if err != nil {
		log.Fatal("can't get a stream", err)
	}

	waitResponse := make(chan error)

	go func() {
		for {
			humanPosition, err := stream.Recv()

			if err == io.EOF {
				log.Fatal("end of corrupted")
				waitResponse <- nil
			}

			if err != nil {
				log.Fatal("stream is corrupted")
				waitResponse <- err
			}

			fmt.Println(humanPosition)
		}
	}()

	go func() {
		var counter int
		for {
			id := strconv.Itoa(counter)
			if err := stream.Send(&api.RawImage{
				Id:   id,
				Image: []byte(fmt.Sprintf("id=%s;data=%s", id, md5.Sum([]byte(id))) ),
			}); err != nil {
				log.Fatal("stream is corrupted; can't send data")
				waitResponse <- err
			}
			counter++
		}
	}()

	return <-waitResponse
}
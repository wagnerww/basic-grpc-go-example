package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/codeedu/fc2-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	//AddUser(client)
	//AddUserVerbose(client)
	//AddUsers(client)
	AddUserStramBoth(client)

}

func AddUser(client pb.UserServiceClient) {

	req := &pb.User{
		Id:    "0",
		Name:  "João",
		Email: "j@j.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)

}

// Função baseada em Stream, ou seja, fica recebendo coisas enquanto o server
/// processa
func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "João",
		Email: "j@j.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not receive the msg: %v", err)
		}

		fmt.Println("Status:", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "w1",
			Name:  "wagner 1",
			Email: "wagner@gmail.com",
		},
		&pb.User{
			Id:    "w2",
			Name:  "wagner 2",
			Email: "wagner@gmail.com",
		},
		&pb.User{
			Id:    "w3",
			Name:  "wagner 3",
			Email: "wagner@gmail.com",
		},
		&pb.User{
			Id:    "w4",
			Name:  "wagner 4",
			Email: "wagner@gmail.com",
		},
		&pb.User{
			Id:    "w5",
			Name:  "wagner 5",
			Email: "wagner@gmail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}

// Implementação bi-direcional
func AddUserStramBoth(client pb.UserServiceClient) {
	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	reqs := []*pb.User{
		&pb.User{
			Id:    "w1",
			Name:  "wagner 1",
			Email: "wagner@gmail.com",
		},
		&pb.User{
			Id:    "w2",
			Name:  "wagner 2",
			Email: "wagner@gmail.com",
		},
		&pb.User{
			Id:    "w3",
			Name:  "wagner 3",
			Email: "wagner@gmail.com",
		},
		&pb.User{
			Id:    "w4",
			Name:  "wagner 4",
			Email: "wagner@gmail.com",
		},
		&pb.User{
			Id:    "w5",
			Name:  "wagner 5",
			Email: "wagner@gmail.com",
		},
	}

	//aPENAS SEGURA O PROGRAMA
	wait := make(chan int)

	// Função assincrona, thread aberta e controlada pelo GO
	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	// Função assincrona que fica em loop escutando e dando o status
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error receiving data: %v", err)
				break
			}

			fmt.Printf("Recebendo user %v com status: %v\n", res.GetUser().GetName(), res.GetStatus())
		}

		// Mata o programa
		close(wait)
	}()

	<-wait
}

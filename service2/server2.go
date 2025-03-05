package main

import (
	"context"
	pb "grpc/server/addressbook"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("couldnot connect to service")
	}
	defer conn.Close()
	client := pb.NewAddressBookServiceClient(conn)
	person := &pb.Person{
		Name: "yehia",
		Id:   1,
	}
	if _, err = client.AddPerson(context.Background(), person); err != nil {
		log.Fatal("couldnot add person")
	}
}

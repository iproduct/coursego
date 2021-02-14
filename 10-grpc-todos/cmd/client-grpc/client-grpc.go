package main

import (
	"context"
	"flag"
	"github.com/golang/protobuf/ptypes"
	"github.com/iproduct/coursego/10-grpc-todos/generated/todo_service"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

func main() {
	// get configuration
	address := flag.String("server", ":9000", "gRPC server in format host:port")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := todo_service.NewToDoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t := time.Now().In(time.UTC)
	reminder, _ := ptypes.TimestampProto(t)
	pfx := t.Format(time.RFC3339Nano)

	// Call Create
	req1 := todo_service.CreateRequest{
		Api: apiVersion,
		ToDo: &todo_service.ToDo{
			Title:       "title (" + pfx + ")",
			Description: "description (" + pfx + ")",
			Reminder:    reminder,
		},
	}
	res1, err := c.Create(ctx, &req1)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res1)

	// Call Create 2
	req2 := todo_service.CreateRequest{
		Api: apiVersion,
		ToDo: &todo_service.ToDo{
			Title:       "title (" + pfx + ")",
			Description: "description (" + pfx + ")",
			Reminder:    reminder,
		},
	}
	res2, err := c.Create(ctx, &req2)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res2)

	id := res1.Id

	// Read
	req3 := todo_service.ReadRequest{
		Api: apiVersion,
		Id:  id,
	}
	res3, err := c.Read(ctx, &req3)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	log.Printf("Read result: <%+v>\n\n", res3)

	// Update
	req4 := todo_service.UpdateRequest{
		Api: apiVersion,
		ToDo: &todo_service.ToDo{
			Id:          res3.ToDo.Id,
			Title:       res3.ToDo.Title,
			Description: res3.ToDo.Description + " + updated",
			Reminder:    res3.ToDo.Reminder,
		},
	}
	res4, err := c.Update(ctx, &req4)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}
	log.Printf("Update result: <%+v>\n\n", res4)

	// Call ReadAll
	req5 := todo_service.ReadAllRequest{
		Api: apiVersion,
	}
	res5, err := c.ReadAll(ctx, &req5)
	if err != nil {
		log.Fatalf("ReadAll failed: %v", err)
	}
	log.Printf("ReadAll result: <%+v>\n\n", res5)

	// Delete
	req6 := todo_service.DeleteRequest{
		Api: apiVersion,
		Id:  id,
	}
	res6, err := c.Delete(ctx, &req6)
	if err != nil {
		log.Fatalf("Delete failed: %v", err)
	}
	log.Printf("Delete result: <%+v>\n\n", res6)

	// Call ReadAll
	req7 := todo_service.ReadAllRequest{
		Api: apiVersion,
	}
	res7, err := c.ReadAll(ctx, &req7)
	if err != nil {
		log.Fatalf("ReadAll failed: %v", err)
	}
	log.Printf("ReadAll result: <%+v>\n\n", res7)


}
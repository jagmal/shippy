//shippy-cli-consignment/main.go
package cli

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"context"

	pb "github.com/jagmal/shippy-service-consignment/proto/consignment"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
	defaultFileName = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment, *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	// Contact the server and print out its response.
	file := defaultFileName
	if len(os.args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)
}

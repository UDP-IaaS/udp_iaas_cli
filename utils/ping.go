package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "udp_iaas_cli/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGrpcConnection() error {
    conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect: %v", err)
	}
	defer conn.Close()
	
    c := pb.NewUtilsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Ping(ctx, &pb.PingRequest{Message: "ping"})
	if err != nil {
		return fmt.Errorf("could not ping: %v", err)
	}
	
	log.Printf("Ping Response: %s (timestamp: %v)", r.GetMessage(), time.Unix(r.GetTimestamp(), 0))
	return nil
}

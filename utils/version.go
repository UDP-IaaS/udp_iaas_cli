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

func Version() error {
    conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect: %v", err)
	}
	defer conn.Close()
    
	c := pb.NewUtilsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetVersion(ctx, &pb.VersionRequest{})
	if err != nil {
		return fmt.Errorf("could not get daemon's version: %v", err)
	}
	
	log.Printf("Daemon version: %s", r.GetVersion())
	return nil
}

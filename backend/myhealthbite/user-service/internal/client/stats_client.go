package client

import (
	"log"
	"time"

	"MyHealthBite/stats-service/proto"

	"google.golang.org/grpc"
)

var StatsClient proto.StatsServiceClient

func InitStatsClient() {
	conn, err := grpc.Dial("stats-service:50054", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		log.Fatalf("Could not connect to stats-service: %v", err)
	}

	StatsClient = proto.NewStatsServiceClient(conn)
}

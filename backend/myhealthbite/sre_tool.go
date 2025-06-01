package main

import (
	"fmt"
	"net"
	"time"
)

type Service struct {
	Name string
	Host string
	Port string
}

var services = []Service{
	{"user-service", "localhost", "50051"},
	{"stats-service", "localhost", "50054"},
	{"email-service", "localhost", "50058"},
}

func checkService(s Service) bool {
	address := net.JoinHostPort(s.Host, s.Port)
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func main() {
	fmt.Println("🔍 Health Check Report (SRE Tool):")
	for _, s := range services {
		status := "❌ DOWN"
		if checkService(s) {
			status = "✅ UP"
		}
		fmt.Printf("- %s (%s:%s): %s\n", s.Name, s.Host, s.Port, status)
	}
}

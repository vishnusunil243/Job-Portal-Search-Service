package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/vishnusunil243/Job-Portal-Search-Service/db"
	"github.com/vishnusunil243/Job-Portal-Search-Service/initializer"
	"github.com/vishnusunil243/Job-Portal-proto-files/pb"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf(err.Error())
	}
	addr := os.Getenv("DB_KEY")
	DB, err := db.InitDB(addr)
	if err != nil {
		log.Fatal("error connecting to database")

	}
	listener, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal("failed to listen on port 8083")
	}
	fmt.Println("email service listening on port 8083")
	services := initializer.Initializer(DB)
	server := grpc.NewServer()
	pb.RegisterSearchServiceServer(server, services)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to listen on port 8083")
	}
}
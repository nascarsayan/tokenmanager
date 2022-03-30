package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc.io/tokenman/src/pb"
)

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {

	var port string
	var hostname string
	var tokenID string
	var tokenName string

	var low uint64
	var mid uint64
	var high uint64

	var isCreate bool
	var isWrite bool
	var isRead bool
	var isDrop bool

	flag.StringVar(&hostname, "hostname", "", "Hostname of the server")
	flag.StringVar(&port, "port", "", "Port of the server")
	flag.StringVar(&tokenID, "id", "", "ID of the token")

	flag.BoolVar(&isCreate, "create", false, "Create token")
	flag.BoolVar(&isWrite, "write", false, "Write token")
	flag.BoolVar(&isRead, "read", false, "Read token")
	flag.BoolVar(&isDrop, "drop", false, "Drop token")

	flag.StringVar(&tokenName, "name", "", "The name of the token.")
	flag.Uint64Var(&low, "low", 0, "The value low.")
	flag.Uint64Var(&mid, "mid", 0, "The value low.")
	flag.Uint64Var(&high, "high", 0, "The value low.")
	flag.Parse()

	if hostname == "" {
		log.Fatalf("hostname should be provided")
	}
	if port == "" {
		log.Fatalf("port should be provided")
	}

	var actionCount uint
	if isCreate {
		actionCount++
	}
	if isWrite {
		actionCount++
	}
	if isRead {
		actionCount++
	}
	if isDrop {
		actionCount++
	}
	if actionCount != 1 {
		log.Fatalf("exactly one action has to be specified in [create, write, read, drop]")
	}

	if isCreate {
		createToken(hostname, port, tokenID)
	}

	if isWrite {
		flags := []string{"name", "low", "mid", "high"}
		for _, f := range flags {
			if !(isFlagPassed(f)) {
				log.Fatalf("%s was not provided", f)
			}
		}
		writeToken(hostname, port, tokenID, tokenName, low, mid, high)
	}

	if isRead {
		readToken(hostname, port, tokenID)
	}

	if isDrop {
		dropToken(hostname, port, tokenID)
	}
}

func createToken(hostname string, port string, id string) {
	log.Println("Creating Token")
	var conn *grpc.ClientConn
	address := fmt.Sprintf("%s:%s", hostname, port)
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewTokenClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := pb.CreateReq{
		Id: id,
	}
	r, err := client.Create(ctx, &req)
	if err != nil {
		log.Fatalf("could not create token: %v", err)
	}
	fmt.Println(r.String())
}

func readToken(hostname string, port string, id string) {
	log.Println("Reading Token")
	var conn *grpc.ClientConn
	address := fmt.Sprintf("%s:%s", hostname, port)
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewTokenClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := pb.ReadReq{
		Id: id,
	}
	r, err := client.Read(ctx, &req)
	if err != nil {
		log.Fatalf("could not read token: %v", err)
	}
	fmt.Println(r.String())
}

func dropToken(hostname string, port string, id string) {
	log.Println("Dropping Token")
	var conn *grpc.ClientConn
	address := fmt.Sprintf("%s:%s", hostname, port)
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewTokenClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := pb.DropReq{
		Id: id,
	}
	r, err := client.Drop(ctx, &req)
	if err != nil {
		log.Fatalf("could not drop token: %v", err)
	}
	fmt.Println(r.String())
}

func writeToken(
	hostname string,
	port string,
	id string,
	name string,
	low uint64,
	mid uint64,
	high uint64) {
	log.Println("Writing Token")
	var conn *grpc.ClientConn
	address := fmt.Sprintf("%s:%s", hostname, port)
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewTokenClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := pb.WriteReq{
		Id:   id,
		Name: name,
		Low:  low,
		High: high,
		Mid:  mid,
	}
	r, err := client.Write(ctx, &req)
	if err != nil {
		log.Fatalf("could not write token: %v", err)
	}
	fmt.Println(r.String())
}

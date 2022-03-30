package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"log"

	"grpc.io/tokenman/src/pb"
)

type Token struct {
	Id         string
	Name       string
	Low        uint64
	Mid        uint64
	High       uint64
	PartialVal *uint64
	FinalVal   *uint64
}

var tokenList map[string]Token
var tokenMu sync.Mutex

func init() {
	tokenList = make(map[string]Token)
}

func printTokenInfo(id string) {
	token, err := json.Marshal(tokenList[id])
	if err != nil {
		log.Println("Could not fetch token details")
	}
	log.Printf("Token content: %+v\n", string(token))

	tokenIds := make([]string, 0, len(tokenList))
	for k := range tokenList {
		tokenIds = append(tokenIds, k)
	}
	log.Printf("Token ids present: %+v", tokenIds)
}

type Server struct {
	pb.UnimplementedTokenServer
}

func (s *Server) Create(
	ctx context.Context, req *pb.CreateReq) (*pb.CreateRes, error) {

	var res pb.CreateRes
	if req == nil {
		msg := "request must not be nil"
		log.Println(msg)
		return &res, fmt.Errorf(msg)
	}

	if req.Id == "" {
		msg := "token id must not be empty"
		log.Println(msg)
		return &res, fmt.Errorf(msg)
	}

	log.Println("Create token")

	tokenMu.Lock()
	defer tokenMu.Unlock()

	tokenList[req.Id] = Token{
		Id: req.Id,
	}

	printTokenInfo(req.Id)
	return &res, nil
}

func (s *Server) Drop(
	ctx context.Context, req *pb.DropReq) (*pb.DropRes, error) {

	var res pb.DropRes
	if req == nil {
		msg := "request must not be nil"
		log.Println(msg)
		return &res, fmt.Errorf(msg)
	}

	if req.Id == "" {
		msg := "token id must not be empty"
		log.Println(msg)
		return &res, fmt.Errorf(msg)
	}

	log.Println("Deleting token")

	tokenMu.Lock()
	defer tokenMu.Unlock()

	delete(tokenList, req.Id)

	printTokenInfo(req.Id)
	return &res, nil
}

func (s *Server) Write(
	ctx context.Context, req *pb.WriteReq) (*pb.WriteRes, error) {

	var res pb.WriteRes
	if req == nil {
		msg := "request must not be nil"
		log.Println(msg)
		return &res, fmt.Errorf(msg)
	}

	if req.Id == "" {
		msg := "token id must not be empty"
		log.Println(msg)
		return &res, fmt.Errorf(msg)
	}

	tokenMu.Lock()
	defer tokenMu.Unlock()

	if tokenList[req.Id].Id == "" {
		msg := "token id not found in server"
		log.Println(msg)
		return &res, fmt.Errorf(msg)
	}

	log.Println("Writing token")

	partial := getArgMin(req.Name, req.Low, req.Mid)

	tokenList[req.Id] = Token{
		Id:         req.Id,
		Name:       req.Name,
		Low:        req.Low,
		High:       req.High,
		Mid:        req.Mid,
		PartialVal: &partial,
		FinalVal:   nil,
	}
	res.PartialVal = partial

	printTokenInfo(req.Id)
	return &res, nil
}

func (s *Server) Read(
	ctx context.Context, req *pb.ReadReq) (*pb.ReadRes, error) {

	var res pb.ReadRes
	if req == nil {
		msg := "request must not be nil"
		log.Println(msg)
		return &res, fmt.Errorf(msg)
	}

	if req.Id == "" {
		msg := "token id must not be empty"
		log.Println(msg)
		return &res, fmt.Errorf(msg)
	}

	tokenMu.Lock()
	defer tokenMu.Unlock()

	if tokenList[req.Id].Id == "" {
		msg := "token id not found in server"
		log.Println(msg)
		return &res, fmt.Errorf(msg)
	}

	if tokenList[req.Id].PartialVal == nil {
		msg := "partial value for token not found"
		log.Println(msg)
		return &res, fmt.Errorf(msg)
	}

	log.Println("Reading token")

	final := getArgMin(tokenList[req.Id].Name, tokenList[req.Id].Mid, tokenList[req.Id].High)

	token := tokenList[req.Id]
	token.FinalVal = &final
	tokenList[req.Id] = token
	res.FinalVal = final

	printTokenInfo(req.Id)
	return &res, nil
}

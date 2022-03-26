package server

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	pb "example.com/tokenmanager/pkg/token"
	"github.com/sirupsen/logrus"
)

type Token struct {
	Id      string
	Name    string
	Low     uint64
	Mid     uint64
	High    uint64
	Partial *uint64
	Final   *uint64
}

var tokens map[string]Token
var tokenMu sync.Mutex

func logTokenInfo(log *logrus.Entry, id string) {
	token, err := json.Marshal(tokens[id])
	if err != nil {
		log.Info("Could not fetch token details")
	}
	log.Infof("Token content: %+v", string(token))

	tokenIds := make([]string, 0, len(tokens))
	for k := range tokens {
		tokenIds = append(tokenIds, k)
	}
	log.Infof("Token ids present: %+v", tokenIds)
}

type Server struct {
	pb.UnimplementedTokenServer
}

func init() {
	tokens = make(map[string]Token)
}

func (s *Server) CreateToken(
	ctx context.Context, req *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error) {

	var res pb.CreateTokenResponse
	if req == nil {
		msg := "request must not be nil"
		logrus.Info(msg)
		return &res, fmt.Errorf(msg)
	}

	log := logrus.WithFields(logrus.Fields{
		"id": req.Id,
	})

	if req.Id == "" {
		msg := "token id must not be empty"
		log.Info(msg)
		return &res, fmt.Errorf(msg)
	}

	log.Info("Creating token")

	tokenMu.Lock()
	defer tokenMu.Unlock()

	tokens[req.Id] = Token{
		Id: req.Id,
	}

	logTokenInfo(log, req.Id)
	return &res, nil
}

func (s *Server) DropToken(
	ctx context.Context, req *pb.DropTokenRequest) (*pb.DropTokenResponse, error) {

	var res pb.DropTokenResponse
	if req == nil {
		msg := "request must not be nil"
		logrus.Info(msg)
		return &res, fmt.Errorf(msg)
	}

	log := logrus.WithFields(logrus.Fields{
		"id": req.Id,
	})

	if req.Id == "" {
		msg := "token id must not be empty"
		log.Info(msg)
		return &res, fmt.Errorf(msg)
	}

	log.Info("Deleting token")

	tokenMu.Lock()
	defer tokenMu.Unlock()

	delete(tokens, req.Id)

	logTokenInfo(log, req.Id)
	return &res, nil
}

func (s *Server) WriteToken(
	ctx context.Context, req *pb.WriteTokenRequest) (*pb.WriteTokenResponse, error) {

	var res pb.WriteTokenResponse
	if req == nil {
		msg := "request must not be nil"
		logrus.Info(msg)
		return &res, fmt.Errorf(msg)
	}

	log := logrus.WithFields(logrus.Fields{
		"id": req.Id,
		// "name": req.Name,
		// "low":  req.Low,
		// "mid":  req.Mid,
		// "high": req.High,
	})

	if req.Id == "" {
		msg := "token id must not be empty"
		log.Info(msg)
		return &res, fmt.Errorf(msg)
	}

	if !((req.Low < req.Mid) && (req.Mid < req.High)) {
		msg := fmt.Sprintf(
			"The following condition is not met: low < mid < high: %d < %d < %d",
			req.Low, req.Mid, req.High)
		log.Info(msg)
		return &res, fmt.Errorf(msg)
	}

	tokenMu.Lock()
	defer tokenMu.Unlock()

	if tokens[req.Id].Id == "" {
		msg := "token id not found in server"
		log.Info(msg)
		return &res, fmt.Errorf(msg)
	}

	log.Info("Writing token")

	partial := computeArgMinHash(req.Name, req.Low, req.Mid)

	tokens[req.Id] = Token{
		Id:      req.Id,
		Name:    req.Name,
		Low:     req.Low,
		High:    req.High,
		Mid:     req.Mid,
		Partial: &partial,
		Final:   nil,
	}
	res.Partial = partial

	logTokenInfo(log, req.Id)
	return &res, nil
}

func (s *Server) ReadToken(
	ctx context.Context, req *pb.ReadTokenRequest) (*pb.ReadTokenResponse, error) {

	var res pb.ReadTokenResponse
	if req == nil {
		msg := "request must not be nil"
		logrus.Info(msg)
		return &res, fmt.Errorf(msg)
	}

	log := logrus.WithFields(logrus.Fields{
		"id": req.Id,
	})

	if req.Id == "" {
		msg := "token id must not be empty"
		log.Info(msg)
		return &res, fmt.Errorf(msg)
	}

	tokenMu.Lock()
	defer tokenMu.Unlock()

	if tokens[req.Id].Id == "" {
		msg := "token id not found in server"
		log.Info(msg)
		return &res, fmt.Errorf(msg)
	}

	if tokens[req.Id].Partial == nil {
		msg := "cannot read token, it has not been written yet"
		log.Info(msg)
		return &res, fmt.Errorf(msg)
	}

	log.Info("Reading token")

	final := computeArgMinHash(tokens[req.Id].Name, tokens[req.Id].Mid, tokens[req.Id].High)

	token := tokens[req.Id]
	token.Final = &final
	tokens[req.Id] = token
	res.Final = final

	logTokenInfo(log, req.Id)
	return &res, nil
}

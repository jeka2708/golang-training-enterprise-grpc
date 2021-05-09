package api

import (
	"context"
	"fmt"
	"github.com/jeka2708/golang-training-enterprise-grpc/pkg/data"
	pb "github.com/jeka2708/golang-training-enterprise-grpc/proto/go_proto"
	log "github.com/sirupsen/logrus"
)

type ClientServer struct {
	data *data.ClientData
}

func (c ClientServer) ReadAllClient(ctx context.Context, request *pb.ListClientRequest) (*pb.ListClientResponse, error) {
	cc, err := c.data.ReadAllClients()
	if err != nil {
		log.Println(err)
	}

	var list []*pb.DataClient
	for _, t := range cc {

		list = append(list, structClientToRes(t))

	}

	return &pb.ListClientResponse{Data: list}, nil
}

func (c ClientServer) CreateClient(ctx context.Context, client *pb.DataClient) (*pb.IdClient, error) {
	var entity = data.Client{
		FirstNameC:   client.FirstNameC,
		LastNameC:    client.LastNameC,
		MiddleNameC:  client.MiddleNameC,
		PhoneNumberC: client.PhoneNumberC,
	}
	id, err := c.data.AddClient(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"client": entity,
		}).Warningf("got an error when tried to create client: %s", err)
		return &pb.IdClient{Id: -1}, fmt.Errorf("got an error when tried to create client: %w", err)
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"client": entity,
	}).Info("create client")
	return &pb.IdClient{Id: int64(id)}, nil
}

func (c ClientServer) DeleteClient(ctx context.Context, client *pb.IdClient) (*pb.StatusClientResponse, error) {
	var entity = new(data.Division)
	entity.Id = int(client.Id)
	err := c.data.DeleteByIdClient(entity.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"client": entity,
		}).Warningf("got an error when tried to delete client: %s", err)
		return &pb.StatusClientResponse{Message: "got an error when tried to delete client"},
			fmt.Errorf("got an error when tried to delete division: %w", err)
	}
	log.WithFields(log.Fields{
		"client": entity,
	}).Info("client was delete")
	return &pb.StatusClientResponse{Message: "client was delete"}, nil
}

func (c ClientServer) UpdateClient(ctx context.Context, client *pb.DataClient) (*pb.StatusClientResponse, error) {
	var entity = data.Client{
		FirstNameC:   client.FirstNameC,
		LastNameC:    client.LastNameC,
		MiddleNameC:  client.MiddleNameC,
		PhoneNumberC: client.PhoneNumberC,
	}
	err := c.data.UpdateClient(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"client": entity,
		}).Warningf("got an error when tried to update client: %s", err)
		return &pb.StatusClientResponse{Message: "got an error when tried to delete client"},
			fmt.Errorf("got an error when tried to update division: %w", err)
	}
	log.WithFields(log.Fields{
		"client": entity,
	}).Info("client was update")
	return &pb.StatusClientResponse{Message: "client was update"}, nil
}

func NewClientServer(d data.ClientData) *ClientServer {
	return &ClientServer{data: &d}
}
func structClientToRes(data data.Client) *pb.DataClient {

	id := data.Id

	d := &pb.DataClient{
		FirstNameC:   data.FirstNameC,
		LastNameC:    data.LastNameC,
		MiddleNameC:  data.MiddleNameC,
		PhoneNumberC: data.PhoneNumberC,
	}

	if id != 0 {
		d.Id = int64(id)
	}

	return d
}

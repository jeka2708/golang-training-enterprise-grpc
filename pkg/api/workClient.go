package api

import (
	"context"
	"fmt"
	"github.com/jeka2708/golang-training-enterprise-grpc/pkg/data"
	pb "github.com/jeka2708/golang-training-enterprise-grpc/proto/go_proto"
	log "github.com/sirupsen/logrus"
)

type WorkClientServer struct {
	data *data.WorkClientData
}

func NewWorkClientServer(d data.WorkClientData) *WorkClientServer {
	return &WorkClientServer{data: &d}
}

func (w WorkClientServer) ReadAllClient(ctx context.Context, request *pb.ListWorkClientRequest) (*pb.ListWorkClientResponse, error) {
	ww, err := w.data.ReadAllWorkClients()
	if err != nil {
		log.Println(err)
	}

	var list []*pb.DataWorkClient
	for _, t := range ww {

		list = append(list, structWorkClientToRes(t))

	}

	return &pb.ListWorkClientResponse{Data: list}, nil
}

func (w WorkClientServer) CreateWorkClient(ctx context.Context, client *pb.DataWorkClient) (*pb.IdWorkClient, error) {
	var entity = data.WorkClient{
		WorkId:   int(client.WorkId),
		ClientId: int(client.ClientId),
	}
	id, err := w.data.AddWorkClient(entity.WorkId, entity.ClientId)
	if err != nil {
		log.WithFields(log.Fields{
			"workClient": entity,
		}).Warningf("got an error when tried to create workClient: %s", err)
		return &pb.IdWorkClient{Id: -1}, fmt.Errorf("got an error when tried to create workClient: %w", err)
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"workClient": entity,
	}).Info("create workClient")
	return &pb.IdWorkClient{Id: int64(id)}, nil
}

func (w WorkClientServer) DeleteWorkClient(ctx context.Context, client *pb.IdWorkClient) (*pb.StatusWorkClientResponse, error) {
	var entity = new(data.Division)
	entity.Id = int(client.Id)
	err := w.data.DeleteByIdWorkClient(entity.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"workClient": entity,
		}).Warningf("got an error when tried to delete workClient: %s", err)
		return &pb.StatusWorkClientResponse{Message: "got an error when tried to delete workClient"},
			fmt.Errorf("got an error when tried to delete workClient: %w", err)
	}
	log.WithFields(log.Fields{
		"workClient": entity,
	}).Info("workClient was delete")
	return &pb.StatusWorkClientResponse{Message: "workClient was delete"}, nil
}

func (w WorkClientServer) UpdateWorkClient(ctx context.Context, client *pb.DataWorkClient) (*pb.StatusWorkClientResponse, error) {
	var entity = data.WorkClient{
		WorkId:   int(client.WorkId),
		ClientId: int(client.ClientId),
	}
	err := w.data.UpdateWorkClient(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"workClient": entity,
		}).Warningf("got an error when tried to update workClient: %s", err)
		return &pb.StatusWorkClientResponse{Message: "got an error when tried to delete workClient"},
			fmt.Errorf("got an error when tried to update work: %w", err)
	}
	log.WithFields(log.Fields{
		"workClient": entity,
	}).Info("workClient was update")
	return &pb.StatusWorkClientResponse{Message: "workClient was update"}, nil
}

func structWorkClientToRes(data data.ResultClientWork) *pb.DataWorkClient {

	id := data.Id

	d := &pb.DataWorkClient{
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		MiddleName:   data.MiddleName,
		PhoneNumber:  data.PhoneNumber,
		FirstNameC:   data.FirstNameC,
		LastNameC:    data.LastNameC,
		MiddleNameC:  data.MiddleNameC,
		PhoneNumberC: data.PhoneNumberC,
		Name:         data.Name,
		Cost:         data.Cost,
	}

	if id != 0 {
		d.Id = int64(id)
	}

	return d
}

package api

import (
	"context"
	"github.com/jeka2708/golang-training-enterprise-grpc/pkg/data"
	pb "github.com/jeka2708/golang-training-enterprise-grpc/proto/go_proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	if err := checkWorkClientRequest(client); err != nil {
		log.WithFields(log.Fields{
			"workClient": client,
		}).Warningf("empty fields error: %s", err)
		return &pb.IdWorkClient{Id: -1}, err
	}
	var entity = data.WorkClient{
		WorkId:   int(client.WorkId),
		ClientId: int(client.ClientId),
	}
	id, err := w.data.AddWorkClient(entity.WorkId, entity.ClientId)
	if err != nil {
		log.WithFields(log.Fields{
			"workClient": entity,
		}).Warningf("got an error when tried to create workClient: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to create workClient: %s, with error: %w", client, err)
		errWithDetails, err := s.WithDetails(client)
		if err != nil {
			return &pb.IdWorkClient{Id: -1}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.IdWorkClient{Id: -1}, errWithDetails.Err()

	}
	entity.Id = id
	log.WithFields(log.Fields{
		"workClient": entity,
	}).Info("create workClient")
	return &pb.IdWorkClient{Id: int64(id)}, nil
}

func (w WorkClientServer) DeleteWorkClient(ctx context.Context, client *pb.IdWorkClient) (*pb.StatusWorkClientResponse, error) {
	if err := checkId(client.GetId()); err != nil {
		log.WithFields(log.Fields{
			"workClient": client,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusWorkClientResponse{Message: "empty fields error"}, err
	}
	var entity = new(data.Division)
	entity.Id = int(client.Id)
	err := w.data.DeleteByIdWorkClient(entity.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"workClient": entity,
		}).Warningf("got an error when tried to delete workClient: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to delete client: %s, with error: %w", client, err)
		errWithDetails, err := s.WithDetails(client)
		if err != nil {
			return &pb.StatusWorkClientResponse{Message: "got an error when tried to delete workClient"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusWorkClientResponse{Message: "got an error when tried to delete workClient"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"workClient": entity,
	}).Info("workClient was delete")
	return &pb.StatusWorkClientResponse{Message: "workClient was delete"}, nil
}

func (w WorkClientServer) UpdateWorkClient(ctx context.Context, client *pb.DataWorkClient) (*pb.StatusWorkClientResponse, error) {
	if err := checkId(client.GetId()); err != nil {
		log.WithFields(log.Fields{
			"workClient": client,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusWorkClientResponse{Message: "empty fields error"}, err
	}
	if err := checkWorkClientRequest(client); err != nil {
		log.WithFields(log.Fields{
			"workClient": client,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusWorkClientResponse{Message: "empty fields error"}, err
	}
	var entity = data.WorkClient{
		WorkId:   int(client.WorkId),
		ClientId: int(client.ClientId),
	}
	err := w.data.UpdateWorkClient(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"workClient": entity,
		}).Warningf("got an error when tried to update workClient: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to update client: %s, with error: %w", client, err)
		errWithDetails, err := s.WithDetails(client)
		if err != nil {
			return &pb.StatusWorkClientResponse{Message: "got an error when tried to update workClient"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusWorkClientResponse{Message: "got an error when tried to update workClient"}, errWithDetails.Err()
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
func checkWorkClientRequest(in *pb.DataWorkClient) error {
	if in.GetWorkId() == 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {WorkId}: %s", in.GetName())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetClientId() == 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {ClientId}: %s", in.GetName())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	return nil
}

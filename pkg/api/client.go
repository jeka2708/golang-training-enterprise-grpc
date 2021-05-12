package api

import (
	"context"
	"github.com/jeka2708/golang-training-enterprise-grpc/pkg/data"
	pb "github.com/jeka2708/golang-training-enterprise-grpc/proto/go_proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	if err := checkClientRequest(client); err != nil {
		log.WithFields(log.Fields{
			"client": client,
		}).Warningf("empty fields error: %s", err)
		return &pb.IdClient{Id: -1}, err
	}
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
		s := status.Newf(codes.Internal, "got an error when tried to create account: %s, with error: %w", client, err)
		errWithDetails, err := s.WithDetails(client)
		if err != nil {
			return &pb.IdClient{Id: -1}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.IdClient{Id: -1}, errWithDetails.Err()
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"client": entity,
	}).Info("create client")
	return &pb.IdClient{Id: int64(id)}, nil
}

func (c ClientServer) DeleteClient(ctx context.Context, client *pb.IdClient) (*pb.StatusClientResponse, error) {
	if err := checkId(client.GetId()); err != nil {
		log.WithFields(log.Fields{
			"client": client,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusClientResponse{Message: "empty fields error"}, err
	}
	var entity = new(data.Division)
	entity.Id = int(client.Id)
	err := c.data.DeleteByIdClient(entity.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"client": entity,
		}).Warningf("got an error when tried to delete client: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to delete client: %s, with error: %w", client, err)
		errWithDetails, err := s.WithDetails(client)
		if err != nil {
			return &pb.StatusClientResponse{Message: "got an error when tried to delete client"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusClientResponse{Message: "got an error when tried to delete client"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"client": entity,
	}).Info("client was delete")
	return &pb.StatusClientResponse{Message: "client was delete"}, nil
}

func (c ClientServer) UpdateClient(ctx context.Context, client *pb.DataClient) (*pb.StatusClientResponse, error) {
	if err := checkId(client.GetId()); err != nil {
		log.WithFields(log.Fields{
			"client": client,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusClientResponse{Message: "empty fields error"}, err
	}
	if err := checkClientRequest(client); err != nil {
		log.WithFields(log.Fields{
			"client": client,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusClientResponse{Message: "empty fields error"}, err
	}

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
		s := status.Newf(codes.Internal, "got an error when tried to update client: %s, with error: %w", client, err)
		errWithDetails, err := s.WithDetails(client)
		if err != nil {
			return &pb.StatusClientResponse{Message: "got an error when tried to update client"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusClientResponse{Message: "got an error when tried to update client"}, errWithDetails.Err()
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

func checkClientRequest(in *pb.DataClient) error {
	if in.GetFirstNameC() == "" {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {FirstNameC}: %s", in.GetFirstNameC())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetLastNameC() == "" {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {LastNameC}: %s", in.GetLastNameC())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetPhoneNumberC() == "" {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {PhoneNumberC}: %s", in.GetPhoneNumberC())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetMiddleNameC() == "" {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {MiddleNameC}: %s", in.GetMiddleNameC())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	return nil
}

func checkId(id int64) error {
	if id <= 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {Id}: %s", id)
		return s.Err()
	}
	return nil
}

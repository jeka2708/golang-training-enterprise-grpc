package api

import (
	"context"
	"github.com/jeka2708/golang-training-enterprise-grpc/pkg/data"
	pb "github.com/jeka2708/golang-training-enterprise-grpc/proto/go_proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServiceServer struct {
	data *data.ServiceData
}

func NewServiceServer(s data.ServiceData) *ServiceServer {
	return &ServiceServer{data: &s}
}

func (s ServiceServer) ReadAllService(ctx context.Context, request *pb.ListServiceRequest) (*pb.ListServiceResponse, error) {
	ss, err := s.data.ReadAllServices()
	if err != nil {
		log.Println(err)
	}

	var list []*pb.DataService
	for _, t := range ss {

		list = append(list, structServiceToRes(t))

	}

	return &pb.ListServiceResponse{Data: list}, nil
}

func (s ServiceServer) CreateService(ctx context.Context, service *pb.DataService) (*pb.IdService, error) {
	if err := checkServiceRequest(service); err != nil {
		log.WithFields(log.Fields{
			"client": service,
		}).Warningf("empty fields error: %s", err)
		return &pb.IdService{Id: -1}, err
	}
	var entity = data.Service{
		Name: service.GetName(),
		Cost: int(service.Cost),
	}
	id, err := s.data.AddService(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"division": entity,
		}).Warningf("got an error when tried to create service: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to create service: %s, with error: %w", service, err)
		errWithDetails, err := s.WithDetails(service)
		if err != nil {
			return &pb.IdService{Id: -1}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.IdService{Id: -1}, errWithDetails.Err()
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"service": entity,
	}).Info("create service")
	return &pb.IdService{Id: int64(id)}, nil
}

func (s ServiceServer) DeleteService(ctx context.Context, service *pb.IdService) (*pb.StatusServiceResponse, error) {
	if err := checkId(service.GetId()); err != nil {
		log.WithFields(log.Fields{
			"service": service,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusServiceResponse{Message: "empty fields error"}, err
	}
	var entity = new(data.Service)
	entity.Id = int(service.Id)
	err := s.data.DeleteByIdService(entity.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"service": entity,
		}).Warningf("got an error when tried to delete service: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to delete client: %s, with error: %w", service, err)
		errWithDetails, err := s.WithDetails(service)
		if err != nil {
			return &pb.StatusServiceResponse{Message: "got an error when tried to delete service"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusServiceResponse{Message: "got an error when tried to delete service"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"service": entity,
	}).Info("division was delete")
	return &pb.StatusServiceResponse{Message: "service was delete"}, nil
}

func (s ServiceServer) UpdateService(ctx context.Context, service *pb.DataService) (*pb.StatusServiceResponse, error) {
	if err := checkId(service.GetId()); err != nil {
		log.WithFields(log.Fields{
			"service": service,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusServiceResponse{Message: "empty fields error"}, err
	}
	if err := checkServiceRequest(service); err != nil {
		log.WithFields(log.Fields{
			"service": service,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusServiceResponse{Message: "empty fields error"}, err
	}
	var entity = data.Service{
		Name: service.GetName(),
		Cost: int(service.Cost),
	}
	err := s.data.UpdateService(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"service": entity,
		}).Warningf("got an error when tried to update service: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to update client: %s, with error: %w", service, err)
		errWithDetails, err := s.WithDetails(service)
		if err != nil {
			return &pb.StatusServiceResponse{Message: "got an error when tried to update service"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusServiceResponse{Message: "got an error when tried to update service"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"service": entity,
	}).Info("service was update")
	return &pb.StatusServiceResponse{Message: "service was update"}, nil
}

func structServiceToRes(data data.Service) *pb.DataService {

	id := data.Id

	d := &pb.DataService{
		Name: data.Name,
		Cost: int64(data.Cost),
	}

	if id != 0 {
		d.Id = int64(id)
	}

	return d
}
func checkServiceRequest(in *pb.DataService) error {
	if in.GetName() == "" {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {GetName}: %s", in.GetName())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	return nil
}

package api

import (
	"context"
	"github.com/jeka2708/golang-training-enterprise-grpc/pkg/data"
	pb "github.com/jeka2708/golang-training-enterprise-grpc/proto/go_proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WorkServer struct {
	data *data.WorkData
}

func NewWorkServer(d data.WorkData) *WorkServer {
	return &WorkServer{data: &d}
}

func (w WorkServer) ReadAllWork(ctx context.Context, request *pb.ListWorkRequest) (*pb.ListWorkResponse, error) {
	dv, err := w.data.ReadAllWorks()
	if err != nil {
		log.Println(err)
	}

	var list []*pb.DataWork
	for _, d := range dv {

		list = append(list, structWorkToRes(d))

	}

	return &pb.ListWorkResponse{Data: list}, nil
}

func (w WorkServer) CreateWork(ctx context.Context, work *pb.DataWork) (*pb.IdWork, error) {
	if err := checkWorkRequest(work); err != nil {
		log.WithFields(log.Fields{
			"work": work,
		}).Warningf("empty fields error: %s", err)
		return &pb.IdWork{Id: -1}, err
	}
	var entity = data.Work{
		WorkerId:  int(work.WorkerId),
		ServiceId: int(work.ServiceId),
	}
	id, err := w.data.AddWork(entity.WorkerId, entity.ServiceId)
	if err != nil {
		log.WithFields(log.Fields{
			"work": entity,
		}).Warningf("got an error when tried to create work: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to create work: %s, with error: %w", work, err)
		errWithDetails, err := s.WithDetails(work)
		if err != nil {
			return &pb.IdWork{Id: -1}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.IdWork{Id: -1}, errWithDetails.Err()
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"work": entity,
	}).Info("create work")
	return &pb.IdWork{Id: int64(id)}, nil
}

func (w WorkServer) DeleteWork(ctx context.Context, work *pb.IdWork) (*pb.StatusWorkResponse, error) {
	if err := checkId(work.GetId()); err != nil {
		log.WithFields(log.Fields{
			"work": work,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusWorkResponse{Message: "empty fields error"}, err
	}
	var entity = new(data.Division)
	entity.Id = int(work.Id)
	err := w.data.DeleteByIdWork(entity.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"work": entity,
		}).Warningf("got an error when tried to delete work: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to delete client: %s, with error: %w", work, err)
		errWithDetails, err := s.WithDetails(work)
		if err != nil {
			return &pb.StatusWorkResponse{Message: "got an error when tried to delete work"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusWorkResponse{Message: "got an error when tried to delete work"}, errWithDetails.Err()

	}
	log.WithFields(log.Fields{
		"work": entity,
	}).Info("work was delete")
	return &pb.StatusWorkResponse{Message: "work was delete"}, nil
}

func (w WorkServer) UpdateWork(ctx context.Context, work *pb.DataWork) (*pb.StatusWorkResponse, error) {
	if err := checkId(work.GetId()); err != nil {
		log.WithFields(log.Fields{
			"work": work,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusWorkResponse{Message: "empty fields error"}, err
	}
	if err := checkWorkRequest(work); err != nil {
		log.WithFields(log.Fields{
			"work": work,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusWorkResponse{Message: "empty fields error"}, err
	}
	var entity = data.Work{
		WorkerId:  int(work.WorkerId),
		ServiceId: int(work.ServiceId),
	}
	err := w.data.UpdateWork(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"work": entity,
		}).Warningf("got an error when tried to update work: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to update client: %s, with error: %w", work, err)
		errWithDetails, err := s.WithDetails(work)
		if err != nil {
			return &pb.StatusWorkResponse{Message: "got an error when tried to update work"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusWorkResponse{Message: "got an error when tried to update work"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"work": entity,
	}).Info("work was update")
	return &pb.StatusWorkResponse{Message: "work was update"}, nil
}

func structWorkToRes(data data.ResultWork) *pb.DataWork {

	id := data.Id

	d := &pb.DataWork{
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		MiddleName:  data.MiddleName,
		PhoneNumber: data.PhoneNumber,
		Name:        data.Name,
	}

	if id != 0 {
		d.Id = int64(id)
	}

	return d
}
func checkWorkRequest(in *pb.DataWork) error {
	if in.GetWorkerId() == 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {WorkerId}: %s", in.GetName())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetServiceId() == 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {ServiceId}: %s", in.GetName())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	return nil
}

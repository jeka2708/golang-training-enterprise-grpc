package api

import (
	"context"
	"github.com/jeka2708/golang-training-enterprise-grpc/pkg/data"
	pb "github.com/jeka2708/golang-training-enterprise-grpc/proto/go_proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WorkerServer struct {
	data *data.WorkerData
}

func NewWorkerServer(s data.WorkerData) *WorkerServer {
	return &WorkerServer{data: &s}
}
func (w WorkerServer) ReadAllWorker(ctx context.Context, request *pb.ListWorkerRequest) (*pb.ListWorkerResponse, error) {
	ww, err := w.data.ReadAllWorkers()
	if err != nil {
		log.Println(err)
	}

	var list []*pb.DataWorker
	for _, t := range ww {

		list = append(list, structWorkerToRes(t))

	}

	return &pb.ListWorkerResponse{Data: list}, nil
}

func (w WorkerServer) CreateWorker(ctx context.Context, worker *pb.DataWorker) (*pb.IdWorker, error) {
	if err := checkWorkerRequest(worker); err != nil {
		log.WithFields(log.Fields{
			"worker": worker,
		}).Warningf("empty fields error: %s", err)
		return &pb.IdWorker{Id: -1}, err
	}
	var entity = data.Worker{
		FirstName:   worker.FirstName,
		LastName:    worker.LastName,
		MiddleName:  worker.MiddleName,
		PhoneNumber: worker.PhoneNumber,
		RoleId:      int(worker.RoleId),
	}
	id, err := w.data.AddWorker(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"worker": entity,
		}).Warningf("got an error when tried to create worker: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to create worker: %s, with error: %w", worker, err)
		errWithDetails, err := s.WithDetails(worker)
		if err != nil {
			return &pb.IdWorker{Id: -1}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.IdWorker{Id: -1}, errWithDetails.Err()
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"worker": entity,
	}).Info("create worker")
	return &pb.IdWorker{Id: int64(id)}, nil
}

func (w WorkerServer) DeleteWorker(ctx context.Context, worker *pb.IdWorker) (*pb.StatusWorkerResponse, error) {
	if err := checkId(worker.GetId()); err != nil {
		log.WithFields(log.Fields{
			"worker": worker,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusWorkerResponse{Message: "empty fields error"}, err
	}
	var entity = new(data.Division)
	entity.Id = int(worker.Id)
	err := w.data.DeleteByIdWorker(entity.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"worker": entity,
		}).Warningf("got an error when tried to delete worker: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to delete worker: %s, with error: %w", worker, err)
		errWithDetails, err := s.WithDetails(worker)
		if err != nil {
			return &pb.StatusWorkerResponse{Message: "got an error when tried to delete worker"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusWorkerResponse{Message: "got an error when tried to delete worker"}, errWithDetails.Err()

	}
	log.WithFields(log.Fields{
		"worker": entity,
	}).Info("client was delete")
	return &pb.StatusWorkerResponse{Message: "worker was delete"}, nil
}

func (w WorkerServer) UpdateWorker(ctx context.Context, worker *pb.DataWorker) (*pb.StatusWorkerResponse, error) {
	if err := checkId(worker.GetId()); err != nil {
		log.WithFields(log.Fields{
			"worker": worker,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusWorkerResponse{Message: "empty fields error"}, err
	}
	if err := checkWorkerRequest(worker); err != nil {
		log.WithFields(log.Fields{
			"worker": worker,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusWorkerResponse{Message: "empty fields error"}, err
	}
	var entity = data.Worker{
		FirstName:   worker.FirstName,
		LastName:    worker.LastName,
		MiddleName:  worker.MiddleName,
		PhoneNumber: worker.PhoneNumber,
		RoleId:      int(worker.RoleId),
	}
	err := w.data.UpdateWorker(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"worker": entity,
		}).Warningf("got an error when tried to update worker: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to update client: %s, with error: %w", worker, err)
		errWithDetails, err := s.WithDetails(worker)
		if err != nil {
			return &pb.StatusWorkerResponse{Message: "got an error when tried to update worker"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusWorkerResponse{Message: "got an error when tried to update worker"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"worker": entity,
	}).Info("worker was update")
	return &pb.StatusWorkerResponse{Message: "worker was update"}, nil
}

func structWorkerToRes(data data.ResultWorker) *pb.DataWorker {

	id := data.Id

	d := &pb.DataWorker{
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
func checkWorkerRequest(in *pb.DataWorker) error {
	if in.GetFirstName() == "" {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {FirstName}: %s", in.GetFirstName())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetLastName() == "" {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {LastName}: %s", in.GetLastName())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetPhoneNumber() == "" {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {PhoneNumber}: %s", in.GetPhoneNumber())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetMiddleName() == "" {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {MiddleName}: %s", in.GetMiddleName())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetRoleId() == 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {RoleId}: %s", in.GetMiddleName())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	return nil
}

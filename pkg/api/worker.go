package api

import (
	"context"
	"fmt"
	"github.com/jeka2708/golang-training-enterprise-grpc/pkg/data"
	pb "github.com/jeka2708/golang-training-enterprise-grpc/proto/go_proto"
	log "github.com/sirupsen/logrus"
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
		return &pb.IdWorker{Id: -1}, fmt.Errorf("got an error when tried to create worker: %w", err)
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"worker": entity,
	}).Info("create worker")
	return &pb.IdWorker{Id: int64(id)}, nil
}

func (w WorkerServer) DeleteWorker(ctx context.Context, worker *pb.IdWorker) (*pb.StatusWorkerResponse, error) {
	var entity = new(data.Division)
	entity.Id = int(worker.Id)
	err := w.data.DeleteByIdWorker(entity.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"worker": entity,
		}).Warningf("got an error when tried to delete worker: %s", err)
		return &pb.StatusWorkerResponse{Message: "got an error when tried to delete worker"},
			fmt.Errorf("got an error when tried to delete worker: %w", err)
	}
	log.WithFields(log.Fields{
		"worker": entity,
	}).Info("client was delete")
	return &pb.StatusWorkerResponse{Message: "worker was delete"}, nil
}

func (w WorkerServer) UpdateWorker(ctx context.Context, worker *pb.DataWorker) (*pb.StatusWorkerResponse, error) {
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
		return &pb.StatusWorkerResponse{Message: "got an error when tried to delete worker"},
			fmt.Errorf("got an error when tried to update worker: %w", err)
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

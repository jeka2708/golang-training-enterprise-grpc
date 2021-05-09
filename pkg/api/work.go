package api

import (
	"context"
	"fmt"
	"github.com/jeka2708/golang-training-enterprise-grpc/pkg/data"
	pb "github.com/jeka2708/golang-training-enterprise-grpc/proto/go_proto"
	log "github.com/sirupsen/logrus"
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
	var entity = data.Work{
		WorkerId:  int(work.WorkerId),
		ServiceId: int(work.ServiceId),
	}
	id, err := w.data.AddWork(entity.WorkerId, entity.ServiceId)
	if err != nil {
		log.WithFields(log.Fields{
			"work": entity,
		}).Warningf("got an error when tried to create work: %s", err)
		return &pb.IdWork{Id: -1}, fmt.Errorf("got an error when tried to create work: %w", err)
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"work": entity,
	}).Info("create work")
	return &pb.IdWork{Id: int64(id)}, nil
}

func (w WorkServer) DeleteWork(ctx context.Context, work *pb.IdWork) (*pb.StatusWorkResponse, error) {
	var entity = new(data.Division)
	entity.Id = int(work.Id)
	err := w.data.DeleteByIdWork(entity.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"work": entity,
		}).Warningf("got an error when tried to delete work: %s", err)
		return &pb.StatusWorkResponse{Message: "got an error when tried to delete work"},
			fmt.Errorf("got an error when tried to delete work: %w", err)
	}
	log.WithFields(log.Fields{
		"work": entity,
	}).Info("work was delete")
	return &pb.StatusWorkResponse{Message: "work was delete"}, nil
}

func (w WorkServer) UpdateWork(ctx context.Context, work *pb.DataWork) (*pb.StatusWorkResponse, error) {
	var entity = data.Work{
		WorkerId:  int(work.WorkerId),
		ServiceId: int(work.ServiceId),
	}
	err := w.data.UpdateWork(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"work": entity,
		}).Warningf("got an error when tried to update work: %s", err)
		return &pb.StatusWorkResponse{Message: "got an error when tried to delete work"},
			fmt.Errorf("got an error when tried to update work: %w", err)
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

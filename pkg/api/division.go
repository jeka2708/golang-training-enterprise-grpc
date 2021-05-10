package api

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/jeka2708/golang-training-enterprise-grpc/pkg/data"
	pb "github.com/jeka2708/golang-training-enterprise-grpc/proto/go_proto"
)

type DivisionServer struct {
	data *data.DivisionData
}

func NewDivisionServer(d data.DivisionData) *DivisionServer {
	return &DivisionServer{data: &d}
}

func (u DivisionServer) ReadAllDivision(ctx context.Context, division *pb.ListDivisionRequest) (*pb.ListDivisionResponse, error) {
	dv, err := u.data.ReadAllDivision()
	if err != nil {
		log.Println(err)
	}

	var list []*pb.DataDivision
	for _, d := range dv {

		list = append(list, structDivisionToRes(d))

	}

	return &pb.ListDivisionResponse{Data: list}, nil
}
func (u DivisionServer) CreateDivision(ctx context.Context, division *pb.DataDivision) (*pb.IdDivision, error) {
	if err := checkDivisionRequest(division); err != nil {
		log.WithFields(log.Fields{
			"client": division,
		}).Warningf("empty fields error: %s", err)
		return &pb.IdDivision{Id: -1}, err
	}
	var entity = data.Division{
		DivisionName: division.GetDivisionName(),
	}
	id, err := u.data.AddDivision(division.DivisionName)
	if err != nil {
		log.WithFields(log.Fields{
			"division": entity,
		}).Warningf("got an error when tried to create divisiion: %s", err)

		s := status.Newf(codes.Internal, "got an error when tried to create division: %s, with error: %w", division, err)
		errWithDetails, err := s.WithDetails(division)
		if err != nil {
			return &pb.IdDivision{Id: -1}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.IdDivision{Id: -1}, errWithDetails.Err()

	}
	entity.Id = id
	log.WithFields(log.Fields{
		"division": entity,
	}).Info("create division")
	return &pb.IdDivision{Id: int64(id)}, nil
}

func (u DivisionServer) DeleteDivision(ctx context.Context, division *pb.IdDivision) (*pb.StatusDivisionResponse, error) {
	if err := checkId(division.GetId()); err != nil {
		log.WithFields(log.Fields{
			"client": division,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusDivisionResponse{Message: "empty fields error"}, err
	}
	var entity = new(data.Division)
	entity.Id = int(division.Id)
	err := u.data.DeleteByIdDivision(entity.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"division": entity,
		}).Warningf("got an error when tried to delete divisiion: %s", err)

		s := status.Newf(codes.Internal, "got an error when tried to delete client: %s, with error: %w", division, err)
		errWithDetails, err := s.WithDetails(division)
		if err != nil {
			return &pb.StatusDivisionResponse{Message: "got an error when tried to delete division"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusDivisionResponse{Message: "got an error when tried to delete division"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"division": entity,
	}).Info("division was delete")
	return &pb.StatusDivisionResponse{Message: "division was delete"}, nil
}

func (u DivisionServer) UpdateDivision(ctx context.Context, division *pb.DataDivision) (*pb.StatusDivisionResponse, error) {
	if err := checkId(division.GetId()); err != nil {
		log.WithFields(log.Fields{
			"client": division,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusDivisionResponse{Message: "empty fields error"}, err
	}
	if err := checkDivisionRequest(division); err != nil {
		log.WithFields(log.Fields{
			"client": division,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusDivisionResponse{Message: "empty fields error"}, err
	}
	var entity = data.Division{
		Id:           int(division.Id),
		DivisionName: division.GetDivisionName(),
	}
	err := u.data.UpdateDivision(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"division": entity,
		}).Warningf("got an error when tried to update divisiion: %s", err)
		return &pb.StatusDivisionResponse{Message: "got an error when tried to delete division"},
			fmt.Errorf("got an error when tried to update division: %w", err)
	}
	log.WithFields(log.Fields{
		"division": entity,
	}).Info("division was update")
	return &pb.StatusDivisionResponse{Message: "division was update"}, nil
}

func structDivisionToRes(data data.Division) *pb.DataDivision {

	id := data.Id

	d := &pb.DataDivision{
		DivisionName: data.DivisionName,
	}

	if id != 0 {
		d.Id = int64(id)
	}

	return d
}
func checkDivisionRequest(in *pb.DataDivision) error {
	if in.GetDivisionName() == "" {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {DivisionName}: %s", in.GetDivisionName())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	return nil
}

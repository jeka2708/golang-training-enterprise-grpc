package api

import (
	"context"
	"github.com/jeka2708/golang-training-enterprise-grpc/pkg/data"
	pb "github.com/jeka2708/golang-training-enterprise-grpc/proto/go_proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RoleServer struct {
	data *data.RoleData
}

func NewRoleServer(r data.RoleData) *RoleServer {
	return &RoleServer{data: &r}
}

func (r RoleServer) ReadAllRole(ctx context.Context, request *pb.ListRoleRequest) (*pb.ListRoleResponse, error) {
	rl, err := r.data.ReadAllRoles()
	if err != nil {
		log.Println(err)
	}

	var list []*pb.DataRole
	for _, t := range rl {

		list = append(list, structRoleToRes(t))

	}

	return &pb.ListRoleResponse{Data: list}, nil
}

func (r RoleServer) DeleteRole(ctx context.Context, role *pb.IdRole) (*pb.StatusRoleResponse, error) {
	if err := checkId(role.GetId()); err != nil {
		log.WithFields(log.Fields{
			"client": role,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusRoleResponse{Message: "empty fields error"}, err
	}
	var entity = new(data.Role)
	entity.Id = int(role.Id)
	err := r.data.DeleteByIdRole(entity.Id)
	if err != nil {
		log.WithFields(log.Fields{
			"division": entity,
		}).Warningf("got an error when tried to delete Role: %s", err)

		s := status.Newf(codes.Internal, "got an error when tried to delete client: %s, with error: %w", role, err)
		errWithDetails, err := s.WithDetails(role)
		if err != nil {
			return &pb.StatusRoleResponse{Message: "got an error when tried to delete Role"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusRoleResponse{Message: "got an error when tried to delete Role"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"Role": entity,
	}).Info("Role was delete")
	return &pb.StatusRoleResponse{Message: "Role was delete"}, nil
}

func structRoleToRes(data data.ResultRoles) *pb.DataRole {

	id := data.Id

	d := &pb.DataRole{
		Name:         data.Name,
		DivisionName: data.DivisionName,
	}

	if id != 0 {
		d.Id = int64(id)
	}

	return d
}

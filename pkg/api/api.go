package api

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jeka2708/golang-training-enterprise-grpc/pkg/data"
	pb "github.com/jeka2708/golang-training-enterprise-grpc/proto/go_proto"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
)

func RegisterAllServices(server *grpc.Server, conn *gorm.DB) {
	pb.RegisterDivisionServiceServer(server, NewDivisionServer(*data.NewDivisionData(conn)))
	pb.RegisterRoleServiceServer(server, NewRoleServer(*data.NewRoleData(conn)))
	pb.RegisterServiceServiceServer(server, NewServiceServer(*data.NewServiceData(conn)))
	pb.RegisterClientServiceServer(server, NewClientServer(*data.NewClientData(conn)))
	pb.RegisterWorkerServiceServer(server, NewWorkerServer(*data.NewWorkerData(conn)))
	pb.RegisterWorkServiceServer(server, NewWorkServer(*data.NewWorkData(conn)))
	pb.RegisterWorkClientServiceServer(server, NewWorkClientServer(*data.NewWorkClientData(conn)))
}

func RegisterAllServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) {
	err := pb.RegisterDivisionServiceHandler(ctx, mux, conn)
	if err != nil {
		log.Fatal(err)
	}
	err = pb.RegisterRoleServiceHandler(ctx, mux, conn)
	if err != nil {
		log.Fatal(err)
	}
	err = pb.RegisterServiceServiceHandler(ctx, mux, conn)
	if err != nil {
		log.Fatal(err)
	}
	err = pb.RegisterClientServiceHandler(ctx, mux, conn)
	if err != nil {
		log.Fatal(err)
	}
	err = pb.RegisterWorkerServiceHandler(ctx, mux, conn)
	if err != nil {
		log.Fatal(err)
	}
	err = pb.RegisterWorkServiceHandler(ctx, mux, conn)
	if err != nil {
		log.Fatal(err)
	}
}

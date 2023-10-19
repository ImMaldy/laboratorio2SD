package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/ImMaldy/laboratorio2SD/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedEnviarServer
}

func (s *server) Enviar_Continente(ctx context.Context, req *pb.Continente) (*pb.ResponderTodo, error) {
	fmt.Println(req.Nombre + " " + req.Apellido + " " + req.Estado)
	return &pb.ResponderTodo{
		Mensaje: "Estado Recibido",
	}, nil
}

func (s *server) RegistrarEn_Datanode(ctx context.Context, req *pb.RegistroDatanode) (*pb.ResponderTodo, error) {
	fmt.Println(req.Id + " " + req.Nombre + " " + req.Apellido)
	return &pb.ResponderTodo{
		Mensaje: "ok",
	}, nil
}

func (s *server) ConsultarPor_Id(ctx context.Context, req *pb.ConsultaId) (*pb.RespuestaId, error) {
	return &pb.RespuestaId{
		Nombre:   "Ignacio",
		Apellido: "Maldonado",
	}, nil
}

func (s *server) ConsultarPor_Estado(ctx context.Context, req *pb.ConsultaEstado) (*pb.RespuestaEstado, error) {
	return nil, nil
}

func main() {
	listner, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterEnviarServer(serv, &server{})

	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())
	}

}

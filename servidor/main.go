package main

import (
	"context"
	"net"

	pb "github.com/ImMaldy/laboratorio2SD/tree/4ab5257af5d27c40df8923a5e516cd35b314c830/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedEnviarServer
}

func (s *server) Enviar_Continente(ctx context.Context, req *pb.Continente) (*pb.ResponderTodo, error) {
	return nil, nil
}

func (s *server) RegistrarEn_Datanode(ctx context.Context, req *pb.RegistroDatanode) (*pb.ResponderTodo, error) {
	return nil, nil
}

func (s *server) ConsultarPor_Id(ctx context.Context, req *pb.ConsultaId) (*pb.RespuestaId, error) {
	return nil, nil
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

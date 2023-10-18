package main

import (
	"context"
	"fmt"

	pb "github.com/ImMaldy/laboratorio2SD/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic("error conexion" + err.Error())
	}

	serviceClient := pb.NewEnviarClient(conn)
	nombre1 := "ignacio"

	//------Continentes envian info a OMS-------------

	res, err := serviceClient.Enviar_Continente(context.Background(), &pb.Continente{

		Nombre:   nombre1,
		Apellido: "maldonado",
		Estado:   "infectado",
	})

	if err != nil {
		panic("error continente" + err.Error())
	}
	fmt.Println(res.Mensaje)

	//-----------OMS envia info a un Datanode---------------
	/*
		res, err := serviceClient.RegistrarEn_Datanode(context.Background(), &pb.RegistroDatanode{
			Id:       "1",
			Nombre:   nombre1,
			Apellido: "Maldonado",
		})

		if err != nil {
			panic("error continente" + err.Error())
		}
		fmt.Println(res.Mensaje)
	*/

	//-----------OMS consulta id a Datanode------------
	// res, err := serviceClient.ConsultarPor_Id(context.Background(), &pb.ConsultaId{
	// 	Id: "1",
	// })

	// if err != nil {
	// 	panic("error continente" + err.Error())
	// }
	// fmt.Println(res.Nombre)
}

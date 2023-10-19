package main

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	pb "github.com/ImMaldy/laboratorio2SD/proto"
	"google.golang.org/grpc"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	//Leer archivo
	archivo, err := os.Open("names.txt")

	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}

	defer archivo.Close() // Asegúrate de cerrar el archivo al final

	// Lee el archivo línea por línea
	var nombres_y_apellidos []string
	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		nombres_y_apellidos = append(nombres_y_apellidos, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	// Seleccionar aleatoriamente 5 nombres y apellidos
	var personas_aleatorias []string
	for i := 0; i < 5; i++ {
		randomIndex := rand.Intn(len(nombres_y_apellidos))
		personas_aleatorias = append(personas_aleatorias, nombres_y_apellidos[randomIndex])

		if randomIndex >= 0 && randomIndex < len(nombres_y_apellidos) {
			nombres_y_apellidos = append(nombres_y_apellidos[:randomIndex], nombres_y_apellidos[randomIndex+1:]...)
		}
	}

	// Imprimir los nombres y apellidos aleatorios

	var lista_personas_final []string
	for _, nameAndLastName := range personas_aleatorias {

		var estado string
		var informacion_persona string

		//Infectado o Muerto
		randomNumber := rand.Intn(100)
		if randomNumber < 55 {
			estado = "Infectado"
		} else {
			estado = "Muerto"
		}

		informacion_persona = nameAndLastName + " " + estado
		lista_personas_final = append(lista_personas_final, informacion_persona)
	}

	//---------------------conexion---------------------
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic("error conexion" + err.Error())
	}

	serviceClient := pb.NewEnviarClient(conn)

	//------Continentes envian info a OMS-------------
	for _, elemento := range lista_personas_final {
		persona := strings.Split(elemento, " ")

		res, err := serviceClient.Enviar_Continente(context.Background(), &pb.Continente{

			Nombre:   persona[0],
			Apellido: persona[1],
			Estado:   persona[2],
		})
		fmt.Println("Estado enviado: " + persona[0] + " " + persona[1] + " " + persona[2])

		if err != nil {
			panic("error continente" + err.Error())
		}
		fmt.Println("[OMS]: " + res.Mensaje + "\n")
	}

	//enviar cada 3 segundos
	cont := 0

	for cont <= len(nombres_y_apellidos) {
		rand.Shuffle(len(nombres_y_apellidos), func(i, j int) {
			nombres_y_apellidos[i], nombres_y_apellidos[j] = nombres_y_apellidos[j], nombres_y_apellidos[i]
		})
		nueva_persona := nombres_y_apellidos[0]
		persona := strings.Split(nueva_persona, " ")
		if len(nombres_y_apellidos) > 0 {
			nombres_y_apellidos = nombres_y_apellidos[1:]

		}
		var estado string
		randomNumber := rand.Intn(100)
		if randomNumber < 55 {
			estado = "Infectado"
		} else {
			estado = "Muerto"
		}
		res, err := serviceClient.Enviar_Continente(context.Background(), &pb.Continente{

			Nombre:   persona[0],
			Apellido: persona[1],
			Estado:   estado,
		})
		fmt.Println("Estado enviado: " + persona[0] + " " + persona[1] + " " + estado)

		if err != nil {
			panic("error continente" + err.Error())
		}
		fmt.Println("[OMS]: " + res.Mensaje + "\n")

		time.Sleep(3 * time.Second)
		cont++
	}

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

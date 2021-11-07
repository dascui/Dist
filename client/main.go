package main

import (
	"context"
	"fmt"
	"log"
	"os"
//	"strconv"
	"time"

	pb "example.com/pb/pb"
	"google.golang.org/grpc"
)
//Castea int32 to string
func String(n int32) string {
    buf := [11]byte{}
    pos := len(buf)
    i := int64(n)
    signed := i < 0
    if signed {
        i = -i
    }
    for {
        pos--
        buf[pos], i = '0'+byte(i%10), i/10
        if i == 0 {
            if signed {
                pos--
                buf[pos] = '-'
            }
            return string(buf[pos:])
        }
    }
}
func avanzar() int32 {
	var eleccion int32 
	log.Printf("Para avanzar ingresa un numero de 1 al 10/n")
	fmt.Scanln(&eleccion)
	return eleccion
}
func muerte() {
	//terminate
}


func solicitarEntrada() string {
	return "Jugador ID: " + "01" + "Quiere entrar a Squid Game"
}

func main() {
	var num int32=1
	conn, err := grpc.Dial("localhost:50055", grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serviceClient := pb.NewLiderClient(conn)

	res, err := serviceClient.Entrar(context.Background(), &pb.EntradaJugador{Aceptar: true})
	if err != nil {
		panic("Cupos de Squid Game llenos" + err.Error())
	}
	var Id string = String(res.IdJugador)
	log.Printf("Has a entrado a Squid Game eres el numero " + Id)
	log.Printf("Esperando a que el Lider de comienzo a los juegos")
	log.Printf("Si los juegos no dan inicio en 15 seg sera desconectado")
	deadline := time.Now().Add(15 * time.Second)
	for {
		res2, err := serviceClient.EntrarRespuesta(context.Background(), &pb.Jugador{IdJugador: res.IdJugador, Estado: res.Estado})
		if err != nil {
			panic("Perdida de conexion" + err.Error())
			os.Exit(1)

		}
		if res2.Aceptar == true || time.Now().After(deadline) {
			break
		}
	}
	log.Printf("Squid Game a dado Comienzo ")
	log.Printf("Primer Juego Luz Verde, Luz Roja, ingresa cuanto deseas avanzar ")
	var i int = 0
	for i <= 3 {
		num = avanzar()
		res3, err := serviceClient.Correr(context.Background(), &pb.Avance{Correr: num, Numeroj: res.IdJugador})
		if err != nil {
			panic("Perdida de conexion" + err.Error())
			os.Exit(1)
		}
		var av string= String(num) 
		if res3.Sn == 4{
			log.Printf("Corres"+ av +"pero te mueres")
			os.Exit(1)
		}else if res3.Sn == 3{
			log.Printf("Ganaste esta Ronda")
		}else {
			log.Printf("Corres"+ av +"sigues con vida")
		}
		i += 1
	}


}

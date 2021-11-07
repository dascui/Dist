package main

import (
	"context"
	"fmt"
	"time"
	"net"
    "math/rand"	
    "log"
	pb "example.com/pb/pb"
	"google.golang.org/grpc"
)


// Logica Juego 1
//Determina el avance que los jugadores no deben superar
func avanzar() int32 {
	var eleccion int32 
	fmt.Scanln(&eleccion)
	return eleccion
}
func avancebot() int32{
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 10
	var i int = (rand.Intn(max - min + 1) + min)
	j := int32(i)
	return j
} 
func ingresarPasoInicial(s []int, j int)[]int{
	s = append(s, j)
	return s
}
func pasos(s []int){}
//Revisa jugadores vs avance dado por el lider
func revisarJugadores(l int32, j int32) bool {
	if l > j {
		return false
	} 
	return true
}
//Suma los avances de cada jugador en Slice
func progresoJugador(s []int , )[]int {
	return s
}
//Eliminar jugadores, cambia valores de jugadores eliminados a 0 en la slice
func eliminarJugador(s []int, numeroj int)[]int {
	numeroj= numeroj-1
	s[numeroj]=0
	//Enviar mensaje para eliminar jugador
	return s
}
// Logica Juego 2
//Numero del Lider
func numLider() int{
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 4
	return(rand.Intn(max - min + 1) + min)
}
//Revisa el numero de jugadores restantes, revisa los valores distintos de 0 para conseguir el numero de jugadores
func numerodeJugadoresRestantes(s []int) int{
	//recorrer lista de jugadores
	var num int = 0
	for _, j := range s {
		if j != 0 {
			num += 1
		}
	}
	return num
}
//Crea equipos
func crearEquipos(e1 []int,e2 []int,s []int) ([]int,[]int){
	for _, j := range s {
		var c int = 0
		var num int = 1
		if j != 0 && c == 0{
			e1 = append(e1, num)
			num += 1
			c=1
		} else if j != 0 && c == 1{
			c=0
			e2 = append(e2, num)
			num += 1
		}
		num += 1
	}
	return e1 , e2
}
//Revisa si el jugador esta en el equipo y suma la fuerza de este al valor final
func TirarlaCuerda(e []int , numeroj int, total int, valorj int) int{
	for _, j := range e {
		if j == numeroj {
			total = valorj + total
		}
	}
	return total
}
//Compara valores
func EquipoGanador(e1 int ,e2 int ,l int) int{
	var g int = 0
	return g
}
//Eliminar jugadores, cambia valores de jugadores eliminados a 0 en la slice
func eliminaEquipo(s []int, e []int)[]int {
	for _, j := range e {
		s[j-1]=0
		//Enviar mensaje para eliminar jugador		
	}
	return s
}
//Logica Juego 3
//Revisa el numero de jugadores restantes, revisa los valores distintos de 0 para conseguir el numero de jugadores
func numerodeJugadoresRestantes2(s []int) int{
	//recorrer lista de jugadores
	var num int = 0
	for _, j := range s {
		if s[j-1] != 1 {
			num += 1
		}
	}
	return num
}
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
//Bots
func EntranlosBots(s *Lider) {
	var j string = String(s.JugadoresRestantes)
	log.Printf("Inicia Squid Game entran " + j + "/16 Jugadores, se rellenaran los cupos restantes" )
	for s.JugadoresRestantes < 16 {
		s.numerobots += 1
		Bot := Bot{
            IdJugador:   s.JugadoresRestantes,
			Estado: true,
			EsBot: true,
		}
		var b string = String((Bot.IdJugador + 1))
		log.Printf("Jugador ID:" + b + "Bot")
		s.JugadoresRestantes += 1
	}
	var r string = String(s.numerobots)
	log.Printf("hay " + r + " bots")	
	log.Printf("Se da inicio a la 1ra Etapa Luz Roja, Luz Verde")

	s.Etapa = 1
	s.Inicio = true
	s.Ronda = 1
}
//Estructura Lider
type Lider struct {
	pb.UnimplementedLiderServer

	JugadoresTotal int32
	JugadoresRestantes int32
	numerobots int32
	Ronda int32
	Etapa int32
	ListaJugadoresRestantes [16]*pb.Jugador
	ListaJugadoresMuertos [16]*pb.Jugador
	Inicio bool
	Jugada []movimientos
	EnvioData int
	Res int32
	
}
type EntradaJugador struct{
    Aceptar bool;
}

type Jugador struct {
	IdJugador int32
	Estado bool
}
type Bot struct {
	IdJugador int32
	Estado bool
	EsBot bool
}
 type movimientos struct{
	IdJugador int32
	Movimientos string
 }

type movimiento struct{
	IdJugador int
	Movimiento int
}

func (s *Lider) Entrar(ctx context.Context, req *pb.EntradaJugador) (*pb.Jugador, error) {
	if s.JugadoresRestantes < 16 {
		s.JugadoresRestantes += 1
		var j string = String(s.JugadoresRestantes)
		log.Printf("El jugador " + j + " ha entrado a Squid Game")
        jugador := &pb.Jugador{
            IdJugador:   s.JugadoresRestantes,
            Estado: true,
        }
        s.ListaJugadoresRestantes [s.JugadoresRestantes - 1] = jugador
		s.JugadoresTotal = s.JugadoresRestantes
		s.Inicio= false
		s.numerobots=0
		s.Res=0
		s.EnvioData=0
		log.Printf("Squid Game: Para dar comienzo a los juegos ingresa 1")
		var eleccion int 
		fmt.Scanln(&eleccion)
		switch eleccion{
		case 1:
			fmt.Println("Dando aviso a jugadores")
			EntranlosBots(s)		
		default:
			fmt.Println("Esperando mas jugadores")
		}
        return jugador, nil
	}
	return nil, nil

} 
func (s *Lider) EntrarRespuesta(ctx context.Context, req *pb.Jugador) (*pb.EntradaJugador, error) {
	entrada := &pb.EntradaJugador{
		Aceptar:   s.Inicio,
	}
	return entrada , nil
}
func (s *Lider) IniciarEtapa(ctx context.Context, req *pb.Inicio) (*pb.Jugadores, error) {
	return nil, nil
}
func (s *Lider) Correr(ctx context.Context, req *pb.Avance) (*pb.Muerte, error) {
	log.Printf("Los jugadores empezaron a correr ingresa un numero del 4 al 10")
	var muerte int32 = avanzar()
	m := &pb.Muerte{
		Sn:   0,
	}
	s.Ronda = 1
	var jugador =s.ListaJugadoresRestantes[req.Numeroj]
	if revisarJugadores(req.Correr, muerte) {
		var Savance = String(req.Correr)
		if s.Ronda == 1 {
			Movimientos := movimientos{
				IdJugador:   req.Numeroj,
				Movimientos: Savance,
			}
			s.Jugada= append(s.Jugada,Movimientos)
			m.Sn = 1
		}else {
			var auxiliar movimientos
			auxiliar = s.Jugada[req.Numeroj-1]
			var Savance = String(req.Correr)
			auxiliar.Movimientos = auxiliar.Movimientos +"/n"  + Savance
			s.Jugada[req.Numeroj-1] = auxiliar
			m.Sn = 1
		}

	}else{
		m.Sn = 4
		s.ListaJugadoresMuertos[req.Numeroj]= jugador
	}
	if s.numerobots + req.Numeroj == 16{
		var i int32 = 0
		var aux int32= 0
		for i <= s.numerobots {
			log.Printf("los bots corren")
			aux=avancebot()
			if revisarJugadores(aux, muerte) {
				var Savance = String(req.Correr)
				if s.Ronda == 1 {
					Movimientos := movimientos{
            			IdJugador:   req.Numeroj,
						Movimientos: Savance,
					}
					s.Jugada[req.Numeroj-1]=Movimientos
					s.Ronda +=1
				}else {
					var auxiliar movimientos
					auxiliar = s.Jugada[req.Numeroj-1]
					auxiliar.Movimientos = auxiliar.Movimientos +"/n" + Savance 
					s.Jugada[req.Numeroj-1] = auxiliar
					s.Ronda +=1
					if s.Ronda>4{
						s.Res = 1
					}
				}
			}else{
				m.Sn = 4
				//s.ListaJugadoresMuertos[aux2-s.numerobots]=16-s.numerobots
			}
			i += 1
	}
	return m,nil
	}

	return nil, nil
}
func (s *Lider) EntrarRespuestaNN(ctx context.Context, req *pb.Res) (*pb.Res, error) {
	if s.Res == req.Res{
		m := &pb.Res{
			Res:   0,
		}
		return m , nil
	}
	if s.Res == 1 {
		m := &pb.Res{
			Res:   1,
		}
		return m, nil
	}
	m := &pb.Res{
		Res:   0,
	}
	return m,nil
}
func (s *Lider) SolData1(ctx context.Context, req *pb.Res) (*pb.MsjLider, error) {
	var auxiliar movimientos
	auxiliar = s.Jugada[s.EnvioData]
	var aux string = String(auxiliar.IdJugador)
	m := &pb.MsjLider{
		Jugador: aux,
		Jugadas: auxiliar.Movimientos,
	}
	s.EnvioData +=1
	if s.EnvioData == 15 {
		s.Res = 0
	}
	return m,nil
}

//func (s *Lider) Tirar(ctx context.Context, req *pb.Empuje) (*pb.Resultado, error) {
//	return nil, nil
//}

func main() {
	//Slice con valores de jugadores
//	var s []int
//	var e1 []int
//	var e2 []int
	listner, err := net.Listen("tcp", ":50055")
	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}
	serv := grpc.NewServer()
	pb.RegisterLiderServer(serv, &Lider{})
	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())
	}
	//Ingresa a los jugadores al slice para llevar la cuenta de cada uno
	log.Printf("Squid Game: Para dar comienzo a los juegos ingresa 1")


}

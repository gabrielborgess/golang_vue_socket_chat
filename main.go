package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/googollee/go-socket.io"

)

func main(){
	server,err:=socketio.NewServer(nil)

	if err!=nil{
		log.Fatal(err)
	}

	//sockets
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("Usuario conectado:", s.ID())
		s.Join("chat_room")
		return nil
	})
	server.OnEvent("/", "chat message", func(s socketio.Conn, msg string) {
		//log.Println(msg)IMPRIME EL MENSAJE EN CONSOLA,SE PUEDE GUARDAR Y ASIGNAR A UNA BASE DE DATOS
		fmt.Println("chat message:", msg)
		s.Emit("reply", "chat message "+msg)
		server.BroadcastToRoom("chat_room", "chat message", msg)
	})


	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/",server)
	http.Handle("/",http.FileServer(http.Dir("./public")))
	log.Println("Servidor en el puerto 3000")
	log.Fatal(http.ListenAndServe(":3000",nil))


}
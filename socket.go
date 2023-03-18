package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("conn", s.ID())
		return nil
	})

	server.OnEvent("/", "test", func(s socketio.Conn, msg string) {
		fmt.Println("test", msg)
		s.Emit("reply", "AAA")
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))

	log.Println("aaa socket io server aaa")
	log.Fatal(http.ListenAndServe(":8000", nil))

}

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

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

	router := gin.Default()

	router.GET("/socket.io/", gin.WrapH(server))
	router.POST("/socket.io/", func(ctx *gin.Context) {
		server.ServeHTTP(ctx.Writer, ctx.Request)
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "aaa",
		})
	})

	router.Run(":8080")

}

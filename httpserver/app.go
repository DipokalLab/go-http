// GIN으로 작성된 http server

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello world")
	// http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
	// 	w.Write([]byte("hello world"))
	// })

	// http.ListenAndServe(":5100", nil)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "aaa",
		})
	})

	r.Run()
}

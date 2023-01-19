package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	router(r)

	port := 8080
	if err := r.Run(fmt.Sprintf(`:%v`, port)); err != nil {
		panic(err)
	}
}

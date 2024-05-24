package main

import (
	"daemon/handler"
	"daemon/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	flights := handler.Manipulate()

	fmt.Sprintf("Available flights : %v", len(flights))

	engine := gin.Default()
	router.InitRouter(flights).Install(engine)

	if err := engine.Run(":8080"); err != nil {
		log.Panic("error when run gin engine ")
	}
}

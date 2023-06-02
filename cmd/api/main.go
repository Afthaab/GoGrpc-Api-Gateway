package main

import (
	"fmt"
	"log"

	"github.com/apiGateway/pkg/auth"
	"github.com/apiGateway/pkg/config"
	"github.com/apiGateway/pkg/profile"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Unable to load the config file : ", err)
	}
	r := gin.Default()

	authsvc := *auth.RegisterRoutes(r, &cfg)
	profile.RegisterProileRoutes(r, &cfg, &authsvc)

	fmt.Println(authsvc)
	r.Run(cfg.Port)

}

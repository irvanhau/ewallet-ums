package cmd

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/api"
	"ewallet-ums/internal/repository"
	"ewallet-ums/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	healthCheckSvc := &services.HealthCheck{}
	healthCheckAPI := &api.HealthCheck{
		HealthCheckServices: healthCheckSvc,
	}

	registerRepo := &repository.RegisterRepository{
		DB: helpers.DB,
	}
	registerSvc := &services.RegisterService{
		RegisterRepository: registerRepo,
	}
	registerAPI := &api.RegisterHandler{
		RegisterService: registerSvc,
	}

	r := gin.Default()

	r.GET("/health", healthCheckAPI.HealthCheckHandlerHTTP)

	userV1 := r.Group("/user/v1")
	userV1.POST("/register", registerAPI.Register)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}

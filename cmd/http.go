package cmd

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/api"
	"ewallet-ums/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	healthCheckSvc := &services.HealthCheck{}
	healthCheckAPI := &api.HealthCheck{
		HealthCheckServices: healthCheckSvc,
	}

	r := gin.Default()

	r.GET("/health", healthCheckAPI.HealthCheckHandlerHTTP)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}

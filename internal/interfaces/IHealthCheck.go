package interfaces

import "github.com/gin-gonic/gin"

type IHealthCheckHandler interface {
	HealthCheckHandlerHTTP(c *gin.Context)
}

type IHealthCheckServices interface {
	HealthCheckServices() (string, error)
}

type IHealthCheckRepo interface {
}

package cmd

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/api"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/repository"
	"ewallet-ums/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	dependency := dependencyInject()
	r := gin.Default()

	r.GET("/health", dependency.HealthCheckAPI.HealthCheckHandlerHTTP)

	userV1 := r.Group("/user/v1")
	userV1.POST("/register", dependency.RegisterAPI.Register)
	userV1.POST("/login", dependency.LoginAPI.Login)

	userV1.DELETE("/logout", dependency.MiddlewareValidateAuth, dependency.LogoutAPI.Logout)
	userV1.PUT("/refresh-token", dependency.MiddlewareRefreshToken, dependency.RefreshTokenAPI.RefreshToken)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	UserRepository interfaces.IUserRepository

	HealthCheckAPI  interfaces.IHealthCheckHandler
	RegisterAPI     interfaces.IRegisterHandler
	LoginAPI        interfaces.ILoginHandler
	LogoutAPI       interfaces.ILogoutHandler
	RefreshTokenAPI interfaces.IRefreshTokenHandler

	TokenValidationAPI *api.TokenValidationHandler
}

func dependencyInject() Dependency {
	healthCheckSvc := &services.HealthCheck{}
	healthCheckAPI := &api.HealthCheck{
		HealthCheckServices: healthCheckSvc,
	}

	userRepo := &repository.UserRepository{
		DB: helpers.DB,
	}
	registerSvc := &services.RegisterService{
		UserRepository: userRepo,
	}
	registerAPI := &api.RegisterHandler{
		RegisterService: registerSvc,
	}

	loginSvc := &services.LoginService{
		UserRepository: userRepo,
	}
	loginAPI := &api.LoginHandler{
		LoginService: loginSvc,
	}

	logoutSvc := &services.LogoutService{
		UserRepository: userRepo,
	}
	logoutAPI := &api.LogoutHandler{
		LogoutService: logoutSvc,
	}

	refreshTokenSvc := &services.RefreshTokenService{
		UserRepository: userRepo,
	}
	refreshTokenAPI := &api.RefreshTokenHandler{
		RefreshTokenService: refreshTokenSvc,
	}

	tokenValidationSvc := &services.TokenValidationService{
		UserRepository: userRepo,
	}
	tokenValidationAPI := &api.TokenValidationHandler{
		TokenValidationService: tokenValidationSvc,
	}

	return Dependency{
		UserRepository:     userRepo,
		HealthCheckAPI:     healthCheckAPI,
		RegisterAPI:        registerAPI,
		LoginAPI:           loginAPI,
		LogoutAPI:          logoutAPI,
		RefreshTokenAPI:    refreshTokenAPI,
		TokenValidationAPI: tokenValidationAPI,
	}
}

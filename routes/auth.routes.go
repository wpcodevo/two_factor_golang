package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/two_factor_golang/controllers"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("auth")

	router.POST("/register", rc.authController.SignUpUser)
	router.POST("/login", rc.authController.LoginUser)
	router.POST("/otp/generate", rc.authController.GenerateOTP)
	router.POST("/otp/verify", rc.authController.VerifyOTP)
	router.POST("/otp/validate", rc.authController.ValidateOTP)
	router.POST("/otp/disable", rc.authController.DisableOTP)
}

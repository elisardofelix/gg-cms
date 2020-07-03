package Controllers

import (
	"gg-cms/DTOs"
	"gg-cms/Services"
	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context)
}

type loginController struct {
	loginService Services.LoginService
	jWtService   Services.JWTService
}

func NewLoginController(loginService Services.LoginService,
	jWtService Services.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) {
	var credentials DTOs.Credentials
	err := ctx.ShouldBindJSON(&credentials)
	if err != nil {
		ctx.JSON(401, nil)
	} else {
		isAuthenticated := controller.loginService.Login(credentials.Username, credentials.Password)
		if isAuthenticated {
			ctx.JSON(200, gin.H{
				"token" : controller.jWtService.GenerateToken(credentials.Username, true),

			})
		} else {
			ctx.JSON(401, nil)
		}
	}


}

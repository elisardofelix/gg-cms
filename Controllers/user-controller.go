package Controllers

import (
	"gg-cms/DTOs"
	"gg-cms/Models"
	"gg-cms/Services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type UserController interface {
	FindAllUsers(ctx *gin.Context)
	FindUser(ctx *gin.Context)
	SaveUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userController struct {
	service Services.UserService
}

func NewUserController(service Services.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (uc *userController) FindAllUsers(ctx *gin.Context) {
	users, err := uc.service.FindAll(20, 0)

	if err == nil {
		ctx.JSON(200, users)
	} else {
		ctx.JSON(500, gin.H {
			"error" : err.Error(),
		})
	}
}

func (uc *userController) FindUser(ctx *gin.Context) {
	permalink := ctx.Param("username")
	post, err := uc.service.Find(permalink)

	if err == nil {
		ctx.JSON(200, post)
	} else {
		ctx.JSON(500, gin.H {
			"error" : err.Error(),
		})
	}
}

func (uc *userController) SaveUser(ctx *gin.Context) {
	var userIn DTOs.Registration
	err := ctx.ShouldBindJSON(&userIn)

	if userIn.Password != userIn.RePasword || userIn.Password == "" {
		ctx.JSON(500, gin.H{
			"error": "Password and RePassword does not match.",
		})
		return
	}
	Services.EncodePassword(userIn.UserName, userIn.Password)

	var user  = Models.User{
		UserName: userIn.UserName,
		Password: Services.EncodePassword(userIn.UserName, userIn.Password),
		Email: userIn.Email,
	}

	if err != nil {
		ctx.JSON(500, gin.H{
			"error" : err.Error(),
		})
	} else {
		user.Status = "Active"
		user.CreatedDate = time.Now()
		cClaims, _ := ctx.Get("jwtClaims")
		tokenClaims := cClaims.(jwt.MapClaims)
		user.CreatedBy = tokenClaims["name"].(string)

		newUser, err := uc.service.Save(user)

		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(200, newUser)
		}
	}
}

func (uc *userController) UpdateUser(ctx *gin.Context) {
	var user Models.User
	var updatedUser DTOs.User
	err := ctx.ShouldBindJSON(&user)

	if err == nil {
		updatedUser, err = uc.service.Update(user)
		if err == nil {
			ctx.JSON(200, updatedUser)
		} else {
			ctx.JSON(500, gin.H {
				"error" : err.Error(),
			})
		}
	} else {
		ctx.JSON(500, gin.H {
			"error" : err.Error(),
		})
	}
}

func (uc *userController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	err := uc.service.Delete(id)

	if err == nil {
		ctx.JSON(200, nil)
	} else {
		ctx.JSON(500, gin.H {
			"error" : err.Error(),
		})
	}
}

package Controllers

import (
	"gg-cms/Models"
	"gg-cms/Services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type PostController interface {
	FindAllActivePosts(ctx *gin.Context)
	FindPost(ctx *gin.Context)
	SavePost(ctx *gin.Context)
	UpdatePost(ctx *gin.Context)
	DeletePost(ctx *gin.Context)
}

type postController struct {
	service Services.PostService
}

func NewPostController(service Services.PostService) PostController {
	return &postController{
		service: service,
	}
}

func (pc *postController) FindAllActivePosts(ctx *gin.Context) {
	posts, err := pc.service.FindAll(20, 0)

	if err == nil {
		ctx.JSON(200, posts)
	} else {
		ctx.JSON(500, gin.H {
			"error" : err.Error(),
		})
	}
}

func (pc *postController) FindPost(ctx *gin.Context) {
	permalink := ctx.Param("permalink")
	post, err := pc.service.Find(permalink)

	if err == nil {
		ctx.JSON(200, post)
	} else {
		ctx.JSON(500, gin.H {
			"error" : err.Error(),
		})
	}
}

func (pc *postController) SavePost(ctx *gin.Context) {
	var post Models.Post
	err := ctx.ShouldBindJSON(&post)

	if err != nil {
		ctx.JSON(500, gin.H{
			"error" : err.Error(),
		})
	} else {
		post.CreatedDate = time.Now()
		cClaims, _ := ctx.Get("jwtClaims")
		tokenClaims := cClaims.(jwt.MapClaims)
		post.CreatedBy = tokenClaims["name"].(string)

		newPost, err := pc.service.Save(post)

		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(200, newPost)
		}
	}
}

func (pc *postController) UpdatePost(ctx *gin.Context) {
	var post Models.Post
	var updatedPost Models.Post
	err := ctx.ShouldBindJSON(&post)

	if err == nil {
		updatedPost, err = pc.service.Update(post)
		if err == nil {
			ctx.JSON(200, updatedPost)
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

func (pc *postController) DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")
	err := pc.service.Delete(id)

	if err == nil {
		ctx.JSON(200, nil)
	} else {
		ctx.JSON(500, gin.H {
			"error" : err.Error(),
		})
	}
}

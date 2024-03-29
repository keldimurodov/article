package api

import (
	_ "projects/article/api-gateway/api/docs" // swag
	"projects/article/api-gateway/api/handlers/middleware"
	"projects/article/api-gateway/api/handlers/tokens"
	v1 "projects/article/api-gateway/api/handlers/v1"
	"projects/article/api-gateway/config"
	"projects/article/api-gateway/pkg/logger"
	"projects/article/api-gateway/queue/kafka/producer"
	"projects/article/api-gateway/services"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	CasbinEnforcer *casbin.Enforcer
	Writer         producer.KafkaProducer
}

// @Title Welcome to swagger service
// @Version 1.0
// @Description you can use this as social network
// @Host localhost:8080
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	router := gin.New()

	jwtHandler := tokens.JWTHandler{
		SigninKey: option.Conf.SigningKey,
	}

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
		JWTHandler:     jwtHandler,
		Enforcer:       option.CasbinEnforcer,
	})

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.NewAuthorizer(option.CasbinEnforcer, jwtHandler, option.Conf))
	// router.Use(middleware.NewAuthorizer(option.CasbinEnforcer, jwtHandler, option.Conf))

	api := router.Group("/v1")

	// users
	api.POST("/users", handlerV1.CreateUser)
	api.GET("/users/:id", handlerV1.GetUser)
	api.GET("/users", handlerV1.GetAllUsers)
	api.PUT("/users/:id", handlerV1.UpdateUser)
	api.DELETE("/users/:id", handlerV1.DeleteUser)

	// posts
	api.POST("/posts", handlerV1.CreatePost)
	api.GET("/posts/:id", handlerV1.GetPost)
	api.GET("/posts", handlerV1.GetAllPosts)
	api.PUT("/posts", handlerV1.UpdatePost)
	api.DELETE("/posts/:id", handlerV1.DeletePost)

	// user registratsiya
	api.POST("/signup", handlerV1.SignUp)
	api.GET("/login", handlerV1.LogIn)
	api.GET("/verification", handlerV1.Verification)
	api.GET("/refreshusertoken", handlerV1.RefreshUserToken)

	// user rbac
	api.POST("/create_role", handlerV1.CreateRole)

	url := ginSwagger.URL("swaggerdoc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

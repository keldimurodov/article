package v1

import (
	"projects/article/api-gateway/api/handlers/tokens"
	"projects/article/api-gateway/config"
	"projects/article/api-gateway/pkg/logger"
	"projects/article/api-gateway/queue/kafka/producer"
	"projects/article/api-gateway/services"

	"github.com/casbin/casbin/v2"
)

type handlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	cfg            config.Config
	jwthandler     tokens.JWTHandler
	enforcer       *casbin.Enforcer
	writer         producer.KafkaProducer
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Cfg            config.Config
	JWTHandler     tokens.JWTHandler
	Enforcer       *casbin.Enforcer
	Writer         producer.KafkaProducer
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
		jwthandler:     c.JWTHandler,
		enforcer:       c.Enforcer,
		writer:         c.Writer,
	}
}

package main

import (
	"projects/article/projects/article/api-gateway/api"
	"projects/article/projects/article/api-gateway/config"
	"projects/article/projects/article/api-gateway/pkg/logger"
	"projects/article/projects/article/api-gateway/queue/kafka/producer"
	"projects/article/projects/article/api-gateway/services"

	"github.com/casbin/casbin/v2"
	defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	"github.com/casbin/casbin/v2/util"
	// db
	// "fmt"
	// gormadapter "github.com/casbin/gorm-adapter/v3"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	// databasa bn accses berish
	// ____________________________________________________________________________________________________________________________________

	// psqlString := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`, "localhost", 5432, "postgres", "asadbek", "userdb")

	// db, err := gormadapter.NewAdapter("postgres", psqlString, true)

	// if err != nil {
	// 	log.Error("gormadapter error", logger.Error(err))
	// 	return
	// }

	// casbinEnforcer, err := casbin.NewEnforcer("auth.conf", db)
	// if err != nil {
	// 	log.Error("NewEnforcer error", logger.Error(err))
	// 	return
	// }
	// _______________________________________________________________________________________________________________________________________

	// csv file bn accses berish
	casbinEnforcer, err := casbin.NewEnforcer(cfg.AuthConfigPath, cfg.CSVFilePath)
	if err != nil {
		log.Fatal("casbin enforcer error", logger.Error(err))
		return
	}

	err = casbinEnforcer.LoadPolicy()
	if err != nil {
		log.Fatal("casbin error load policy", logger.Error(err))
		return
	}

	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManagerImpl).AddMatchingFunc("keyMatch", util.KeyMatch)
	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManagerImpl).AddMatchingFunc("keyMatch3", util.KeyMatch3)

	// ________________________________________________________________________________________________________________________________________
	// kafka
	writer, err := producer.NewKafkaProducerInit([]string{"locolhost:9092"})
	if err != nil {
		log.Fatal("NewKafkaProducerInit", logger.Error(err))
	}

	defer writer.Close()
	// _____________________________________________________________________________________________________________________________

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
		CasbinEnforcer: casbinEnforcer,
		Writer:         writer,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}

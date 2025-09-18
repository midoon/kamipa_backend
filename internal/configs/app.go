package configs

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/midoon/kamipa_backend/internal/controller"
	"github.com/midoon/kamipa_backend/internal/delivery/http/route"
	"github.com/midoon/kamipa_backend/internal/repository"
	"github.com/midoon/kamipa_backend/internal/usecase"
	"github.com/midoon/kamipa_backend/internal/util"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	KamipaDB *gorm.DB
	SimipaDB *gorm.DB
	Router   *mux.Router
	Validate *validator.Validate
	// config    *Config
}

func BootStrap(bs *BootstrapConfig) {

	cnf := GetConfig()

	redisClient := GetRedisClient(cnf.Redis.Addr, 0)

	tokenUtil := util.NewTokenUtil(cnf.JWT.Key, redisClient)

	// setup repository
	redisRepository := repository.NewRedisRepository(redisClient)
	userRepository := repository.NewUserRepository(bs.KamipaDB)
	studentRepository := repository.NewStudentRepository(bs.SimipaDB)

	// setup usecase
	userUsecase := usecase.NewUserUsecase(bs.Validate, userRepository, studentRepository, tokenUtil, redisRepository)

	// setup controller
	userController := controller.NewUserController(userUsecase)

	//setup middleware

	routeConfig := route.RouteConfig{
		Router:         bs.Router,
		UserController: userController,
	}

	routeConfig.Setup()

}

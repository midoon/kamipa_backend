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

	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr: "localhost:6379",
	// 	DB:   0,
	// })

	tokenUtil := util.NewTokenUtil(cnf.JWT.Key, nil)

	// setup repository
	userRepository := repository.NewUserRepository(bs.KamipaDB)
	studentRepository := repository.NewStudentRepository(bs.SimipaDB)

	// setup usecase
	userUsecase := usecase.NewUserUsecase(bs.Validate, userRepository, studentRepository, tokenUtil)

	// setup controller
	userController := controller.NewUserController(userUsecase)

	//setup middleware

	routeConfig := route.RouteConfig{
		Router:         bs.Router,
		UserController: userController,
	}

	routeConfig.Setup()

}

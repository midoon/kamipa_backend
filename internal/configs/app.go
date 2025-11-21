package configs

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/midoon/kamipa_backend/internal/controller"
	"github.com/midoon/kamipa_backend/internal/delivery/http/route"
	"github.com/midoon/kamipa_backend/internal/repository"
	"github.com/midoon/kamipa_backend/internal/usecase"
	"github.com/midoon/kamipa_backend/internal/util"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	KamipaDB    *gorm.DB
	SimipaDB    *gorm.DB
	Router      *mux.Router
	HttpClient  *http.Client
	Validate    *validator.Validate
	Cnf         *ConfigApp
	RedisClient *redis.Client
}

func BootStrap(bs *BootstrapConfig) {

	// setup repository
	redisRepository := repository.NewRedisRepository(bs.RedisClient)
	userRepository := repository.NewUserRepository(bs.KamipaDB)
	studentRepository := repository.NewStudentRepository(bs.SimipaDB)
	dashboardApiRepository := repository.NewDashboardApiRepository(bs.HttpClient, bs.Cnf.Mediamipa.BaseUrl)
	attendanceRepository := repository.NewAttendanceRepository(bs.SimipaDB)
	feeRepository := repository.NewFeeRepository(bs.SimipaDB)
	topupRepository := repository.NewTopupRepository(bs.KamipaDB)

	tokenUtil := util.NewTokenUtil(bs.Cnf.JWT.Key, redisRepository)

	// setup usecase
	userUsecase := usecase.NewUserUsecase(bs.Validate, userRepository, studentRepository, tokenUtil, redisRepository)
	dashboardUsecase := usecase.NewDashboardUsecase(dashboardApiRepository)
	attendanceUsecase := usecase.NewAttendanceUsecase(attendanceRepository, userRepository, studentRepository)
	feeUsecase := usecase.NewFeeUsecase(feeRepository, userRepository, studentRepository)
	topupUsecase := usecase.NewTopupUsecase(bs.Cnf.Midtrans.Key, bs.Cnf.Midtrans.IsProd, topupRepository, feeRepository, userRepository)

	// setup controller
	userController := controller.NewUserController(userUsecase)
	dashboardController := controller.NewDashboardController(dashboardUsecase)
	attendanceController := controller.NewAttendanceController(attendanceUsecase)
	feeController := controller.NewFeeController(feeUsecase)
	topupController := controller.NewTopupController(topupUsecase)

	//setup middleware

	routeConfig := route.RouteConfig{
		Router:               bs.Router,
		UserController:       userController,
		TokenUtil:            tokenUtil,
		DashboardController:  dashboardController,
		AttendanceController: attendanceController,
		FeeController:        feeController,
		TopupController:      topupController,
	}

	routeConfig.Setup()

}

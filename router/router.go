package router

import (
	"github.com/gin-gonic/gin"
	"kecamatan_app/controller"
	"kecamatan_app/driver"
	"kecamatan_app/middlewares"
	"kecamatan_app/models"
	"kecamatan_app/repo"
	"kecamatan_app/usecase"
)

func Router() *gin.Engine {
	router := gin.New()
	//port := os.Getenv("PORT_LOGIN")
	router.Use(middlewares.CORSMiddleware())
	db := driver.ConnectDB()
	models.InitTable(db)

	adminRepo := repo.CreateRegisterRepoImpl(db)
	adminService := usecase.CreateRegisterServiceImpl(adminRepo)

	bantuanRepo := repo.CreateBantuanRepoImpl(db)
	bantuanUsecase := usecase.CreateBantuanUsecase(bantuanRepo)

	daerahRepo := repo.CreateDaerahIndoRepo()
	daerahUsecase := usecase.CreateDaerahUsecase(daerahRepo)

	dataWargaRepo := repo.CreateDataWargaRepoImpl(db)
	dataWargaUsecase := usecase.CreateDataWargaUsecaseImpl(dataWargaRepo, daerahRepo)

	v1 := router.Group("api/v1")

	{
		newRoute := v1.Group("ms-kecamatan")
		controller.CreateAdminController(newRoute, adminService)
		controller.CreateBantuanController(newRoute, bantuanUsecase)
		controller.CreateDaerahController(newRoute, daerahUsecase)
		controller.CreateDataWargaController(newRoute, dataWargaUsecase)
	}
	//s.ListenAndServe()
	//router.Run(port)
	return router
}

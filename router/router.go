package router

import (
	"github.com/gin-gonic/gin"
	"kecamatan_app/controller"
	"kecamatan_app/driver"
	"kecamatan_app/models"
	"kecamatan_app/repo"
	"kecamatan_app/usecase"
)

func Router() *gin.Engine {
	router := gin.New()
	//port := os.Getenv("PORT_LOGIN")
	db := driver.ConnectDB()
	models.InitTable(db)

	adminRepo := repo.CreateRegisterRepoImpl(db)
	adminService := usecase.CreateRegisterServiceImpl(adminRepo)

	v1 := router.Group("api/v1")

	{
		newRoute := v1.Group("ms-kecamatan")
		controller.CreateAdminController(newRoute, adminService)
	}
	//s.ListenAndServe()
	//router.Run(port)
	return router
}

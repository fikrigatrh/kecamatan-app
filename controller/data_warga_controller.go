package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kecamatan_app/models"
	"kecamatan_app/usecase"
	"kecamatan_app/utils"
	"net/http"
)

type DataWargaController struct {
	dataWargaService usecase.DataWargaUsecaseInterface
}

func CreateDataWargaController(router *gin.RouterGroup, dataWargaService usecase.DataWargaUsecaseInterface) {
	inDB := DataWargaController{dataWargaService}

	router.POST("/add-data-warga", inDB.AddDataWarga)
	router.GET("/warga-all", inDB.GetAllDataWarga)

}

func (a *DataWargaController) AddDataWarga(c *gin.Context) {
	var dataWarga models.DataWarga

	err := c.ShouldBindJSON(&dataWarga)
	if err != nil {
		utils.ErrorMessage(c, http.StatusBadRequest, "Oppss, something error")
		fmt.Printf("[DataWargaController.AddDataWargaPemerintah] Error when decoder data from body with error : %v\n", err)
		return
	}

	data, err := a.dataWargaService.CreateDataWarga(&dataWarga)
	if err != nil {
		utils.ErrorMessage(c, http.StatusBadRequest, "OPPSSSSSSSSS")
		fmt.Printf("[DataWargaController.AddDataWargaPemerintah] Error when request data to usecase with error: %v\n", err)
		return
	}

	utils.SuccessData(c, http.StatusOK, data)
}

func (a *DataWargaController) GetAllDataWarga(c *gin.Context) {
	data, err := a.dataWargaService.GetAllData()
	if err != nil {
		utils.ErrorMessage(c, http.StatusBadRequest, "Oppss, something error ")
		fmt.Printf("[ClassController.GetAdmin] Error when request data to usecase with error: %v\n", err)
		return
	}

	utils.SuccessData(c, http.StatusOK, data)
}

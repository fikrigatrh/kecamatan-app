package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kecamatan_app/usecase"
	"kecamatan_app/utils"
	"net/http"
	"strconv"
)

type DaerahController struct {
	daerahService usecase.DaerahIndoUsecaseInterface
}

func CreateDaerahController(router *gin.RouterGroup, daerahService usecase.DaerahIndoUsecaseInterface) {
	inDB := DaerahController{daerahService}

	router.GET("/provinsi/:id", inDB.GetProvinsiByID)
}

func (a *DaerahController) GetProvinsiByID(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	data, err := a.daerahService.GetDataProvinsiByID(id)
	if err != nil {
		utils.ErrorMessage(c, http.StatusBadRequest, "Oppss, something error ")
		fmt.Printf("[ClassController.GetAdmin] Error when request data to usecase with error: %v\n", err)
		return
	}

	utils.SuccessData(c, http.StatusOK, data)
}

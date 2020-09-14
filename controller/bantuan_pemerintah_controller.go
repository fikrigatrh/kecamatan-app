package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kecamatan_app/models"
	"kecamatan_app/usecase"
	"kecamatan_app/utils"
	"net/http"
	"strconv"
)

type BantuanController struct {
	bantuanService usecase.BantuanPemerintahUsecaseInterface
}

func CreateBantuanController(router *gin.RouterGroup, bantuanService usecase.BantuanPemerintahUsecaseInterface) {
	inDB := BantuanController{bantuanService}

	router.POST("/add-bantuan", inDB.AddBantuan)
	router.GET("/bantuan-all", inDB.GetAllAdmin)
	router.GET("/bantuan/:id", inDB.GetByID)
	router.PUT("/bantuan/:id", inDB.UpdateData)

}

func (a *BantuanController) AddBantuan(c *gin.Context) {
	var bantuan models.BantuanPemerintah

	err := c.ShouldBindJSON(&bantuan)
	if err != nil {
		utils.ErrorMessage(c, http.StatusBadRequest, "Oppss, something error")
		fmt.Printf("[AdminController.AddBantuanPemerintah] Error when decoder data from body with error : %v\n", err)
		return
	}

	data, err := a.bantuanService.CreateBantuan(&bantuan)
	if err != nil {
		utils.ErrorMessage(c, http.StatusBadRequest, "Tabungan Sudah Terdaftar")
		fmt.Printf("[AdminController.AddBantuanPemerintah] Error when request data to usecase with error: %v\n", err)
		return
	}

	utils.SuccessData(c, http.StatusOK, data)
}

func (a *BantuanController) GetByID(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	data, err := a.bantuanService.GetDataByID(id)
	if err != nil {
		utils.ErrorMessage(c, http.StatusBadRequest, "Oppss, something error ")
		fmt.Printf("[StudentController.GetByID] Error when request data to usecase with error: %v\n", err)
		return
	}

	utils.SuccessData(c, http.StatusOK, data)
}

func (a *BantuanController) GetAllAdmin(c *gin.Context) {
	data, err := a.bantuanService.GetAllData()
	if err != nil {
		utils.ErrorMessage(c, http.StatusBadRequest, "Oppss, something error ")
		fmt.Printf("[ClassController.GetAdmin] Error when request data to usecase with error: %v\n", err)
		return
	}

	utils.SuccessData(c, http.StatusOK, data)
}

func (a *BantuanController) UpdateData(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	var dataCurrent models.BantuanPemerintah

	err := c.ShouldBindJSON(&dataCurrent)
	if err != nil {
		utils.ErrorMessage(c, http.StatusBadRequest, "Opps, something error")
		fmt.Printf("[StudentController.UpdateData] Error when decoder data from body with error : %v\n", err)
		return
	}

	dataCheck := a.bantuanService.CheckData(dataCurrent.NamaBantuan)
	if dataCheck == true {
		utils.ErrorMessage(c, http.StatusBadRequest, "Data Bantuan Sudahh Terdaftar")
		return
	}

	result, err := a.bantuanService.UpdateData(id, &dataCurrent)
	if err != nil {
		utils.ErrorMessage(c, http.StatusBadRequest, err.Error())
		fmt.Printf("[StudentController.updateStudent] Error when request data to usecase with error : %v", err)
		return
	}
	utils.SuccessData(c, http.StatusOK, result)
}

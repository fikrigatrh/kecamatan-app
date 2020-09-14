package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kecamatan_app/middlewares"
	"kecamatan_app/models"
	"kecamatan_app/usecase"
	"kecamatan_app/utils"
	"log"
	"net/http"
	"strconv"
)

type AdminController struct {
	adminService usecase.RegisterServiceInterface
}

func CreateAdminController(router *gin.RouterGroup, adminService usecase.RegisterServiceInterface) {
	inDB := AdminController{adminService}

	router.POST("/login", inDB.LoginController)
	router.POST("/add-role", inDB.AddRole)
	router.POST("/register", inDB.RegisterAdmin)
	router.GET("/admins", inDB.GetAllAdmin)
	router.GET("/admin/:uuid", inDB.GetByUUID)
	router.POST("/logout", middlewares.TokenAuthMiddleware(), inDB.logout)
	router.PUT("/admin/:id", inDB.UpdateData)
	router.DELETE("admin/:id", inDB.DeleteAdmin)
}

func (a *AdminController) AddRole(c *gin.Context) {
	var role models.Role

	err := c.ShouldBindJSON(&role)
	if err != nil {
		utils.ErrorMessage(c, http.StatusInternalServerError, "Oppss, something error")
		fmt.Printf("[AdminController.AddRole] Error when decoder data from body with error : %v\n", err)
		return
	}

	data, err := a.adminService.AddRole(&role)
	if err != nil {
		utils.ErrorMessage(c, http.StatusInternalServerError, "Tabungan Sudah Terdaftar")
		fmt.Printf("[AdminController.AddRole] Error when request data to usecase with error: %v\n", err)
		return
	}

	utils.SuccessData(c, http.StatusOK, data)
}

func (a *AdminController) RegisterAdmin(c *gin.Context) {
	var admin models.User

	err := c.ShouldBindJSON(&admin)
	if err != nil {
		utils.SuccessMessage(c, http.StatusBadRequest, "Oppss, something error")
		log.Printf("[RegisterAdmin.Register] Error when decoder data from body with error : %v\n", err)
		return
	}

	dataCheck := a.adminService.CheckData(admin.Username)
	if dataCheck == true {
		utils.ErrorMessage(c, http.StatusInternalServerError, "Data Sudahh Terdaftar")
		return
	}

	userHashed, errs := utils.HashPassword(&admin)
	if errs != nil {
		utils.SuccessMessage(c, http.StatusBadRequest, "Oppss, something error")
		log.Printf("[RegisterAdmin.Register] Error when HashPassword: %v\n", err)
		return
	}

	_, err = a.adminService.RegisterAdmin(userHashed)
	if err != nil {
		utils.ErrorMessage(c, http.StatusBadRequest, "username has been registered")
		log.Printf("[RegisterAdmin.Register] Error when request data to service with error: %v\n", err)
		return
	}

	utils.SuccessData(c, http.StatusOK, "Data Berhasil Ditambahkan")
}

func (a *AdminController) GetByUUID(c *gin.Context) {
	uuid := c.Param("uuid")
	data, err := a.adminService.GetAdminByUUID(uuid)
	if err != nil {
		utils.ErrorMessage(c, http.StatusInternalServerError, "Oppss, something error ")
		fmt.Printf("[StudentController.GetByID] Error when request data to usecase with error: %v\n", err)
		return
	}
	data.ID = 0
	data.Password = ""

	utils.SuccessData(c, http.StatusOK, data)
}

func (a *AdminController) GetAllAdmin(c *gin.Context) {
	dataAdmin, err := a.adminService.GetAllAdmin()
	if err != nil {
		utils.ErrorMessage(c, http.StatusInternalServerError, "Oppss, something error ")
		fmt.Printf("[ClassController.GetAdmin] Error when request data to usecase with error: %v\n", err)
		return
	}

	utils.SuccessData(c, http.StatusOK, dataAdmin)
}

func (a *AdminController) UpdateData(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var admin models.User

	err := c.ShouldBindJSON(&admin)
	if err != nil {
		utils.ErrorMessage(c, http.StatusInternalServerError, "Opps, something error")
		fmt.Printf("[StudentController.UpdateData] Error when decoder data from body with error : %v\n", err)
		return
	}

	dataCheck := a.adminService.CheckData(admin.Username)
	if dataCheck == true {
		utils.ErrorMessage(c, http.StatusInternalServerError, "Kelas Sudahh Terdaftar")
		return
	}

	data, err := a.adminService.UpdateAdmin(id, &admin)
	if err != nil {
		utils.ErrorMessage(c, http.StatusInternalServerError, err.Error())
		fmt.Printf("[StudentController.updateStudent] Error when request data to usecase with error : %v", err)
		return
	}
	utils.SuccessData(c, http.StatusOK, data)
}

func (a *AdminController) DeleteAdmin(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorMessage(c, http.StatusInternalServerError, "UUID not valid ")
		fmt.Printf("[ClassController.UpdateData] Error when convert pathvar with error: %v\n", err)
		return
	}
	err = a.adminService.DeleteAdmin(id)
	if err != nil {
		utils.ErrorMessage(c, http.StatusInternalServerError, "Oppss, something error")
		fmt.Printf("[ClassController.DeleteAdmin]Error when request data to usecase with error : %v\n", err)
		return
	}
	utils.SuccessMessage(c, http.StatusOK, "Delete data success")
}


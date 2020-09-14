package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"kecamatan_app/auth"
	"kecamatan_app/models"
	"kecamatan_app/utils"
	"log"
	"net/http"
	"os"
)

func (a *AdminController) LoginController(c *gin.Context) {
	key := os.Getenv("KEY_DECRYPT")
	var encrpytData models.Decrypt

	err := c.ShouldBindJSON(&encrpytData)
	if err != nil {
		utils.ErrorMessage(c, http.StatusBadRequest, "Ops, Error when bind json from body")
		log.Printf("[login Controller] error when encode data enkripsi : %v\n", err)
		return
	}

	decrypt, err := utils.KeyDecrypt(key, encrpytData.Encrypt)
	if err != nil {
		utils.ErrorMessage(c, http.StatusBadRequest, "Ops, Something went wrong")
		fmt.Printf("[login Controller] error when decrypt data enkripsi : %v\n", err)
		return
	}

	admin := models.User{}

	err = json.Unmarshal([]byte(decrypt), &admin)
	if err != nil {
		utils.ErrorMessage(c, http.StatusInternalServerError, "Ops, Something went wrong")
		log.Printf("[login Controller] error when decrypt data enkripsi to struct : %v\n", err)
		return
	}

	result, errs := a.adminService.LoginService(admin.Username, admin.Password)
	if errs != nil || result == "" {
		utils.ErrorMessage(c, http.StatusBadRequest, "invalid username or password")
		return
	}

	utils.SuccessData(c, http.StatusOK, result)
}

func (a *AdminController) logout(c *gin.Context) {
	au, err := auth.ExtractTokenAuth(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	deleted, errs := a.adminService.DeleteAuthData(au.AuthUUID)
	if errs != nil || deleted != 0 { //if any goes wrong
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	utils.SuccessMessage(c, http.StatusOK, "Successfully logged out")
}

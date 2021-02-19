package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../model"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Initlize Database
//var db = database.Conn()

func UserLogin(c *gin.Context) {
	var login model.UserLogin
	RequestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Invalid Parameter"})
		return
	}
	json.Unmarshal(RequestBody, &login)
	fmt.Println("Email :: ", login.Email)
	fmt.Println("Password :: ", login.Password)
	result, err := db.Query("select * FROM user_record where email = ? AND password = ? ", login.Email, login.Password)
	defer result.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid Credentials",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Succesfull",
	})

	return
}

package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"../database"
	"../model"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// variable
var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

// Initlize Database
var db = database.Conn()

func UserRegistration(c *gin.Context) {
	var user model.UserRegistration
	RequestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Invalid Parameter"})
		return
	}
	json.Unmarshal(RequestBody, &user)
	fmt.Print("User Data :: ", user)
	user.Password = generatePassword()
	result, err := db.ExecContext(c, "insert into user_record (name,email,mobile,password) values (?,?,?,?)", user.Name, user.Email, user.Mobile, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Email Id already exists"})
		return
	}
	result.RowsAffected()
	userDetails := map[string]string{
		"username": user.Email,
		"password": user.Password,
	}
	c.JSON(http.StatusCreated, gin.H{
		"user credentials": userDetails,
		"message":          "User Created",
	})
	return
}

func generatePassword() string {
	rand.Seed(time.Now().Unix())
	minSpecialChar := 1
	minNum := 6
	minUpperCase := 1
	passwordLength := 8

	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

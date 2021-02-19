package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"../model"
	"github.com/gin-gonic/gin"
)

func VehicleRegistration(c *gin.Context) {
	var vehicle model.VehicleRecord
	RequestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Invalid Parameter"})
		return
	}
	json.Unmarshal(RequestBody, &vehicle)
	t := time.Now()
	t.String()
	vehicle.ParkIntime = t.Format("2006-01-02 15:04:05")
	vehicle.TokenNumber = generateToken()
	vehicle.ParkOuttime = ""
	vehicle.Day = ""
	vehicle.Rent = ""
	fmt.Print("Vehicle Data :: ", vehicle)
	result, err := db.ExecContext(c, "insert into vehicle_record (name,vehiclename,vehiclenumber,parkingtime,tokennumber,agentid,parkingoutime,day,rent) values (?,?,?,?,?,?,?,?,?)", vehicle.Name, vehicle.Vehiclename, vehicle.Vehiclenumber, vehicle.ParkIntime, vehicle.TokenNumber, vehicle.Agentid, vehicle.ParkOuttime, vehicle.Day, vehicle.Rent)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error"})
		return
	}
	result.RowsAffected()
	c.JSON(http.StatusCreated, gin.H{
		"Vehicle Details": vehicle,
		"Token Number":    vehicle.TokenNumber,
		"message":         "Your Vehicle Parked",
	})
	return
}

func generateToken() string {
	rand.Seed(time.Now().Unix())
	minNum := 6

	var token strings.Builder

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		token.WriteString(string(numberSet[random]))
	}
	inRune := []rune(token.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return "PMS" + string(inRune)
}

func VehicleDetails(c *gin.Context) {
	var vehicle model.VehicleDetails
	var vehicledata []model.VehicleDetails
	// Vehicle Number
	vehiclenumber := c.Param("vehiclenumber")
	log.Println("Id is : ", vehiclenumber)
	result, err := db.Query("select name,vehiclename,vehiclenumber, parkingtime,parkingoutime,tokennumber,day,rent from vehicle_record where vehiclenumber = ?", vehiclenumber)
	defer result.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server error"})
		return
	}
	for result.Next() {
		err := result.Scan(&vehicle.Name, &vehicle.Vehiclename, &vehicle.Vehiclenumber, &vehicle.ParkIntime, &vehicle.ParkOuttime, &vehicle.TokenNumber, &vehicle.Day, &vehicle.Rent)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal Server error"})
			return
		}
		vehicledata = append(vehicledata, vehicle)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": vehicledata,
	})
	return
}

func VehicleParkOut(c *gin.Context) {
	var vehicle model.VehicleParkOut
	// Vehicle Number
	tokennumber := c.Param("tokennumber")
	log.Println("Token Number is : ", tokennumber)
	t := time.Now()
	t.String()
	vehicle.ParkOuttime = t.Format("2006-01-02 15:04:05")
	vehicle.Day = "1"
	vehicle.Rent = "50"
	result, err := db.ExecContext(c, "update vehicle_record set parkingoutime = ?, day = ? , rent = ? where tokennumber = ? ", vehicle.ParkOuttime, vehicle.Day, vehicle.Rent, tokennumber)
	rows, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Invalid Parameter"})
		return
	}
	if rows != 1 {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found"})
		return
	}

	data, err := db.Query("select *from vehicle_record where tokennumber = ?", tokennumber)
	defer data.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server error"})
		return
	}
	for data.Next() {
		var vehicle model.VehicleRecord
		err := data.Scan(&vehicle.ID, &vehicle.Name, &vehicle.Vehiclename, &vehicle.ParkIntime, &vehicle.ParkOuttime, &vehicle.Agentid, &vehicle.TokenNumber, &vehicle.Vehiclenumber, &vehicle.Day, &vehicle.Rent)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal Server error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Vehicle Data": vehicle,
			"message":      "You Take Out your vehicle form parking",
		})
	}
	return
}

func AllVehicle(c *gin.Context) {
	var vehicle model.VehicleRecord
	var vehicledata []model.VehicleRecord
	result, err := db.Query("select *from vehicle_record")
	defer result.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server error"})
		return
	}
	for result.Next() {
		err := result.Scan(&vehicle.ID, &vehicle.Name, &vehicle.Vehiclename, &vehicle.ParkIntime, &vehicle.ParkOuttime, &vehicle.Agentid, &vehicle.TokenNumber, &vehicle.Vehiclenumber, &vehicle.Day, &vehicle.Rent)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Server error"})
			return
		}
		vehicledata = append(vehicledata, vehicle)
	}
	c.JSON(http.StatusOK, gin.H{
		"vehicle_data": vehicledata,
	})
	return
}

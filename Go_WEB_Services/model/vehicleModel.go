package model

type VehicleRecord struct {
	ID            int    `json:"-"`
	Name          string `json:"name"`
	Vehiclename   string `json:"vehiclename"`
	Vehiclenumber string `json:"vehiclenumber"`
	ParkIntime    string `json:"parkintime"`
	ParkOuttime   string `json:"ParkOuttime"`
	Agentid       string `json:"agentid"`
	TokenNumber   string `json:"tokennumber"`
	Day           string `json:"day"`
	Rent          string `json:"rent"`
}

type VehicleDetails struct {
	Name          string `json:"name"`
	Vehiclename   string `json:"vehiclename"`
	Vehiclenumber string `json:"vehiclenumber"`
	ParkIntime    string `json:"parkintime"`
	ParkOuttime   string `json:"parkouttime"`
	TokenNumber   string `json:"tokennumber"`
	Day           string `json:"day"`
	Rent          string `json:"rent"`
}

type VehicleParkOut struct {
	TokenNumber string `json:"tokennumber"`
	ParkOuttime string `json:"parkouttime"`
	Day         string `json:"day"`
	Rent        string `json:"rent"`
}

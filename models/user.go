package models

// Creo mis modelos para los datos

// Este es un struct Adress que va a ir dentro del struc principal User

type Address struct {
	State      string `json:"state" bson:"state"`
	City       string `json:"city" bson:"city"`
	PostalCode int    `json:"postalcode" bson:"postalcode"`
}

// Struct principal

type User struct {
	Name    string  `json:"name" bson:"user_name"`
	Surname string  `json:"surname" bson:"user_surname"`
	Dni     string  `json:"dni" bson:"user_dni"`
	Age     int     `json:"age" bson:"user_age"`
	Addrees Address `json:"address" bson:"user_address"`
}

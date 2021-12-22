package main

import (
	"github.com/gin-gonic/gin"
	"encoding/json"
)

type Products struct {
	Id int
	Nombre string
	Color string
	Precio float64
	Stock int
	Codigo string
	Publicado bool
}

func main(){

	// Crea un router con gin
	router := gin.Default()
	// Captura la solicitud GET “/hello-world”
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
		"message": "Hello Victor",
		})
	})

	router.GET("/products", GetAll)

	// Corremos nuestr
	router.Run()
}

func GetAll(c *gin.Context){
	p1 := Products{
		Id: 1,
		Nombre: "remera",
		Color: "rojo",
    	Precio: 10.00,
        Stock : 5,
        Codigo: "A30",
        Publicado: true,
	}

	p2 := Products{
		Id: 2,
		Nombre: "pantalon",
		Color: "azul",
    	Precio: 14.00,
        Stock : 10,
        Codigo: "A10",
        Publicado: false,
	}

	producList := []Products{}
	producList = append(producList,p1,p2)
	
	jsonData, _ := json.Marshal(producList)
	c.JSON(200, string(jsonData))
}
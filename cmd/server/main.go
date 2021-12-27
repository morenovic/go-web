package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Go-web-c1/practicaTM/go-web/internal/productos"
	"github.com/Go-web-c1/practicaTM/go-web/cmd/server/handler"
)

func main() {
	repo := productos.NewRepository()
	service := productos.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()

	pr := r.Group("/productos")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateName())
	pr.DELETE("/:id", p.Delete())

	r.Run()

}

//go run cmd/server/main.go


/* type Products struct {
	Id int 
	Nombre string
	Color string
	Precio float64
	Stock int
	Codigo string
	Publicado bool
}

type request struct {
	Id int `json:"id"`
	Nombre string `json:"nombre" binding:"required"`
	Color string `json: "color" binding:"required"`
	Precio float64 `json:"precio" binding:"required"`
	Stock int `json:"stock" binding:"required"`
	Codigo string `json:"codigo" binding:"required"`
	Publicado bool `json:"publicado" binding:"required"`
}

var lastId int
var productos []request

func main(){

	// Crea un router con gin
	router := gin.Default()
	// Captura la solicitud GET “/hello-world”
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
		"message": "Hello Victor",
		})
	})

	
	pr := router.Group("/products")
	
	pr.POST("/", Auth, Guardar)
	pr.GET("/", Auth, GetAll)

	// Corremos nuestr
	router.Run()
}


func Auth(c *gin.Context){
	token := c.GetHeader("token")

	if token != "123456" {
		c.JSON(400, gin.H{
			"error": "token invalido",
		})
		return
	}

}


func Guardar(c *gin.Context) {
	var req request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
		"error": err.Error(),
		})
		return
	}
	lastId = len(productos)
	req.Id = lastId +1
	productos = append(productos, req)
	c.JSON(200, req)
}


func GetAll(c *gin.Context){
	
	queries := c.Request.URL.Query()

	fmt.Println(queries)	

	
	c.JSON(200, productos)
}*/
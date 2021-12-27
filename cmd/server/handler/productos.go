package handler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Go-web-c1/practicaTM/go-web/internal/productos"
)

type request struct {
	Id int `json:"id"`
	Nombre string `json:"nombre" binding:"required"`
	Color string `json: "color" binding:"required"`
	Precio float64 `json:"precio" binding:"required"`
	Stock int `json:"stock" binding:"required"`
	Codigo string `json:"codigo" binding:"required"`
	Publicado bool `json:"publicado" binding:"required"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")

		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")

		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		p, err := c.service.Store(req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//Validar token
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		//Parsear id string a int
		Id, err := strconv.ParseInt(ctx.Param("Id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		//Validaciones

		if req.Nombre == "" {
			ctx.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
			return
		}
		if req.Color == "" {
			ctx.JSON(400, gin.H{"error": "El Color del producto es requerido"})
			return
		}
		if req.Precio == 0 {
			ctx.JSON(400, gin.H{"error": "La Precio es requerida"})
			return
		}
		if req.Stock == 0 {
			ctx.JSON(400, gin.H{"error": "El Stock es requerido"})
			return
		}
		if req.Codigo == 0 {
			ctx.JSON(400, gin.H{"error": "El Codigo es requerido"})
			return
		}
		if req.Publicado == 0 {
			ctx.JSON(400, gin.H{"error": "El campo Publicado es requerido"})
			return
		}

		//Hago el Update
		p, err := c.service.Update(int(Id), req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *Product) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//Validar token
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		//Parsear id string a int
		Id, err := strconv.ParseInt(ctx.Param("Id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		//Validaciones
		if req.Nombre == "" {
			ctx.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
			return
		}

		//Hago el patch del producto
		p, err := c.service.UpdateName(int(Id), req.Nombre)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//Validar token
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		//Parsear id string a int
		Id, err := strconv.ParseInt(ctx.Param("Id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		//Hago la baja fisica del producto
		err = c.service.Delete(int(Id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"data": fmt.Sprintf("El producto %d ha sido eliminado", Id)})
	}
}

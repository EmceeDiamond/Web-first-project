package main

import (
	"net/http"

	"API/models"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	//"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	router := gin.Default()
	/*router.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))*/

	//router.LoadHTMLGlob("template/*")
	router.Use(cors.Default())
	router.GET("/providers", getProviders)

	router.GET("/products", getProducts)

	router.GET("/provider/:provider_name", getProvider)

	router.POST("/add_providers", addProvider)

	router.POST("/add_products", addProduct)

	router.Run("localhost:8000")
}

func getProviders(c *gin.Context) {

	providers := models.GetProviders()
	if providers == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, providers)
	}
}

func getProvider(c *gin.Context) {

	provider_name := c.Param("provider_name")
	provider := models.GetProvider(provider_name)
	if provider == nil {
		c.AbortWithStatus(http.StatusNotFound)

	} else {

		c.IndentedJSON(http.StatusOK, provider)

	}

}

func addProvider(c *gin.Context) {
	var prov models.Provider

	if err := c.BindJSON(&prov); err != nil {

		c.AbortWithStatus(http.StatusBadRequest)
	} else {

		models.AddProvider(prov)
		c.IndentedJSON(http.StatusCreated, prov)
	}

}

func addProduct(c *gin.Context) {
	var prod models.Product

	if err := c.BindJSON(&prod); err != nil {

		c.AbortWithStatus(http.StatusBadRequest)
	} else {

		models.AddProduct(prod)
		c.IndentedJSON(http.StatusCreated, prod)
	}

}

func getProducts(c *gin.Context) {

	products := models.GetProducts()
	if products == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, products)
	}
}

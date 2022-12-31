package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/customers", getCustomers)
	router.GET("/customers/:id", getCustomersById)
	router.POST("/customers", postCustomers)

	router.Run("localhost:8080")
}

type customer struct {
	ID		string	`json:"id"`
	Name	string	`json:"name"`
}

var customers = []customer{
	{ID: "0001", Name: "Pedro"},
	{ID: "0002", Name: "José"},
	{ID: "0003", Name: "Paulo"},
}

func getCustomers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, customers)
}

func postCustomers(c *gin.Context) {
	var newCustomer customer

	if err := c.BindJSON(&newCustomer); err != nil {
		return
	}

	customers = append(customers, newCustomer)
	c.IndentedJSON(http.StatusCreated, newCustomer)
}

func getCustomersById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range customers {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "customer not found"})
}
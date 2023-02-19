package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Person struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

var people = []Person{
	{ID:1, Name:"Ava Max", Age:29},
	{ID:2, Name:"Bebe Rexha", Age:33},
	{ID:3, Name:"Charlie Puth", Age:31},
}

func getPeople(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, people)
}

func createPeople(c *gin.Context) {
	var newPerson Person

	if err:= c.BindJSON(&newPerson); err != nil {
		return
	}

	people = append(people, newPerson)
	c.IndentedJSON(http.StatusCreated, newPerson)
}

func main() {
	r := gin.Default()
	r.GET("/person", getPeople)
	r.POST("/person", createPeople)
	r.Run()
}
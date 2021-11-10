package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type UserDetails struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Date      string `json:"date"`
	Email     string `json:"email"`
	Income    int    `json:"income"`
	IpAddress string `json:"ipaddress"`
}

var data []UserDetails

func postUser(c *gin.Context) {

	var user UserDetails

	err := c.BindJSON(&user)
	if err != nil {
		panic(err)
	}

	jsonfile, err1 := ioutil.ReadFile("customer.json")
	if err1 != nil {
		panic(err1)
	}

	json.Unmarshal(jsonfile, &data)

	data = append(data, user)

	newUser, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("customer.json", newUser, 0664)
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func getUser(c *gin.Context) {

	//Header().Set("Access-Control-Allow-Origin", "*")

	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	jsonfile, err1 := ioutil.ReadFile("customer.json")
	if err1 != nil {
		panic(err1)
	}

	json.Unmarshal(jsonfile, &data)

	c.JSON(200, data)

}

func deleteUser(c *gin.Context) {

	id, err1 := strconv.Atoi(c.Param("id"))
	if err1 != nil {
		panic(err1)
	}

	jsonfile, err := ioutil.ReadFile("customer.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(jsonfile, &data)

	for key, user := range data {

		if id == user.ID {

			data = append(data[:key], data[key+1:]...)

			newUser, err := json.MarshalIndent(data, "", "\t")
			if err != nil {
				panic(err)
			}

			ioutil.WriteFile("customer.json", newUser, 0664)
			c.JSON(http.StatusOK, gin.H{"message": "deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "not found"})

}

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	router.POST("/api/postData/", postUser)
	router.GET("/api/getData/", getUser)
	router.DELETE("/api/deleteRow/:id", deleteUser)
	router.Run("localhost:8080")
}

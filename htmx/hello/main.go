package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type Response struct {
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	Data       []User `json:"data"`
}

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("./templates/**/*")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{})
	})

	r.GET("/admin", func(c *gin.Context) {
		url := "https://reqres.in/api/users"
		res, _ := http.Get(url)
		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err.Error())
		}
		// fmt.Println(string(body))
		var data Response
		err = json.Unmarshal(body, &data)
		c.HTML(http.StatusOK, "admin.html", gin.H{
			"data": data.Data,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

package handlers

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"

	"github.com/dolldot/scrapyard/bookmarkd/internal/repos"
	"github.com/dolldot/scrapyard/bookmarkd/views/components/account"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	repo repos.AccountRepo
}

func NewAccountHandler(repo repos.AccountRepo) *AccountHandler {
	return &AccountHandler{
		repo: repo,
	}
}

func generateRandomNumber() string {
	var number string
	for i := 0; i < 12; i++ {
		digit, _ := rand.Int(rand.Reader, big.NewInt(10)) // Generate a random digit between 0 and 9.
		number += digit.String()
	}
	return number
}

func (handler AccountHandler) CreateAccount(c *gin.Context) {
	accountNumber := generateRandomNumber()

	data, err := handler.repo.CreateAccount(c, accountNumber)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "error creating an account",
		})
		return
	}

	c.Writer.Header().Add("HX-Redirect", fmt.Sprintf("/%s/bookmark", data.Number))
	c.Status(http.StatusCreated)
}

func (handler AccountHandler) RenderLogin(c *gin.Context) {
	render(c, http.StatusOK, components.Login())
}

func (handler AccountHandler) Login(c *gin.Context) {
	accountNumber := c.PostForm("account")

	data, err := handler.repo.FindAccountByNumber(c, accountNumber)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "account not found",
		})
		return
	}

	c.Writer.Header().Add("HX-Redirect", fmt.Sprintf("/%s/bookmark", data.Number))
	c.Status(http.StatusOK)
}

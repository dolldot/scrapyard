package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/dolldot/scrapyard/bookmarkd/generated/ent"
	"github.com/dolldot/scrapyard/bookmarkd/internal/repos"
	"github.com/dolldot/scrapyard/bookmarkd/views/components/bookmark"
	"github.com/dolldot/scrapyard/bookmarkd/views/layout"
	"github.com/dolldot/scrapyard/bookmarkd/views/pages"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookmarkHandler struct {
	repo repos.BookmarkRepo
}

func NewBookmarkHandler(repo repos.BookmarkRepo) *BookmarkHandler {
	return &BookmarkHandler{
		repo: repo,
	}
}

func render(ctx *gin.Context, status int, template templ.Component) error {
	ctx.Status(status)
	return template.Render(ctx.Request.Context(), ctx.Writer)
}

func (handler BookmarkHandler) GetBookmarks(c *gin.Context) {
	number := c.Param("number")
	bookmarks, err := handler.repo.FindAll(c, number)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "error get posts",
		})
		return
	}

	data := []*BookmarkReponse{}
	for _, bookmark := range bookmarks {
		data = append(data, bookmarkTransformer(bookmark))
	}

	render(c, http.StatusOK, layout.Content("Bookmark List", number, pages.BookmarkList(number, bookmarks)))
}

func (handler BookmarkHandler) AddBookmark(c *gin.Context) {
	number := c.Param("number")
	render(c, http.StatusOK, components.BookmarkAdd(number))
}

func (handler BookmarkHandler) CreateBookmark(c *gin.Context) {
	number := c.Param("number")

	fmt.Println(number)

	var bookmark ent.Bookmark
	bookmark.Name = c.PostForm("name")
	bookmark.URL = c.PostForm("url")
	bookmark.AccountNumber = number

	_, err := handler.repo.Create(c, bookmark)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "error creating bookmark",
		})
		return
	}

	c.Writer.Header().Add("HX-Refresh", "true")
	c.Status(http.StatusCreated)
}

func (handler BookmarkHandler) EditBookmark(c *gin.Context) {
	number := c.Param("number")
	uuid, _ := uuid.Parse(c.Param("uuid"))

	data, err := handler.repo.FindByUUID(c, number, uuid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "error get bookmark",
		})
		return
	}

	render(c, http.StatusOK, components.BookmarkEdit(data))
}

func (handler BookmarkHandler) UpdateBookmark(c *gin.Context) {
	number := c.Param("number")
	uuid, _ := uuid.Parse(c.Param("uuid"))

	var bookmark ent.Bookmark
	bookmark.Name = c.PostForm("name")
	bookmark.URL = c.PostForm("url")

	err := handler.repo.UpdateByUUID(c, number, uuid, bookmark)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "error updating bookmark",
		})
		return
	}

	c.Writer.Header().Add("HX-Refresh", "true")
	c.Status(http.StatusOK)
}

func (handler BookmarkHandler) DeleteBookmark(c *gin.Context) {
	number := c.Param("number")
	uuid, _ := uuid.Parse(c.Param("uuid"))

	err := handler.repo.DeleteByUUID(c, number, uuid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "error deleting bookmark",
		})
		return
	}

	c.Writer.Header().Add("HX-Refresh", "true")
	c.Status(http.StatusOK)
}

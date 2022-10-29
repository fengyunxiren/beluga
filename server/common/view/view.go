package view

import (
	"beluga/global"
	"beluga/server/common/response"
	"errors"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QueryBind struct {
	Page     int    `form:"page" binding:"omitempty,gte=1"`
	PageSize int    `form:"page_size" binding:"omitempty,gte=10"`
	Query    string `form:"q"`
}
type CommonView struct {
	EndPoint  string
	UrlPrefix string
}

type IDBind struct {
	ID string `uri:"id" binding:"required"`
}

func (v CommonView) CreateAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		response.ResponseWithStatus(
			http.StatusMethodNotAllowed,
			response.ERR_NOT_ALLOWED,
			nil,
			c,
		)
	}
}

func (v CommonView) ListAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		response.ResponseWithStatus(
			http.StatusMethodNotAllowed,
			response.ERR_NOT_ALLOWED,
			nil,
			c,
		)
	}
}

func (v CommonView) UpdateAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		response.ResponseWithStatus(
			http.StatusMethodNotAllowed,
			response.ERR_NOT_ALLOWED,
			nil,
			c,
		)
	}
}

func (v CommonView) DeleteAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		response.ResponseWithStatus(
			http.StatusMethodNotAllowed,
			response.ERR_NOT_ALLOWED,
			nil,
			c,
		)
	}
}

func (v CommonView) GetDetailAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		response.ResponseWithStatus(
			http.StatusMethodNotAllowed,
			response.ERR_NOT_ALLOWED,
			nil,
			c,
		)
	}
}

func (v CommonView) GetDB(c *gin.Context) (*gorm.DB, error) {
	db, exist := c.Get(global.CONTEXT_DB)
	if !exist {
		return nil, errors.New("db connect not exist")
	}
	switch db := db.(type) {
	case *gorm.DB:
		//新增操作
		return db, nil
	default:
		return nil, errors.New("db connect not support")
	}
}

func (v CommonView) GetViewPath() string {
	uri := filepath.Join(v.UrlPrefix, v.EndPoint)
	return uri
}

func RegisterView(view View, router gin.IRouter) {
	uri := view.GetViewPath()
	router.GET(uri, view.ListAction())
	router.POST(uri, view.CreateAction())
	individualUri := filepath.Join(uri, ":id")
	router.GET(individualUri, view.GetDetailAction())
	router.PUT(individualUri, view.UpdateAction())
	router.DELETE(individualUri, view.DeleteAction())
}

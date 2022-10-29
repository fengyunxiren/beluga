package view

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type View interface {
	GetViewPath() string
	CreateAction() gin.HandlerFunc
	ListAction() gin.HandlerFunc
	UpdateAction() gin.HandlerFunc
	DeleteAction() gin.HandlerFunc
	GetDetailAction() gin.HandlerFunc
	GetDB(c *gin.Context) (*gorm.DB, error)
}

package view

import (
	"beluga/server/common/database"
	"beluga/server/common/logger"
	"beluga/server/common/response"
	"beluga/utils"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type DBView struct {
	CommonView
	generator utils.Generator
}

func NewDBView(endPoint string, UrlPrefix string, generator utils.Generator) DBView {
	return DBView{
		CommonView: CommonView{
			EndPoint:  endPoint,
			UrlPrefix: UrlPrefix,
		},
		generator: generator,
	}
}

func (v DBView) CreateAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.GetContextLogger(c)
		db, err := v.GetDB(c)
		if err != nil {
			log.Error("Get db Failed: ", err)
			response.AbortWithError(response.ERR_SERVER_DB_NOT_FOUND, c)
			return
		}
		instance := v.generator(utils.PTR)
		log.Info("type of instance: ", reflect.TypeOf(instance))
		if err = c.ShouldBindBodyWith(instance, binding.JSON); err != nil {
			log.Error("Bind body failed: ", err)
			response.AbortWithError(response.ERR_BAD_REQUEST, c)
			return
		}
		log.Info("instance: ", instance)
		log.Info("type of instance: ", reflect.TypeOf(instance))
		result := db.Create(instance)
		if result.Error != nil {
			log.Error("Create failed: ", result.Error)
			response.ResponseError(response.ERR_SERVER_VIEW_CREATE, c)
		} else {
			response.ResponseOkWithData(instance, c)
		}
	}
}

func (v DBView) ListAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.GetContextLogger(c)
		db, err := v.GetDB(c)
		if err != nil {
			log.Error("Get db Failed: ", err)
			response.AbortWithError(response.ERR_SERVER_DB_NOT_FOUND, c)
			return
		}
		query := QueryBind{}
		if err = c.ShouldBindQuery(&query); err != nil {
			log.Error("Query parameter illegal: ", err)
			response.ResponseError(response.ERR_BAD_REQUEST, c)
			return
		}
		if query.Page == 0 {
			query.Page = 1
		}
		if query.PageSize == 0 {
			query.PageSize = 10
		}
		log.Info("query: ", query)
		model := v.generator(utils.PTR)
		instances := v.generator(utils.LIST)
		var count int64
		offset := (query.Page - 1) * query.PageSize
		db = database.FuzzyQuery(db, model, query.Query, nil)
		db.Model(model).Count(&count)
		db.Model(model).Offset(offset).Limit(query.PageSize).Find(&instances)
		response.ResponseOkWithData(gin.H{
			"pagination": gin.H{
				"page":      query.Page,
				"page_size": query.PageSize,
				"total":     count,
			},
			"datas": instances,
		}, c)
	}
}

func (v DBView) GetDetailAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.GetContextLogger(c)
		db, err := v.GetDB(c)
		if err != nil {
			log.Error("Get db Failed: ", err)
			response.AbortWithError(response.ERR_SERVER_DB_NOT_FOUND, c)
			return
		}
		id := IDBind{}
		if err := c.ShouldBindUri(&id); err != nil {
			log.Error("Uri parameter illegal: ", err)
			response.ResponseError(response.ERR_BAD_REQUEST, c)
			return
		}
		log.Info("id: ", id)
		model := v.generator(utils.PTR)
		instances := v.generator(utils.LIST)
		result := db.Model(model).Where("ID = ?", id.ID).First(&instances)
		if result.Error != nil {
			log.Error("Create failed: ", result.Error)
			response.ResponseError(response.ERR_SERVER_VIEW_GET, c)
		} else {
			res, err := utils.InterfaceToArray(instances)
			if err != nil || len(res) == 0 {
				response.ResponseOkWithData(instances, c)
			} else {
				response.ResponseOkWithData(res[0], c)
			}
		}
	}
}

func (v DBView) UpdateAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.GetContextLogger(c)
		db, err := v.GetDB(c)
		if err != nil {
			log.Error("Get db Failed: ", err)
			response.AbortWithError(response.ERR_SERVER_DB_NOT_FOUND, c)
			return
		}
		id := IDBind{}
		if err := c.ShouldBindUri(&id); err != nil {
			log.Error("Uri parameter illegal: ", err)
			response.ResponseError(response.ERR_BAD_REQUEST, c)
			return
		}
		model := v.generator(utils.PTR)
		update := v.generator(utils.PTR)
		if err = c.ShouldBindBodyWith(update, binding.JSON); err != nil {
			log.Error("Bind body failed: ", err)
			response.AbortWithError(response.ERR_BAD_REQUEST, c)
			return
		}
		log.Info("update: ", update)
		result := db.Model(model).Where("id = ?", id.ID).Updates(update)
		if result.Error != nil {
			log.Error("Update model failed: ", result.Error)
			response.ResponseError(response.ERR_SERVER_VIEW_UPDATE, c)
		} else {
			response.ResponseOk(c)
		}
	}
}

func (v DBView) DeleteAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.GetContextLogger(c)
		db, err := v.GetDB(c)
		if err != nil {
			log.Error("Get db Failed: ", err)
			response.AbortWithError(response.ERR_SERVER_DB_NOT_FOUND, c)
			return
		}
		id := IDBind{}
		if err := c.ShouldBindUri(&id); err != nil {
			log.Error("Uri parameter illegal: ", err)
			response.ResponseError(response.ERR_BAD_REQUEST, c)
			return
		}
		model := v.generator(utils.PTR)
		instances := v.generator(utils.LIST)
		result := db.Model(model).Where("id = ?", id.ID).Delete(&instances)
		if result.Error != nil {
			log.Error("Delete model failed: ", result.Error)
			response.ResponseError(response.ERR_SERVER_VIEW_DELETE, c)
		} else {
			response.ResponseOk(c)
		}
	}
}

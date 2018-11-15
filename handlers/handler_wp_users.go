package handlers

import (
	"github.com/gin-gonic/gin"
	"mojotv/zerg/models"
)

func init() {
	groupApi.GET("wp-user", wpUserAll)
	groupApi.GET("wp-user/:id", wpUserOne)
	groupApi.POST("wp-user", wpUserCreate)
	groupApi.PATCH("wp-user", wpUserUpdate)
	groupApi.DELETE("wp-user/:id", wpUserDelete)
}

func wpUserAll(c *gin.Context) {
	mdl := models.WpUser{}
	query := &models.PaginationQuery{}
	err := c.ShouldBindQuery(query)
	if handleError(c, err) {
		return
	}
	list, total, err := mdl.All(query)
	if handleError(c, err) {
		return
	}
	jsonPagination(c, list, total, query)
}
func wpUserOne(c *gin.Context) {
	var mdl models.WpUser
	id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}
	mdl.Id = id
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}
func wpUserCreate(c *gin.Context) {
	var mdl models.WpUser
	err := c.ShouldBind(&mdl)
	if handleError(c, err) {
		return
	}
	err = mdl.Create()
	if handleError(c, err) {
		return
	}
	jsonData(c, mdl)
}

func wpUserUpdate(c *gin.Context) {
	var mdl models.WpUser
	err := c.ShouldBind(&mdl)
	if handleError(c, err) {
		return
	}
	err = mdl.Update()
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)
}

func wpUserDelete(c *gin.Context) {
	var mdl models.WpUser
	id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}
	mdl.Id = id
	err = mdl.Delete()
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)
}

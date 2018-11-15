package handlers

import (
	"github.com/gin-gonic/gin"
	"mojotv/zerg/models"
)

func init() {
	groupApi.GET("wp-post", wpPostAll)
	groupApi.GET("wp-post/:id", wpPostOne)
	groupApi.POST("wp-post", wpPostCreate)
	groupApi.PATCH("wp-post", wpPostUpdate)
	groupApi.DELETE("wp-post/:id", wpPostDelete)
}

func wpPostAll(c *gin.Context) {
	mdl := models.WpPost{}
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
func wpPostOne(c *gin.Context) {
	var mdl models.WpPost
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
func wpPostCreate(c *gin.Context) {
	var mdl models.WpPost
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

func wpPostUpdate(c *gin.Context) {
	var mdl models.WpPost
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

func wpPostDelete(c *gin.Context) {
	var mdl models.WpPost
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

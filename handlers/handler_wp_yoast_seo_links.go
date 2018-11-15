package handlers

import (
	"github.com/gin-gonic/gin"
	"mojotv/zerg/models"
)

func init() {
	groupApi.GET("wp-yoast-seo-link", wpYoastSeoLinkAll)
	groupApi.GET("wp-yoast-seo-link/:id", wpYoastSeoLinkOne)
	groupApi.POST("wp-yoast-seo-link", wpYoastSeoLinkCreate)
	groupApi.PATCH("wp-yoast-seo-link", wpYoastSeoLinkUpdate)
	groupApi.DELETE("wp-yoast-seo-link/:id", wpYoastSeoLinkDelete)
}

func wpYoastSeoLinkAll(c *gin.Context) {
	mdl := models.WpYoastSeoLink{}
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
func wpYoastSeoLinkOne(c *gin.Context) {
	var mdl models.WpYoastSeoLink
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
func wpYoastSeoLinkCreate(c *gin.Context) {
	var mdl models.WpYoastSeoLink
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

func wpYoastSeoLinkUpdate(c *gin.Context) {
	var mdl models.WpYoastSeoLink
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

func wpYoastSeoLinkDelete(c *gin.Context) {
	var mdl models.WpYoastSeoLink
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

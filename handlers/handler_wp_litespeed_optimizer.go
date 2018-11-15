package handlers

import (
	"github.com/gin-gonic/gin"
	"mojotv/zerg/models"
)

func init() {
	groupApi.GET("wp-litespeed-optimizer", wpLitespeedOptimizerAll)
	groupApi.GET("wp-litespeed-optimizer/:id", wpLitespeedOptimizerOne)
	groupApi.POST("wp-litespeed-optimizer", wpLitespeedOptimizerCreate)
	groupApi.PATCH("wp-litespeed-optimizer", wpLitespeedOptimizerUpdate)
	groupApi.DELETE("wp-litespeed-optimizer/:id", wpLitespeedOptimizerDelete)
}

func wpLitespeedOptimizerAll(c *gin.Context) {
	mdl := models.WpLitespeedOptimizer{}
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
func wpLitespeedOptimizerOne(c *gin.Context) {
	var mdl models.WpLitespeedOptimizer
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
func wpLitespeedOptimizerCreate(c *gin.Context) {
	var mdl models.WpLitespeedOptimizer
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

func wpLitespeedOptimizerUpdate(c *gin.Context) {
	var mdl models.WpLitespeedOptimizer
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

func wpLitespeedOptimizerDelete(c *gin.Context) {
	var mdl models.WpLitespeedOptimizer
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

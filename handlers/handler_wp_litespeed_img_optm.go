package handlers

import (
	"github.com/gin-gonic/gin"
	"mojotv/zerg/models"
)

func init() {
	groupApi.GET("wp-litespeed-img-optm", wpLitespeedImgOptmAll)
	groupApi.GET("wp-litespeed-img-optm/:id", wpLitespeedImgOptmOne)
	groupApi.POST("wp-litespeed-img-optm", wpLitespeedImgOptmCreate)
	groupApi.PATCH("wp-litespeed-img-optm", wpLitespeedImgOptmUpdate)
	groupApi.DELETE("wp-litespeed-img-optm/:id", wpLitespeedImgOptmDelete)
}

func wpLitespeedImgOptmAll(c *gin.Context) {
	mdl := models.WpLitespeedImgOptm{}
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
func wpLitespeedImgOptmOne(c *gin.Context) {
	var mdl models.WpLitespeedImgOptm
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
func wpLitespeedImgOptmCreate(c *gin.Context) {
	var mdl models.WpLitespeedImgOptm
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

func wpLitespeedImgOptmUpdate(c *gin.Context) {
	var mdl models.WpLitespeedImgOptm
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

func wpLitespeedImgOptmDelete(c *gin.Context) {
	var mdl models.WpLitespeedImgOptm
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

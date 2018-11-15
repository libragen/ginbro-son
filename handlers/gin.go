package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"mojotv/zerg/models"
	"path"
	"strconv"
	"time"
)

var r = gin.Default()

var groupApi *gin.RouterGroup

//in the same package init executes in file'name alphabet order
func init() {
	if viper.GetBool("app.enable_cors") {
		enableCors()
	}

	if sp := viper.GetString("app.static_path"); sp != "" {
		r.Use(static.Serve("/", static.LocalFile(sp, true)))
		if viper.GetBool("app.enable_not_found") {
			r.NoRoute(func(c *gin.Context) {
				file := path.Join(sp, "index.html")
				c.File(file)
			})
		}
	}

	if viper.GetBool("app.enable_swagger") && viper.GetString("app.env") != "prod" {
		//add edit your own swagger.doc.yml file in ./swagger/doc.yml
		//generateSwaggerDocJson()
		r.Static("swagger", "./swagger")
	}
	prefix := viper.GetString("app.api_prefix")
	api := "api"
	if prefix != "" {
		api = fmt.Sprintf("%s/%s", api, prefix)
	}
	groupApi = r.Group(api)
	if viper.GetString("app.env") != "prod" {
		r.GET("/app/info", func(c *gin.Context) {
			c.JSON(200, viper.GetStringMapString("app"))
		})
	}
}

func ServerRun() {

	addr := viper.GetString("app.addr")
	if viper.GetBool("app.enable_https") {
		logrus.Fatal(autotls.Run(r, addr))
	} else {
		r.Run(addr)
	}
}

func Close() {
	models.Close()
}

func jsonError(c *gin.Context, msg string) {
	c.JSON(200, gin.H{"code": 0, "msg": msg})
}
func jsonData(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{"code": 1, "data": data})
}
func jsonPagination(c *gin.Context, list interface{}, total uint, query *models.PaginationQuery) {
	c.JSON(200, gin.H{"code": 1, "data": list, "total": total, "offset": query.Offset, "limit": query.Limit})
}
func jsonSuccess(c *gin.Context) {
	c.JSON(200, gin.H{"code": 1, "msg": "success"})
}

func handleError(c *gin.Context, err error) bool {
	if err != nil {
		jsonError(c, err.Error())
		return true
	}
	return false
}

func parseParamID(c *gin.Context) (uint, error) {
	id := c.Param("id")
	if parseId, err := strconv.ParseUint(id, 10, 32); err != nil {
		return 0, errors.New("id must be an unsigned int")
	} else {
		return uint(parseId), nil
	}
}

func enableCors() {
	//https://github.com/gin-contrib/cors

	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, //https://foo.com
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false, //enable cookie
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour, //cache options result decrease request lag
	}))
}

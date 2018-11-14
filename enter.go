package grf

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"net/http"
	"gopkg.in/yaml.v2"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/swaggo/gin-swagger"
)
func Enter(r *gin.Engine,db *gorm.DB) *Register {
	reg := new(Register)
	reg.route = r
	reg.db = db
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/api/doc", func(c *gin.Context) {
		s, err := yaml.Marshal(Swagger)
		if err != nil {
			panic(err)
		}
		c.String(http.StatusOK, string(s))
	})
	return reg
}

type Register struct {
	db *gorm.DB
	route *gin.Engine
}

func (this *Register) Inset(view Viewer, path string) {
	CreateView(this.route,this.db,view, path)
}




# grf
*grf 依赖gin web 还有grom orm*
### 注意： 这是想法实现不要用在线上业务上

### 实现功能
1.不需要注释，自动生成部分swagger文档  
2.实现interface自动生成curd  
3.序列化可以自定义

### 后续功能
1.完善swagger文档生成部分
2.重构代码
3.结合prometheus实现项目依赖图

### 例子
```
package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/gin-gonic/gin"
	"grf"
	"net/http"
	"log"
)
type User struct {
	gorm.Model
	Name string `json:"name"`
	Age int `json:"age"`
}
type UserInterface struct {
	Name string `json:"name" form:"name"`
}

type UserResponses struct {
	Data string `json:"data"`
	Msg string `json:"msg"`
	Status string `json:"status"`
	Gorm []*UserInterface `json:"gorm"`
}

type UserObj struct {
}

func (this *UserObj) GetModelObj() interface{} {
	return new(User)
}

func (this *UserObj) GetModelObjs() interface{} {
	objs := make([]*User, 0, 0)
	return &objs
}

func (this *UserObj) GetInterfaceObj() interface{} {
	return new(UserInterface)
}

func (this *UserObj) GetInterfaceObjs() interface{} {
	objs := make([]*UserInterface, 0, 0)
	return &objs
}

func (this *UserObj) GetResponsesObj() interface{} {
	return new(UserResponses)
}

func (this *UserObj) GetNotes() (notes *grf.Notes)  {
	notes = new(grf.Notes)
	notes.GetNote.Summary = "获取用户"
	notes.GetIdNote.Description = "根据id 获取用户"
	return
}
func (this *UserObj) GetPaths() []string {
	return grf.PassDefault()
}


type UserView struct {
	grf.Serializers
	UserObj
}


var R *gin.Engine
var Db *gorm.DB
var Reg *grf.Register

func main()  {
	R = gin.Default()
	db, err := gorm.Open("sqlite3", "./data.db")
	if err != nil {
		log.Println(err)
	}
	Db = db
	Reg = grf.Enter(R,Db)
	Reg.Inset(&UserView{}, "/user")

	http.ListenAndServe(":10800", R)
}
```

### 生成的文档
![文档生成](https://github.com/fitan/grf/blob/master/readmeimage/swaggerimage.png)


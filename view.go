package grf

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Viewer interface {
	SerializersDataer
	Serializerser
}

type View struct {
	ser Serializerser
	data SerializersDataer
}

func PassDefault() []string {
	return []string{"GET","POST","PUT","DELETE"}
}
func CreateView(route *gin.Engine,db *gorm.DB,view Viewer, path string)  {

	for _, value := range view.GetPaths() {
		switch value {
		case "GET":
			GetAddPaths(path, view.GetInterfaceObj(), view.GetResponsesObj(), view.GetNotes())
			route.GET(path, func(c *gin.Context) {
				interfaceObj := view.GetInterfaceObj()
				modelObjs := view.GetModelObjs()
				if len(c.Request.URL.Query()) != 0  {
					err := c.Bind(interfaceObj)
					if err != nil {
						ErrCheck(err, c)
						return
					}
					data, err := view.Query(db, interfaceObj, modelObjs)
					if err != nil {
						ErrCheck(err, c)
						return
					}
					OkCheck(data,c)
				} else {
					data, err := view.Query(db, nil, modelObjs)
					if err != nil {
						ErrCheck(err, c)
						return
					}
					OkCheck(data,c)
				}
			})
			GetIdPaths(path, view.GetInterfaceObj(), view.GetResponsesObj(),view.GetNotes())
			route.GET(path+"/:id", func(c *gin.Context) {
				id,err := GetId(c)
				if err != nil {
					ErrCheck(err, c)
					return
				}
				obj := view.GetModelObj()
				data, err := view.IdQuery(db,id, obj)
				if err != nil {
					ErrCheck(err, c)
					return
				}
				OkCheck(data, c)
			})
		case "POST":
			PostAddPaths(path, view.GetInterfaceObj(), view.GetResponsesObj(),view.GetNotes())
			route.POST(path, func(c *gin.Context) {
				modelObj := view.GetModelObj()
				err := c.Bind(modelObj)
				if err != nil {
					ErrCheck(err,c)
					return
				}
				interfaceObj := view.GetInterfaceObj()
				err = view.Create(db,interfaceObj, modelObj)
				if err != nil {
					ErrCheck(err, c)
					return
				}
				OkCheck("", c)
			})
		case "PUT":
			PutAddPaths(path, view.GetInterfaceObj(), view.GetResponsesObj(),view.GetNotes())
			route.PUT(path+"/:id", func(c *gin.Context) {
				id, err := GetId(c)
				if err != nil {
					ErrCheck(err, c)
					return
				}
				interfaceObj := view.GetInterfaceObj()
				err = c.Bind(interfaceObj)
				if err != nil {
					ErrCheck(err, c)
					return
				}
				modelObj := view.GetModelObj()
				err = view.UpData(db,id, interfaceObj, modelObj)
				if err != nil {
					ErrCheck(err,c)
				}
				OkCheck("", c)
			})
		case "DELETE":
			DeleteIdPaths(path, view.GetInterfaceObj(), view.GetResponsesObj(),view.GetNotes())
			route.DELETE(path+"/:id", func(c *gin.Context) {
				id, err := GetId(c)
				if err != nil {
					ErrCheck(err, c)
					return
				}
				obj := view.GetModelObj()
				err = view.IdDelete(db,id, obj)
				if err != nil {
					ErrCheck(err,c)
					return
				}
				OkCheck("", c)
			})
		}
	}
}

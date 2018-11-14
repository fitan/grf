package grf

import "github.com/jinzhu/gorm"

type Serializerser interface {
	Create(db *gorm.DB,InterfaceObj interface{},ModelObj interface{}) (err error)
	UpData(db *gorm.DB,id int, InterfaceObj interface{}, ModelObj interface{}) (err error)
	Query(db *gorm.DB,InterfaceObj interface{}, ModelObjs interface{}) (data interface{}, err error)
	IdQuery(db *gorm.DB,id int, ModelObj interface{}) (data interface{}, err error)
	IdDelete(db *gorm.DB,id int, ModelObj interface{}) (err error)
}

type Serializers struct {
}

func (this *Serializers) Create(db *gorm.DB,InterfaceObj interface{},ModelObj interface{}) (err error) {
	err = db.Create(ModelObj).Error
	return
}

func (this *Serializers) Query(db *gorm.DB,InterfaceObj interface{}, ModelObjs interface{}) (data interface{}, err error) {
	if InterfaceObj == nil {
		err = db.Find(ModelObjs).Error
	} else {
		err = db.Where(InterfaceObj).Find(ModelObjs).Error
	}
	data = ModelObjs
	return
}

func (this *Serializers) IdQuery(db *gorm.DB,id int, ModelObj interface{}) (data interface{}, err error) {
	err = db.First(ModelObj, id).Error
	data = ModelObj
	return
}

func (this *Serializers) IdDelete(db *gorm.DB,id int, ModelObj interface{}) (err error) {
	err = db.Where("ID = ?", id).Delete(ModelObj).Error
	return
}
func (this *Serializers ) UpData(db *gorm.DB,id int,InterfaceObj interface{}, ModelObj interface{}) (err error)  {
	err = db.Model(ModelObj).Where("ID = ?", id).Updates(InterfaceObj).Error
	return
}

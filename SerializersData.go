package grf

import "github.com/jinzhu/gorm"

type DefaultModel struct {
	gorm.Model
}

type DefaultInterface struct {
	gorm.Model
}

type DefaultResponse struct {
	gorm.Model
}

type NoteMsg struct {
	Summary string
	Description string
}

type Notes struct {
	GetNote	NoteMsg
	GetIdNote NoteMsg
	PostNote NoteMsg
	PutNote NoteMsg
	DeleteNote NoteMsg
}

type SerializersDataer interface {
	GetModelObj() interface{}
	GetModelObjs() interface{}
	GetInterfaceObj() interface{}
	GetInterfaceObjs() interface{}
	GetResponsesObj() interface{}
	GetNotes() (notes *Notes)
	GetPaths() []string
}

type SerializersData struct {
}

func (this *SerializersData) GetModelObj() interface{} {
	return new(DefaultModel)
}

func (this *SerializersData) GetModelObjs() interface{} {
	objs := make([]*DefaultModel, 0, 0)
	return &objs
	
}

func (this *SerializersData) GetInterfaceObj() interface{} {
	return new(DefaultInterface)
	
}
func (this *SerializersData) GetInterfaceObjs() interface{} {
	objs := make([]*DefaultInterface, 0, 0)
	return &objs
}
func (this *SerializersData) GetResponsesObj() interface{} {
	return new(DefaultResponse)

}

func (this *SerializersData) GetNotes() (notes *Notes)  {
	notes = new(Notes)
	return
}

package grf

import (
	"reflect"
	//"fmt"
	//"log"
	"unicode"
)
var Swagger *DefaultSwagger
func init() {
	Swagger = new(DefaultSwagger)
	Swagger.Swagger = "3.0"
	Swagger.Info.Description = ""
	Swagger.Info.Title = ""
	Swagger.Info.Version = ""
	Swagger.Schemes = append(Swagger.Schemes, []string{"http", "https"}...)
	Swagger.Paths = make(map[string]map[string]methodData)
}


type DefaultSwagger struct {
	Swagger string `yaml:"swagger"`
	Info struct{
		Description string `yaml:"description"`
		Title string `yaml:"title"`
		Version string `yaml:"version"`
	} `yaml:"info"`
	Schemes []string `yaml: ","`
	Paths map[string]map[string]methodData
}

type methodData struct {
	Summary string
	Description string
	Parameters interface{}
	Responses interface{}
}


func getStructTag(data reflect.StructField, t string) string {
	return data.Tag.Get(t)
}

func PostAddPaths(path string, data interface{}, responseData interface{}, notes *Notes) {
	_,haveData := jsonDecode(reflect.TypeOf(data))
	_,haveResponseData := jsonDecode(reflect.TypeOf(responseData))
	method := methodData{}
	method.Summary = notes.PostNote.Summary
	method.Description = notes.PostNote.Description
	method.Parameters = []interface{}{map[string]interface{}{"name": "person", "in": "body", "schema": haveData, "required": true}}
	method.Responses = map[string]interface{}{"200": map[string]interface{}{"schema": haveResponseData,"description": "persons succesfully created"}}
	if _,ok := Swagger.Paths[path]; ok {
		Swagger.Paths[path]["post"] = method
	} else {
		Swagger.Paths[path] = map[string]methodData{"post": method}
	}
	//return map[string]string{}{path: map[string]interface{}{"post": map[string]interface{}{"parameters": []interface{}{map[string]interface{}{"name": "person", "in": "body", "schema": haveData}}, "responses": map[string]interface{}{"200": map[string]string{"description": "persons succesfully created"}, "400": map[string]string{"description": "persions couldnt have been created"}}}}}
}
func PutAddPaths(path string, data interface{}, responseData interface{}, notes *Notes)  {
	path = path + "/{id}"
	_,haveData := jsonDecode(reflect.TypeOf(data))
	_,haveResponseData := jsonDecode(reflect.TypeOf(responseData))
	method := methodData{}
	method.Summary = notes.PutNote.Summary
	method.Description = notes.PutNote.Description
	method.Parameters = []interface{}{map[string]interface{}{"name": "id", "in": "path", "required": "true"},map[string]interface{}{"name": "body", "in": "body", "schema": haveData}}
	method.Responses = map[string]interface{}{"200": map[string]interface{}{"schema": haveResponseData,"description": "persons succesfully created"}}
	if _,ok := Swagger.Paths[path]; ok {
		Swagger.Paths[path]["put"] = method
	} else {
		Swagger.Paths[path] = map[string]methodData{"put": method}
	}
}

func GetAddPaths(path string, data interface{}, responseData interface{}, notes *Notes)  {
	_, havaData := formDecode(reflect.TypeOf(data))
	_,haveResponseData := jsonDecode(reflect.TypeOf(responseData))
	method := methodData{}
	method.Summary = notes.GetNote.Summary
	method.Description = notes.GetNote.Description
	method.Parameters = havaData
	method.Responses = map[string]interface{}{"200": map[string]interface{}{"schema": haveResponseData,"description": "persons succesfully created"}}
	if _, ok := Swagger.Paths[path]; ok {
		Swagger.Paths[path]["get"] = method
	} else {
		Swagger.Paths[path] = map[string]methodData{"get": method}
	}
}

func GetIdPaths(path string, data interface{}, responseData interface{},notes *Notes)  {
	path = path + "/{id}"
	_,haveResponseData := jsonDecode(reflect.TypeOf(responseData))
	method := methodData{}
	method.Summary = notes.GetIdNote.Summary
	method.Description = notes.GetIdNote.Description
	method.Parameters = []interface{}{map[string]string{"name": "id","in": "path", "required": "true"}}
	method.Responses = map[string]interface{}{"200": map[string]interface{}{"schema": haveResponseData,"description": "persons succesfully created"}}
	if _, ok := Swagger.Paths[path]; ok {
		Swagger.Paths[path]["get"] = method
	} else {
		Swagger.Paths[path] = map[string]methodData{"get": method}
	}

}

func DeleteIdPaths(path string, data interface{}, responseData interface{}, notes *Notes)  {
	path = path + "/{id}"
	_,haveResponseData := jsonDecode(reflect.TypeOf(responseData))
	method := methodData{}
	method.Summary = notes.DeleteNote.Summary
	method.Description = notes.DeleteNote.Description
	method.Parameters = []interface{}{map[string]string{"name": "id","in": "path", "required": "true" }}
	method.Responses = map[string]interface{}{"200": map[string]interface{}{"schema": haveResponseData,"description": "persons succesfully created"}}
	if _, ok := Swagger.Paths[path]; ok {
		Swagger.Paths[path]["delete"] = method
	} else {
		Swagger.Paths[path] = map[string]methodData{"delete": method}
	}

}

func formDecode(data reflect.Type) (string ,interface{})  {
	switch data.Kind() {
	case reflect.Struct:
		swaggerNode := make([]map[string]interface{},0,0)
		for i := 0; i < data.NumField(); i++ {
			ignoring := getStructTag(data.Field(i), "ignoring")
			if ignoring == "true" {
				continue
			}
			name := getStructTag(data.Field(i), "form")
			if name == "" {
				continue
			}
			dataType := map[string]interface{}{}

			swichType,nameType := formDecode(data.Field(i).Type)
			switch swichType {
			case "struct", "Array", "Ptr":
				//swaggerNode[name] = nameType
				continue
			case "String", "Bool", "Int", "Uint", "Float32", "Float64":
				required := getStructTag(data.Field(i), "binding")
				if required == "" {
					dataType = map[string]interface{}{"name":name,"type": nameType, "in": "query"}
				} else {
					dataType = map[string]interface{}{"name":name,"type": nameType, "in": "query", "required": true}
				}
				swaggerNode = append(swaggerNode, dataType)
			default:
				panic("未知类型" + swichType)
			}
		}
		return "struct" , swaggerNode
	case reflect.Slice, reflect.Array:
		swichType, haveData := formDecode(data.Elem())
		if swichType == "struct" {
			return "Array",map[string]interface{}{"type": "array", "items": haveData}
		} else {
			return "Array", map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": haveData}}
		}
	case reflect.Ptr:
		swichType, haveData := formDecode(data.Elem())
		return swichType,haveData
	case reflect.String:
		return "String", "string"
	case reflect.Bool:
		return "Bool", "boolean"
	case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
		return  "Int", "integer"
	case reflect.Uint,reflect.Uint8,reflect.Uint16,reflect.Uint32,reflect.Uint64:
		return "Uint", "number"
	case reflect.Float32:
		return "Float32", "number"
	case reflect.Float64:
		return "Float64", "number"
	default:
		return "default", "default"
	}
}
func jsonDecode(data reflect.Type) (string ,interface{}) {
	switch data.Kind() {
	case reflect.Struct:
		swaggerNode := make(map[string]interface{})
		for i := 0; i < data.NumField(); i++ {
			ignoring := getStructTag(data.Field(i), "ignoring")
			if ignoring == "true" {
				continue
			}
			if unicode.IsUpper([]rune(data.Field(i).Name)[0]) != true {
				continue
			}
			name := getStructTag(data.Field(i), "json")
			if name == "" {
				name = data.Field(i).Name
			}
			dataType := map[string]interface{}{}
			if data.Field(i).Type.String() == "time.Time" || data.Field(i).Type.String() == "*time.Time" {
				dataType = map[string]interface{}{"name": name, "type": "string", "in": "query"}
				swaggerNode[name] = dataType
				continue
			}
			swichType, nameType := jsonDecode(data.Field(i).Type)
			if data.Field(i).Anonymous == true && swichType == "struct" {
				//fmt.Println("anonymous is true")
				swichType = "Anonymous"
			}
			switch swichType {
			case "Anonymous":
				for k,v := range nameType.(map[string]interface{})["properties"].(map[string]interface{}) {
					swaggerNode[k] = v
				}
			case "struct", "Array", "Ptr":
				swaggerNode[name] = nameType
			case "String", "Bool", "Int", "Uint", "Float32", "Float64":
				dataType = map[string]interface{}{"name": name, "type": nameType, "in": "query"}
				//dataType = map[string]interface{}{"type": nameType}
				swaggerNode[name] = dataType
			default:
				panic("未知类型" + swichType)
			}
		}
		return "struct", map[string]interface{}{"type": "object", "properties": swaggerNode}
	case reflect.Slice, reflect.Array:
		swichType, haveData := jsonDecode(data.Elem())
		if swichType == "struct" {
			return "Array", map[string]interface{}{"type": "array", "items": haveData}
		} else {
			return "Array", map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": haveData}}
		}
	case reflect.Ptr:
		swichType, haveData := jsonDecode(data.Elem())
		return swichType, haveData
	case reflect.String:
		return "String", "string"
	case reflect.Bool:
		return "Bool", "boolean"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "Int", "integer"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "Uint", "number"
	case reflect.Float32:
		return "Float32", "number"
	case reflect.Float64:
		return "Float64", "number"
	default:
		return "default", "default"
	}
}

func oldjsonDecode(data reflect.Type) (string ,interface{})  {
	switch data.Kind() {
	case reflect.Struct:
		swaggerNode := make(map[string]interface{})
		for i := 0; i < data.NumField(); i++ {
			ignoring := getStructTag(data.Field(i), "ignoring")
			if ignoring == "true" {
				continue
			}
			name := getStructTag(data.Field(i), "json")
			if name == "" {
				//continue
				name = data.Field(i).Name
			}
			dataType := map[string]interface{}{}
			swichType,nameType := oldjsonDecode(data.Field(i).Type)
			switch swichType {
			case "struct", "Array", "Ptr":
				swaggerNode[name] = nameType
			case "String", "Bool", "Int", "Uint", "Float32", "Float64":
				dataType = map[string]interface{}{"name":name,"type": nameType, "in": "query"}
				//dataType = map[string]interface{}{"type": nameType}
				swaggerNode[name] = dataType
			default:
				panic("未知类型" + swichType)
			}
		}
		return "struct" , map[string]interface{}{"type": "object", "properties": swaggerNode}
	case reflect.Slice, reflect.Array:
		swichType, haveData := oldjsonDecode(data.Elem())
		if swichType == "struct" {
			return "Array",map[string]interface{}{"type": "array", "items": haveData}
		} else {
			return "Array", map[string]interface{}{"type": "array", "items": map[string]interface{}{"type": haveData}}
		}
	case reflect.Ptr:
		swichType, haveData := oldjsonDecode(data.Elem())
		return swichType,haveData
	case reflect.String:
		return "String", "string"
	case reflect.Bool:
		return "Bool", "boolean"
	case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
		return  "Int", "integer"
	case reflect.Uint,reflect.Uint8,reflect.Uint16,reflect.Uint32,reflect.Uint64:
		return "Uint", "number"
	case reflect.Float32:
		return "Float32", "number"
	case reflect.Float64:
		return "Float64", "number"
	default:
		return "default", "default"
	}
}


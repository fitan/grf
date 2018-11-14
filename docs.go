// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-11-09 11:36:47.4165845 +0800 DST m=+0.032466001

package grf

import (
	//"bytes"

	//"github.com/alecthomas/template"
	"github.com/swaggo/swag"
	"gopkg.in/yaml.v2"
)

//var doc = `{
//    "swagger": "2.0",
//    "info": {
//        "description": "This is a sample server Petstore server.",
//        "title": "Swagger Example API",
//        "termsOfService": "http://swagger.io/terms/",
//        "contact": {
//            "name": "API Support",
//            "url": "http://www.swagger.io/support",
//            "email": "support@swagger.io"
//        },
//        "license": {
//            "name": "Apache 2.0",
//            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
//        },
//        "version": "1.0"
//    },
//    "host": "petstore.swagger.io",
//    "basePath": "/v2",
//    "paths": {
//        "/api/v1/tags": {
//            "post": {
//                "summary": "新增文章标签",
//                "responses": {
//                    "200": {
//                        "description": "OK",
//                        "schema": {
//                            "type": "object",
//                            "$ref": "#/definitions/model.InterCmdb"
//                        }
//                    }
//                }
//            }
//        }
//    },
//    "definitions": {
//        "model.InterCmdb": {
//            "type": "object",
//            "required": [
//                "cpu",
//                "ip"
//            ],
//            "properties": {
//                "cpu": {
//                    "type": "string"
//                },
//                "dist": {
//                    "type": "string"
//                },
//                "ip": {
//                    "type": "string"
//                },
//                "mem": {
//                    "type": "string"
//                }
//            }
//        }
//    }
//}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	//t, err := template.New("swagger_info").Parse(doc)
	//if err != nil {
	//	return doc
	//}
	//
	//var tpl bytes.Buffer
	//if err := t.Execute(&tpl, SwaggerInfo); err != nil {
	//	return doc
	//}
	//
	//return tpl.String()
	doc, err := yaml.Marshal(Swagger)

	if err != nil {
		panic(err)
	}
	return string(doc)
}

func init() {
	swag.Register(swag.Name, &s{})
}

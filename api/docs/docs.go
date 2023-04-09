// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/clubs": {
            "get": {
                "summary": "クラブ一覧を取得",
                "responses": {}
            }
        },
        "/v1/clubs/{id}": {
            "get": {
                "summary": "クラブを取得",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Club ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/v1/clubs/{id}/activities": {
            "get": {
                "summary": "クラブの活動を取得",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Club ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/v1/me": {
            "get": {
                "summary": "自分のユーザー情報を取得する",
                "responses": {}
            }
        },
        "/v1/users": {
            "post": {
                "summary": "ユーザーを作成する",
                "parameters": [
                    {
                        "description": "ユーザー情報",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.requestUser"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "handler.requestUser": {
            "type": "object",
            "properties": {
                "class": {
                    "type": "string",
                    "example": "IE3A"
                },
                "icon": {
                    "type": "string",
                    "example": ""
                },
                "name": {
                    "type": "string",
                    "example": "ぴよ太郎"
                },
                "readme": {
                    "type": "string",
                    "example": ""
                },
                "uid": {
                    "type": "integer",
                    "example": 220000
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "soberkoder@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/groups": {
            "get": {
                "description": "Get all groups",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Get all groups",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/groups.Dto"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "create a new groups",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Create a new groups",
                "parameters": [
                    {
                        "description": "Create groups",
                        "name": "groupPayLoad",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/groups.PayLoad"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/groups.Dto"
                        }
                    }
                }
            }
        },
        "/api/v1/groups/{idGroup}": {
            "put": {
                "description": "Update the group corresponding to the input groupId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Update group identified by the given orderId",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the group to be updated",
                        "name": "idGroup",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "update groups",
                        "name": "groupPayLoad",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/groups.PayLoad"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/groups.Dto"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete the group corresponding to the input idGroup",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Delete group identified by the given idGroup",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the group to be updated",
                        "name": "idGroup",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        }
    },
    "definitions": {
        "groups.Dto": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nameGroup": {
                    "type": "string"
                },
                "owner": {
                    "type": "string"
                },
                "private": {
                    "type": "boolean"
                },
                "thumbnail": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "groups.PayLoad": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "nameGroup": {
                    "type": "string"
                },
                "private": {
                    "type": "boolean"
                },
                "type": {
                    "type": "string"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:5000",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "vdatchat API",
	Description: "This is a sample serice for managing orders",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}

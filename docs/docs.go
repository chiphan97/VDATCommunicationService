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
        "contact": {},
        "license": {},
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
                "summary": "Update group by groupId",
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
        },
        "/api/v1/groups/{idGroup}/members": {
            "get": {
                "description": "Get all member groups",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groupUser"
                ],
                "summary": "Get all member groups",
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
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/userdetail.Dto"
                                }
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "delete user to group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groupUser"
                ],
                "summary": "delete user to group",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the group to be add user",
                        "name": "idGroup",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            },
            "patch": {
                "description": "add user to group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groupUser"
                ],
                "summary": "add user to group",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the group to be updated",
                        "name": "idGroup",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "add user to group",
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
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/groups.Dto"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/groups/{idGroup}/members/{userId}": {
            "delete": {
                "description": "delete user to group by admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groupUser"
                ],
                "summary": "delete user to group by admin",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID group",
                        "name": "idGroup",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "ID user want delete",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/api/v1/user": {
            "get": {
                "description": "find user by keyword",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "find user by keyword",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name search by keyword",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "pageSize",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/userdetail.Dto"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/user/info": {
            "get": {
                "description": "check user api",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "check user api",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userdetail.Dto"
                        }
                    }
                }
            }
        },
        "/api/v1/user/online": {
            "delete": {
                "description": "user logout api",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "user logout",
                "parameters": [
                    {
                        "type": "string",
                        "description": "hostName",
                        "name": "hostName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "socketId",
                        "name": "socketId",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
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
        },
        "userdetail.Dto": {
            "type": "object",
            "properties": {
                "first": {
                    "type": "string"
                },
                "fullName": {
                    "type": "string"
                },
                "hostName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "socketId": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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

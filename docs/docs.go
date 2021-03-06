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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/book": {
            "post": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "Create new book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create new book",
                "parameters": [
                    {
                        "description": "book data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateBookInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/book/:id": {
            "get": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "Find a book",
                "produces": [
                    "application/json"
                ],
                "summary": "Find a book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "book identifier",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "Delete a book",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "book identifier",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "Update a book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "book identifier",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "book data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateBookInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/books": {
            "get": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "Get all books",
                "produces": [
                    "application/json"
                ],
                "summary": "Get all books",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CreateBookInput": {
            "type": "object",
            "required": [
                "author",
                "author_id",
                "title"
            ],
            "properties": {
                "author": {
                    "type": "string"
                },
                "author_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Authorization": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	Version:     "beta",
	Host:        "localhost:7010",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "rest-api DOC",
	Description: "This is the rest-api documentation",
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

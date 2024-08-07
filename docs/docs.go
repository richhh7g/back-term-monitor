// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Richhh7g",
            "url": "https://github.com/richhh7g",
            "email": "richhh7g@protonmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/richhh7g/back-term-monitor/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/check-term": {
            "post": {
                "description": "Faça a checagem de termos de marca nos resultados de pesquisa do Google.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1",
                    "Termos"
                ],
                "summary": "Checar termos de marca",
                "parameters": [
                    {
                        "description": "Corpo da requisição",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ChecarTermoBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "ChecarTermoBody": {
            "type": "object",
            "required": [
                "email",
                "termos"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "email@example.com"
                },
                "termos": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Term One",
                        "Term Two"
                    ]
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Term Monitor",
	Description:      "API do [Term Monitor](https://github.com/richhh7g/back-term-monitor) para monitorar concorrentes que usam termos de marca em resultados patrocinados do Google.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

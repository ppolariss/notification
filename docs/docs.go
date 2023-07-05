// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Maintainer OpenTreeHole",
            "email": "dev@fduhole.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "https://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.MessageModel"
                        }
                    }
                }
            }
        },
        "/callback/mipush": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Callback"
                ],
                "summary": "Mipush Callback",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/messages": {
            "post": {
                "description": "Send to multiple recipients and save to db, admin only.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Message"
                ],
                "summary": "Send a message",
                "parameters": [
                    {
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    }
                }
            }
        },
        "/users/push-tokens": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "List Tokens of a User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.PushToken"
                            }
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Add Token of a User",
                "parameters": [
                    {
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PushToken"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.PushToken"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Delete the token of a user's certain device",
                "parameters": [
                    {
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/token.DeleteModel"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/users/push-tokens/_all": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Delete all tokens of a user",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Map": {
            "type": "object",
            "additionalProperties": {}
        },
        "models.Message": {
            "type": "object",
            "required": [
                "code",
                "data",
                "description",
                "message",
                "recipients",
                "url"
            ],
            "properties": {
                "code": {
                    "type": "string"
                },
                "data": {
                    "$ref": "#/definitions/models.Map"
                },
                "description": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "recipients": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "models.MessageModel": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.PushService": {
            "type": "integer",
            "enum": [
                0,
                1,
                2
            ],
            "x-enum-varnames": [
                "ServiceAPNS",
                "ServiceFCM",
                "ServiceMipush"
            ]
        },
        "models.PushToken": {
            "type": "object",
            "required": [
                "device_id",
                "token"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "device_id": {
                    "type": "string",
                    "maxLength": 64
                },
                "service": {
                    "$ref": "#/definitions/models.PushService"
                },
                "token": {
                    "type": "string",
                    "maxLength": 64
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "description": "not required",
                    "type": "integer"
                }
            }
        },
        "token.DeleteModel": {
            "type": "object",
            "properties": {
                "device_id": {
                    "type": "string",
                    "maxLength": 64
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Notification Center",
	Description:      "This is a notification microservice.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

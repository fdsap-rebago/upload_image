// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "FDSAP Support"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/fetch": {
            "post": {
                "description": "FETCH IMAGE DATA",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "KPLUS"
                ],
                "summary": "FETCH IMAGE DATA",
                "parameters": [
                    {
                        "description": "Image ID",
                        "name": "img_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.ImageID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Uploaded_Images"
                        }
                    }
                }
            }
        },
        "/fetch_image_data/{img_id}": {
            "get": {
                "description": "FETCH IMAGE DATA",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "KPLUS"
                ],
                "summary": "FETCH IMAGE DATA",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Image ID",
                        "name": "img_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.ImageSrc"
                        }
                    }
                }
            }
        },
        "/get_um/{userInput}": {
            "get": {
                "description": "Fetch User Data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "KPLUS"
                ],
                "summary": "Fetch User Data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Input",
                        "name": "userInput",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.UserManagement"
                        }
                    }
                }
            }
        },
        "/upload": {
            "post": {
                "description": "UPLOAD IMAGE",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "KPLUS"
                ],
                "summary": "UPLOAD IMAGE",
                "parameters": [
                    {
                        "description": "Complete image path",
                        "name": "img_src",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.ImageSrc"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Uploaded_Images"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.ImageID": {
            "type": "object",
            "properties": {
                "img_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "main.ImageSrc": {
            "type": "object",
            "properties": {
                "img_src": {
                    "type": "string"
                }
            }
        },
        "main.Uploaded_Images": {
            "type": "object",
            "properties": {
                "img_data": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "img_id": {
                    "type": "integer"
                },
                "img_type": {
                    "type": "string"
                }
            }
        },
        "main.UserManagement": {
            "type": "object",
            "properties": {
                "branch_names": {
                    "type": "string"
                },
                "check_status": {
                    "type": "string"
                },
                "given_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "middle_name": {
                    "type": "string"
                },
                "roles": {
                    "type": "string"
                },
                "user_name": {
                    "description": "RoleID   int    ` + "`" + `json:\"role_id\"` + "`" + `\nRoleName string ` + "`" + `json:\"role_name\"` + "`" + `\nRoleDesc string ` + "`" + `json:\"role_desc\"` + "`" + `",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "192.168.0.138:4040",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Image Upload",
	Description:      "Image Upload Swagger",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

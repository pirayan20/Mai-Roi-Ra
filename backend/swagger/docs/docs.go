// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/events": {
            "post": {
                "description": "Create a new event with the provided details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Create new event",
                "parameters": [
                    {
                        "description": "Create Event Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structure.CreateEventRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structure.CreateEventResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/test": {
            "get": {
                "description": "Get a test message",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Test"
                ],
                "summary": "GetTest",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structure.TestResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "structure.CreateEventRequest": {
            "type": "object",
            "required": [
                "activities",
                "admin_id",
                "deadline",
                "description",
                "end_date",
                "event_name",
                "location_id",
                "organizer_id",
                "participant_fee",
                "start_date",
                "status"
            ],
            "properties": {
                "activities": {
                    "type": "string"
                },
                "admin_id": {
                    "type": "string"
                },
                "deadline": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "event_image": {
                    "type": "string"
                },
                "event_name": {
                    "type": "string"
                },
                "location_id": {
                    "type": "string"
                },
                "organizer_id": {
                    "type": "string"
                },
                "participant_fee": {
                    "type": "number"
                },
                "start_date": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "structure.CreateEventResponse": {
            "type": "object",
            "properties": {
                "event_id": {
                    "type": "string"
                }
            }
        },
        "structure.TestResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Mai-Roi-Ra Swagger API",
	Description:      "This is a sample server Mai-Roi-Ra api gateway.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

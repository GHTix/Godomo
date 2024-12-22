// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/v1/battery": {
            "get": {
                "description": "get the last battery mesure",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sensor"
                ],
                "summary": "get the last battery mesure",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Sensor"
                        }
                    }
                }
            }
        },
        "/api/v1/humidity": {
            "get": {
                "description": "get the last humidity mesure",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sensor"
                ],
                "summary": "get the last humidity mesure",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Sensor"
                        }
                    }
                }
            }
        },
        "/api/v1/pressure": {
            "get": {
                "description": "get the last pressure mesure",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sensor"
                ],
                "summary": "get the last pressure mesure",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Sensor"
                        }
                    }
                }
            }
        },
        "/api/v1/sensor": {
            "get": {
                "description": "get the last data of the sensor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sensor"
                ],
                "summary": "get the last data of the sensor",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/aqara.Aqara"
                        }
                    }
                }
            }
        },
        "/api/v1/temperature": {
            "get": {
                "description": "get the last temperature mesure",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sensor"
                ],
                "summary": "get the last temperature mesure",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Sensor"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "aqara.Aqara": {
            "type": "object",
            "properties": {
                "battery": {
                    "type": "integer"
                },
                "humidity": {
                    "type": "number"
                },
                "linkquality": {
                    "type": "integer"
                },
                "power_outage_count": {
                    "type": "integer"
                },
                "pressure": {
                    "type": "number"
                },
                "temperature": {
                    "type": "number"
                },
                "voltage": {
                    "type": "integer"
                }
            }
        },
        "common.Sensor": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "unit": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
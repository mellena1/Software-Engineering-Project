// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-02-26 18:47:48.842125 -0500 EST m=+0.041514844

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "The API for the code camp counting program.",
        "title": "Code Camp Counter API",
        "contact": {},
        "license": {},
        "version": "1.0"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/rooms": {
            "get": {
                "description": "Return a row of the room",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a room",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Room"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            }
        },
        "/api/v1/session": {
            "get": {
                "description": "Return a list of all sessions",
                "produces": [
                    "application/json"
                ],
                "summary": "Get all sessions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Session"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            }
        },
        "/api/v1/speaker": {
            "get": {
                "description": "Return a list of all speakers",
                "produces": [
                    "application/json"
                ],
                "summary": "Get all speakers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Speaker"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "db.Room": {
            "type": "object",
            "properties": {
                "capacity": {
                    "type": "integer",
                    "example": 50
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "My Room Name"
                }
            }
        },
        "db.Session": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "Session Name"
                },
                "room": {
                    "type": "object",
                    "$ref": "#/definitions/db.Room"
                },
                "speaker": {
                    "type": "object",
                    "$ref": "#/definitions/db.Speaker"
                },
                "timeslot": {
                    "type": "object",
                    "$ref": "#/definitions/db.Timeslot"
                }
            }
        },
        "db.Speaker": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "firstname.lastname@gmail.com"
                },
                "firstName": {
                    "type": "string",
                    "example": "Bob"
                },
                "lastName": {
                    "type": "string",
                    "example": "Smith"
                }
            }
        },
        "db.Timeslot": {
            "type": "object",
            "properties": {
                "endTime": {
                    "type": "string",
                    "example": "2019-10-01 23:00:00"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "startTime": {
                    "type": "string",
                    "example": "2019-02-18 21:00:00"
                }
            }
        }
    }
}`

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
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}

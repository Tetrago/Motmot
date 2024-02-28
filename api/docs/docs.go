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
        "/auth/login": {
            "post": {
                "description": "Log in to user and authenticate with the backend",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "User login information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/course/department/{dep}": {
            "get": {
                "description": "Queries for all UF courses in a three-letter department prefix",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "course"
                ],
                "summary": "Get courses in department",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Three-letter department prefix",
                        "name": "dep",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request"
                    },
                    "503": {
                        "description": "Service Unavailable"
                    }
                }
            }
        },
        "/course/group/{dep}/{code}": {
            "get": {
                "description": "Gets (or creates) group of specified course",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "course"
                ],
                "summary": "Get group of specified course",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Three-letter department prefix",
                        "name": "dep",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Four-digit (with potential postfix) course code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    },
                    "503": {
                        "description": "Service Unavailable"
                    }
                }
            }
        },
        "/group/all": {
            "get": {
                "description": "Gets all public groups",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "group"
                ],
                "summary": "Gets groups",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/group.AllResponseItem"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/group/get/{id}": {
            "get": {
                "description": "Gets group information",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "group"
                ],
                "summary": "Get group",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Group ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/group.GetResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/group/history/{id}": {
            "get": {
                "description": "Gets message history from a group in descending order",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "group"
                ],
                "summary": "Gets group messages",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Group ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Max number of messages to retreive (\u003c= 20)",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "UTC time cutoff; searches in reverse from this point (inclusive)",
                        "name": "before",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/group.HistoryResponseItem"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/user/bio": {
            "post": {
                "description": "Updates a user's bio",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Updates bio",
                "parameters": [
                    {
                        "description": "User token and new bio",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.BioRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/user/get/{ident}": {
            "get": {
                "description": "Fetches publically available user information and groups.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Fetch user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User identifier",
                        "name": "ident",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.GetResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/user/join": {
            "post": {
                "description": "Adds a user to a group",
                "tags": [
                    "user"
                ],
                "summary": "Join group",
                "parameters": [
                    {
                        "description": "User token and group to join",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.JoinRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/user/profile_picture": {
            "post": {
                "description": "Uploads a new profile picture, replacing the old one",
                "tags": [
                    "user"
                ],
                "summary": "Upload profile picture",
                "parameters": [
                    {
                        "description": "User token and profile picture",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.PostProfilePictureRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/user/profile_picture/{ident}": {
            "get": {
                "description": "Gets a user's profile picture from their identifier",
                "produces": [
                    "image/jpeg"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Retrieves profile picture",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User identifier",
                        "name": "ident",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "Registers a new user given the provided arguments",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User registration information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.RegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/ws/{group}": {
            "get": {
                "description": "Opens a WebSocket for a user on a group",
                "tags": [
                    "ws"
                ],
                "summary": "Opens a WebSocket",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Group ID",
                        "name": "group",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "401": {
                        "description": "Unauthorized"
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "group.AllResponseItem": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "group.GetResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "group.HistoryResponseItem": {
            "type": "object",
            "properties": {
                "contents": {
                    "type": "string"
                },
                "iat": {
                    "type": "integer"
                },
                "message_id": {
                    "type": "integer"
                },
                "user_ident": {
                    "type": "string"
                }
            }
        },
        "user.BioRequest": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string"
                }
            }
        },
        "user.GetResponse": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string"
                },
                "display_name": {
                    "type": "string"
                },
                "groups": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/user.GetResponseGroup"
                    }
                },
                "ident": {
                    "type": "string"
                }
            }
        },
        "user.GetResponseGroup": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "user.JoinRequest": {
            "type": "object",
            "properties": {
                "group_id": {
                    "type": "integer"
                }
            }
        },
        "user.PostProfilePictureRequest": {
            "type": "object",
            "properties": {
                "jpeg": {
                    "type": "string"
                }
            }
        },
        "user.RegisterRequest": {
            "type": "object",
            "properties": {
                "display_name": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "user.RegisterResponse": {
            "type": "object",
            "properties": {
                "ident": {
                    "type": "string"
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

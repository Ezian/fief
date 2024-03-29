// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API of the server for Fief Diplomatie",
    "title": "Fief Diplomatie API",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/auth/signin": {
      "post": {
        "description": "Signin with login/password and retrieve JWT token for further requests",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "operationId": "Login",
        "parameters": [
          {
            "name": "Login",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/LoginInfo"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful login",
            "schema": {
              "$ref": "#/definitions/LoginSuccess"
            }
          },
          "401": {
            "description": "Wrong Login/Password",
            "schema": {
              "description": "Error message",
              "type": "string"
            }
          },
          "500": {
            "description": "Server error",
            "schema": {
              "description": "Error message",
              "type": "string"
            }
          }
        }
      }
    },
    "/auth/signup": {
      "post": {
        "description": "Register a new user",
        "tags": [
          "user"
        ],
        "operationId": "Register",
        "parameters": [
          {
            "description": "Registration payload",
            "name": "signup",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/RegisterUser"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful registration"
          },
          "400": {
            "description": "Registration failure",
            "schema": {
              "description": "Error message",
              "type": "string"
            }
          },
          "500": {
            "description": "Server error",
            "schema": {
              "description": "Error message",
              "type": "string"
            }
          }
        }
      }
    },
    "/games": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "List existing games",
        "tags": [
          "game"
        ],
        "responses": {
          "200": {
            "description": "TODO",
            "schema": {
              "$ref": "#/definitions/GamesSuccess"
            }
          },
          "401": {
            "description": "User not authorized to view games",
            "schema": {
              "description": "Error Message",
              "type": "string"
            }
          }
        }
      }
    },
    "/games/new": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Create a new game",
        "tags": [
          "manage"
        ],
        "responses": {
          "200": {
            "description": "TODO"
          }
        }
      }
    },
    "/games/{id}": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Get current state of an existing game",
        "tags": [
          "game"
        ],
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "TODO"
          }
        }
      }
    },
    "/games/{id}/instructions": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Retrieve the current instructions for the authenticated user",
        "tags": [
          "game"
        ],
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "TODO"
          }
        }
      },
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Post an instruction for the authenticated user on the session, if no instruction defined",
        "tags": [
          "game"
        ],
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "instruction",
            "in": "body",
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "TODO"
          }
        }
      }
    },
    "/games/{id}/join": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Join an existing game",
        "tags": [
          "manage"
        ],
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "TODO"
          }
        }
      }
    }
  },
  "definitions": {
    "GameInfo": {
      "description": "General information about existing games",
      "type": "object",
      "required": [
        "id",
        "name",
        "players",
        "status"
      ],
      "properties": {
        "id": {
          "description": "UID of the game",
          "type": "string"
        },
        "joined": {
          "description": "If the current user as joined this game",
          "type": "boolean"
        },
        "name": {
          "description": "Human readable name of the game",
          "type": "string"
        },
        "players": {
          "type": "object",
          "required": [
            "joined",
            "required"
          ],
          "properties": {
            "joined": {
              "description": "Current count of players who have joined the game",
              "type": "number"
            },
            "required": {
              "description": "Total players required to launch the game",
              "type": "number"
            }
          }
        },
        "status": {
          "description": "A simple string describing the status (defined on server side)",
          "type": "string"
        }
      }
    },
    "GamesSuccess": {
      "type": "object",
      "properties": {
        "available": {
          "description": "Game created but waiting for enough players",
          "type": "array",
          "items": {
            "$ref": "#/definitions/GameInfo"
          }
        },
        "joined": {
          "description": "Games joined by the current player",
          "type": "array",
          "items": {
            "$ref": "#/definitions/GameInfo"
          }
        }
      }
    },
    "LoginInfo": {
      "type": "object",
      "required": [
        "login",
        "password"
      ],
      "properties": {
        "login": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "LoginSuccess": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "RegisterUser": {
      "type": "object",
      "required": [
        "email",
        "login",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string"
        },
        "login": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API of the server for Fief Diplomatie",
    "title": "Fief Diplomatie API",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/auth/signin": {
      "post": {
        "description": "Signin with login/password and retrieve JWT token for further requests",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "operationId": "Login",
        "parameters": [
          {
            "name": "Login",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/LoginInfo"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful login",
            "schema": {
              "$ref": "#/definitions/LoginSuccess"
            }
          },
          "401": {
            "description": "Wrong Login/Password",
            "schema": {
              "description": "Error message",
              "type": "string"
            }
          },
          "500": {
            "description": "Server error",
            "schema": {
              "description": "Error message",
              "type": "string"
            }
          }
        }
      }
    },
    "/auth/signup": {
      "post": {
        "description": "Register a new user",
        "tags": [
          "user"
        ],
        "operationId": "Register",
        "parameters": [
          {
            "description": "Registration payload",
            "name": "signup",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/RegisterUser"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful registration"
          },
          "400": {
            "description": "Registration failure",
            "schema": {
              "description": "Error message",
              "type": "string"
            }
          },
          "500": {
            "description": "Server error",
            "schema": {
              "description": "Error message",
              "type": "string"
            }
          }
        }
      }
    },
    "/games": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "List existing games",
        "tags": [
          "game"
        ],
        "responses": {
          "200": {
            "description": "TODO",
            "schema": {
              "$ref": "#/definitions/GamesSuccess"
            }
          },
          "401": {
            "description": "User not authorized to view games",
            "schema": {
              "description": "Error Message",
              "type": "string"
            }
          }
        }
      }
    },
    "/games/new": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Create a new game",
        "tags": [
          "manage"
        ],
        "responses": {
          "200": {
            "description": "TODO"
          }
        }
      }
    },
    "/games/{id}": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Get current state of an existing game",
        "tags": [
          "game"
        ],
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "TODO"
          }
        }
      }
    },
    "/games/{id}/instructions": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Retrieve the current instructions for the authenticated user",
        "tags": [
          "game"
        ],
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "TODO"
          }
        }
      },
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Post an instruction for the authenticated user on the session, if no instruction defined",
        "tags": [
          "game"
        ],
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "instruction",
            "in": "body",
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "TODO"
          }
        }
      }
    },
    "/games/{id}/join": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Join an existing game",
        "tags": [
          "manage"
        ],
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "TODO"
          }
        }
      }
    }
  },
  "definitions": {
    "GameInfo": {
      "description": "General information about existing games",
      "type": "object",
      "required": [
        "id",
        "name",
        "players",
        "status"
      ],
      "properties": {
        "id": {
          "description": "UID of the game",
          "type": "string"
        },
        "joined": {
          "description": "If the current user as joined this game",
          "type": "boolean"
        },
        "name": {
          "description": "Human readable name of the game",
          "type": "string"
        },
        "players": {
          "type": "object",
          "required": [
            "joined",
            "required"
          ],
          "properties": {
            "joined": {
              "description": "Current count of players who have joined the game",
              "type": "number"
            },
            "required": {
              "description": "Total players required to launch the game",
              "type": "number"
            }
          }
        },
        "status": {
          "description": "A simple string describing the status (defined on server side)",
          "type": "string"
        }
      }
    },
    "GameInfoPlayers": {
      "type": "object",
      "required": [
        "joined",
        "required"
      ],
      "properties": {
        "joined": {
          "description": "Current count of players who have joined the game",
          "type": "number"
        },
        "required": {
          "description": "Total players required to launch the game",
          "type": "number"
        }
      }
    },
    "GamesSuccess": {
      "type": "object",
      "properties": {
        "available": {
          "description": "Game created but waiting for enough players",
          "type": "array",
          "items": {
            "$ref": "#/definitions/GameInfo"
          }
        },
        "joined": {
          "description": "Games joined by the current player",
          "type": "array",
          "items": {
            "$ref": "#/definitions/GameInfo"
          }
        }
      }
    },
    "LoginInfo": {
      "type": "object",
      "required": [
        "login",
        "password"
      ],
      "properties": {
        "login": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "LoginSuccess": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "RegisterUser": {
      "type": "object",
      "required": [
        "email",
        "login",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string"
        },
        "login": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
}

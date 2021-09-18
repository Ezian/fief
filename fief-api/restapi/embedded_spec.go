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
  "basePath": "/v1",
  "paths": {
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
            "description": "TODO"
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
    },
    "/games/{id}/status": {
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
    "/signin": {
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
            "name": "credentials",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/credentials"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Todo"
          }
        }
      }
    },
    "/signup/{code}": {
      "post": {
        "description": "Create user with login/password",
        "parameters": [
          {
            "name": "credentials",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/credentials"
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
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "User not found",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "Authorization code that allow player to signup in server",
          "name": "code",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "LoginSuccess": {
      "type": "object",
      "properties": {
        "success": {
          "type": "boolean"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "credentials": {
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
    }
  },
  "securityDefinitions": {
    "isAuthenticated": {
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
  "basePath": "/v1",
  "paths": {
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
            "description": "TODO"
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
    },
    "/games/{id}/status": {
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
    "/signin": {
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
            "name": "credentials",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/credentials"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Todo"
          }
        }
      }
    },
    "/signup/{code}": {
      "post": {
        "description": "Create user with login/password",
        "parameters": [
          {
            "name": "credentials",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/credentials"
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
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "User not found",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "Authorization code that allow player to signup in server",
          "name": "code",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "LoginSuccess": {
      "type": "object",
      "properties": {
        "success": {
          "type": "boolean"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "credentials": {
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
    }
  },
  "securityDefinitions": {
    "isAuthenticated": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
}

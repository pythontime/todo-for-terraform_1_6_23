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
    "application/spkane.todo-list.v1+json"
  ],
  "produces": [
    "application/spkane.todo-list.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "The product of a tutorial on goswagger.io",
    "title": "A To Do list application",
    "version": "1.0.0"
  },
  "paths": {
    "/": {
      "get": {
        "tags": [
          "todos"
        ],
        "operationId": "findTodos",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list the todo operations",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/item"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "todos"
        ],
        "operationId": "addOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/item"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/{id}": {
      "put": {
        "tags": [
          "todos"
        ],
        "operationId": "updateOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/item"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "todos"
        ],
        "operationId": "destroyOne",
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "item": {
      "type": "object",
      "required": [
        "description"
      ],
      "properties": {
        "completed": {
          "type": "boolean"
        },
        "description": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/spkane.todo-list.v1+json"
  ],
  "produces": [
    "application/spkane.todo-list.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "The product of a tutorial on goswagger.io",
    "title": "A To Do list application",
    "version": "1.0.0"
  },
  "paths": {
    "/": {
      "get": {
        "tags": [
          "todos"
        ],
        "operationId": "findTodos",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list the todo operations",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/item"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "todos"
        ],
        "operationId": "addOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/item"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/{id}": {
      "put": {
        "tags": [
          "todos"
        ],
        "operationId": "updateOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/item"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "todos"
        ],
        "operationId": "destroyOne",
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "item": {
      "type": "object",
      "required": [
        "description"
      ],
      "properties": {
        "completed": {
          "type": "boolean"
        },
        "description": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        }
      }
    }
  }
}`))
}

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
        "/admin/add-dish": {
            "get": {
                "description": "\"添加一道菜\"",
                "tags": [
                    "public"
                ],
                "summary": "AddDish",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "price",
                        "name": "price",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "discount",
                        "name": "discount",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "detail",
                        "name": "detail",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/admin/add-table": {
            "get": {
                "description": "\"加台\"",
                "tags": [
                    "public"
                ],
                "summary": "AddTable",
                "parameters": [
                    {
                        "type": "string",
                        "description": "table_id",
                        "name": "table_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/admin/complete-table": {
            "get": {
                "description": "\"结台\"",
                "tags": [
                    "public"
                ],
                "summary": "CompleteTable",
                "parameters": [
                    {
                        "type": "string",
                        "description": "table_id",
                        "name": "table_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/admin/del-table": {
            "get": {
                "description": "\"减台\"",
                "tags": [
                    "public"
                ],
                "summary": "DelTable",
                "parameters": [
                    {
                        "type": "string",
                        "description": "table_id",
                        "name": "table_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/admin/delete-dish": {
            "get": {
                "description": "\"删除一道菜，如果dish_id不存在，不会返回失败\"",
                "tags": [
                    "public"
                ],
                "summary": "DeleteDish",
                "parameters": [
                    {
                        "type": "string",
                        "description": "dish_id",
                        "name": "dish_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/admin/get-consume-record": {
            "get": {
                "description": "\"获取营业额\"",
                "tags": [
                    "public"
                ],
                "summary": "GetTurnover",
                "parameters": [
                    {
                        "type": "string",
                        "description": "start",
                        "name": "start",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "end",
                        "name": "end",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/check-user-password": {
            "get": {
                "description": "\"检查md5加密用户密码是否正确\"",
                "tags": [
                    "public"
                ],
                "summary": "CallCheckUserPassword",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uin",
                        "name": "uin",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "phone_number",
                        "name": "phone_number",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/get-dish-info": {
            "get": {
                "description": "\"获得菜品详情\"",
                "tags": [
                    "public"
                ],
                "summary": "GetDishInfo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "dish_id",
                        "name": "dish_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/get-dish-list": {
            "get": {
                "description": "\"获取菜品列表\"",
                "tags": [
                    "public"
                ],
                "summary": "GetDishList",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/get-table-info": {
            "get": {
                "description": "\"获取桌台详情，包括桌台状态、点菜等\"",
                "tags": [
                    "public"
                ],
                "summary": "GetTableInfo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "桌台id",
                        "name": "table_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/get-tables-status": {
            "get": {
                "description": "\"获得全部桌台状态\"",
                "tags": [
                    "public"
                ],
                "summary": "GetTablesStatus",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "\"登录\"",
                "tags": [
                    "public"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "手机号",
                        "name": "phone_number",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "MD5加密密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/login-by-token": {
            "post": {
                "description": "\"登录\"",
                "tags": [
                    "public"
                ],
                "summary": "LoginByJWT",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/modify-password": {
            "post": {
                "description": "\"修改密码\"",
                "tags": [
                    "public"
                ],
                "summary": "ModifyPassword",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uin",
                        "name": "uin",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "phone_numer",
                        "name": "phone_number",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "old_password",
                        "name": "old_password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "new_password",
                        "name": "new_password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/open-table": {
            "get": {
                "description": "\"开台\"",
                "tags": [
                    "public"
                ],
                "summary": "OpenTable",
                "parameters": [
                    {
                        "type": "string",
                        "description": "table_id",
                        "name": "table_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/order-dish": {
            "post": {
                "description": "\"点菜\"",
                "tags": [
                    "public"
                ],
                "summary": "OrderDish",
                "parameters": [
                    {
                        "type": "string",
                        "description": "table_id",
                        "name": "table_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "dish_id",
                        "name": "dish_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "\"ping一下服务器\"",
                "tags": [
                    "example"
                ],
                "summary": "ping",
                "parameters": [
                    {
                        "type": "string",
                        "description": "回复",
                        "name": "message",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "\"注册新用户\"",
                "tags": [
                    "public"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "MD5加密密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "手机",
                        "name": "phone_number",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    }
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

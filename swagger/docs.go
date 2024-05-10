// Code generated by swaggo/swag. DO NOT EDIT.

package swagger

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
        "/v1/common/generateId": {
            "get": {
                "description": "生成id-描述",
                "tags": [
                    "公共接口"
                ],
                "summary": "雪花ID生成",
                "parameters": [
                    {
                        "type": "string",
                        "description": "认证信息 eg:xxxx-xxxx-xxxx-xxx",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "生成id数量 1-1000",
                        "name": "num",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功",
                        "schema": {
                            "$ref": "#/definitions/internal_api_http_servicev1.HttpGenerateIDResponse"
                        }
                    }
                }
            }
        },
        "/v1/exampleGet": {
            "get": {
                "description": "get接口示例",
                "tags": [
                    "Example"
                ],
                "summary": "get接口示例",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "create_time",
                        "name": "create_time",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "query_id",
                        "name": "query_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user_id 用户id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功",
                        "schema": {
                            "$ref": "#/definitions/_internal_api_http_dto.ExampleGetResponse"
                        }
                    }
                }
            }
        },
        "/v1/exampleGetOne": {
            "get": {
                "description": "getOne接口示例",
                "tags": [
                    "Example"
                ],
                "summary": "getOne接口示例",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "create_time",
                        "name": "create_time",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "query_id",
                        "name": "query_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user_id 用户id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功",
                        "schema": {
                            "$ref": "#/definitions/_internal_api_http_dto.ExampleGetOneResponse"
                        }
                    }
                }
            }
        },
        "/v1/examplePost": {
            "post": {
                "description": "\nios购买项类型 \u003ca href=\"https://developer.apple.com/documentation/appstoreconnectapi/list_all_in-app_purchases_for_an_app\"\u003e 详情\u003c/a\u003e \u003c/br\u003e \n\n\nandroid订阅 \u003ca href=\"https://developers.google.com/android-publisher/api-ref/rest/v3/monetization.subscriptions/list?hl=zh-cn\"\u003e 详情\u003c/a\u003e \u003c/br\u003e \n\nandroid非订阅\u003ca href=\"https://developers.google.com/android-publisher/api-ref/rest/v3/inappproducts/list?hl=zh-cn\"\u003e 详情\u003c/a\u003e \u003c/br\u003e \n\nandroid订阅产品的类型\u003ca href=\"https://developers.google.com/android-publisher/api-ref/rest/v3/inappproducts?hl=zh-cn#PurchaseType\"\u003e 详情\u003c/a\u003e \u003c/br\u003e",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Example"
                ],
                "summary": "post 接口 示例",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/_internal_api_http_dto.ExamplePostRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "请求成功",
                        "schema": {
                            "$ref": "#/definitions/_internal_api_http_dto.ExamplePostResponse"
                        },
                        "headers": {
                            "Location": {
                                "type": "string",
                                "description": "/entity/1"
                            },
                            "Token": {
                                "type": "string",
                                "description": "token"
                            },
                            "Token2": {
                                "type": "string",
                                "description": "token2"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "_internal_api_http_dto.ExampleGetOneResponse": {
            "type": "object",
            "required": [
                "code",
                "data",
                "msg",
                "traceId"
            ],
            "properties": {
                "code": {
                    "description": "code:  0 成功; 非0失败;",
                    "type": "integer",
                    "example": 1
                },
                "data": {
                    "$ref": "#/definitions/_internal_api_http_dto.UserPortraitData"
                },
                "msg": {
                    "description": "错误消息",
                    "type": "string",
                    "example": "success"
                },
                "traceId": {
                    "description": "traceId",
                    "type": "string"
                }
            }
        },
        "_internal_api_http_dto.ExampleGetResponse": {
            "type": "object",
            "required": [
                "code",
                "data",
                "msg",
                "traceId"
            ],
            "properties": {
                "code": {
                    "description": "code:  0 成功; 非0失败;",
                    "type": "integer",
                    "example": 1
                },
                "data": {
                    "$ref": "#/definitions/_internal_api_http_dto.UserPortraitData"
                },
                "msg": {
                    "description": "错误消息",
                    "type": "string",
                    "example": "success"
                },
                "traceId": {
                    "description": "traceId",
                    "type": "string"
                }
            }
        },
        "_internal_api_http_dto.ExamplePostRequest": {
            "type": "object",
            "required": [
                "app_id",
                "c_ver",
                "cmd_id",
                "fid",
                "lang",
                "question",
                "token"
            ],
            "properties": {
                "app_id": {
                    "description": "AppID 应用id",
                    "type": "string"
                },
                "c_ver": {
                    "description": "CVer 客户端协议版本号",
                    "type": "string"
                },
                "cmd_id": {
                    "description": "CmdId 功能类型枚举:\n1 AUTH鉴权;\n2 QA 问答；\n4 总结SUMMARY;\n5 KEYWORDS 关键词；\n6 CLASSIFICATION 文档分类；\n7 RENAME  重命名；\n8 mindmap 思维导图；",
                    "type": "integer",
                    "example": 1
                },
                "fid": {
                    "description": "Fid： 文档id",
                    "type": "string"
                },
                "ignore_feature_prompt": {
                    "description": "IgnoreFeaturePrompt 是否忽略功能检查  1 是； 0 否；",
                    "type": "integer"
                },
                "lang": {
                    "description": "Lang 语言",
                    "type": "string"
                },
                "question": {
                    "description": "Question 问题",
                    "type": "string"
                },
                "show_gpt_2": {
                    "description": "show_gpt_2",
                    "type": "string"
                },
                "token": {
                    "description": "Token 鉴权token",
                    "type": "string",
                    "example": "sdfsdsdfsd"
                }
            }
        },
        "_internal_api_http_dto.ExamplePostResponse": {
            "type": "object",
            "required": [
                "code",
                "data",
                "msg",
                "traceId"
            ],
            "properties": {
                "code": {
                    "description": "code:  0 成功; 非0失败;",
                    "type": "integer",
                    "example": 1
                },
                "data": {
                    "$ref": "#/definitions/_internal_api_http_dto.UserPortraitData"
                },
                "msg": {
                    "description": "错误消息",
                    "type": "string",
                    "example": "success"
                },
                "traceId": {
                    "description": "traceId",
                    "type": "string"
                }
            }
        },
        "_internal_api_http_dto.UserPortraitData": {
            "type": "object",
            "properties": {
                "country": {
                    "description": "国家",
                    "type": "string"
                },
                "last_login": {
                    "description": "上次登陆时间",
                    "type": "string"
                },
                "user_id": {
                    "description": "用户id",
                    "type": "string"
                },
                "vip_info": {
                    "description": "是否为VIP，0/1",
                    "type": "integer"
                }
            }
        },
        "internal_api_http_servicev1.HttpGenerateIDResponse": {
            "type": "object",
            "required": [
                "code",
                "data",
                "msg",
                "traceId"
            ],
            "properties": {
                "code": {
                    "description": "code:  0 成功; 非0失败;",
                    "type": "integer",
                    "example": 1
                },
                "data": {
                    "description": "生成id数组",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "msg": {
                    "description": "错误消息",
                    "type": "string",
                    "example": "success"
                },
                "traceId": {
                    "description": "traceId",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:3013",
	BasePath:         "/goodsCenterLogic",
	Schemes:          []string{"http"},
	Title:            "gin-http API",
	Description:      "gin-http服务文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

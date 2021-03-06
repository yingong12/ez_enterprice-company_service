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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/audit": {
            "get": {
                "description": "审核搜索",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "审核"
                ],
                "summary": "审核搜索",
                "parameters": [
                    {
                        "description": "字段注解",
                        "name": "xxx",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/audit.Search"
                        }
                    }
                ],
                "responses": {}
            },
            "post": {
                "description": "审核提交",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "审核"
                ],
                "summary": "审核提交",
                "parameters": [
                    {
                        "description": "字段注解",
                        "name": "xxx",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/audit.Create"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/enterprise": {
            "post": {
                "description": "企业新建",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "企业"
                ],
                "summary": "企业新建",
                "parameters": [
                    {
                        "description": "字段注解",
                        "name": "xxx",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/request.Create"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/valuate/search": {
            "get": {
                "description": "估值搜索",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "估值"
                ],
                "summary": "估值搜索",
                "parameters": [
                    {
                        "type": "string",
                        "description": "企业ID",
                        "name": "appID",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页大小",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Valuate"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "audit.Create": {
            "type": "object",
            "properties": {
                "app_id": {
                    "type": "string"
                },
                "app_type": {
                    "description": "0-企业  1-机构",
                    "type": "integer"
                },
                "form_data": {
                    "description": "审核表单信息 json",
                    "type": "string"
                }
            }
        },
        "audit.Search": {
            "type": "object",
            "properties": {
                "app_id": {
                    "description": "企业id精确查询",
                    "type": "string"
                },
                "app_name": {
                    "description": "企业名字模糊查询",
                    "type": "string"
                },
                "audit_ids": {
                    "description": "审核id",
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "registration_number": {
                    "description": "注册号",
                    "type": "string"
                },
                "states": {
                    "description": "状态",
                    "type": "string"
                }
            }
        },
        "model.EnterpriseMuttable": {
            "type": "object",
            "properties": {
                "business_scope": {
                    "type": "string"
                },
                "company_type": {
                    "type": "integer"
                },
                "district": {
                    "type": "string"
                },
                "estimate_value": {
                    "type": "integer"
                },
                "industry": {
                    "type": "string"
                },
                "introduction": {
                    "type": "string"
                },
                "legal_representative": {
                    "type": "string"
                },
                "legal_representative_id_img": {
                    "type": "string"
                },
                "license_img": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "register_capital": {
                    "type": "integer"
                },
                "registration_address": {
                    "type": "string"
                },
                "registration_number": {
                    "type": "string"
                },
                "shar_holders_json": {
                    "type": "string"
                },
                "share_holder_info": {
                    "type": "string"
                },
                "stage": {
                    "type": "integer"
                },
                "state": {
                    "description": "审核状态",
                    "type": "integer"
                }
            }
        },
        "model.Valuate": {
            "type": "object",
            "properties": {
                "app_id": {
                    "type": "string"
                },
                "created_at": {
                    "description": "返回给业务侧",
                    "type": "string"
                },
                "form_data": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "requested_at": {
                    "type": "string"
                },
                "result": {
                    "description": "估值结果",
                    "type": "string"
                },
                "state": {
                    "type": "integer"
                },
                "udated_at": {
                    "type": "string"
                },
                "valuate_id": {
                    "type": "string"
                }
            }
        },
        "request.Create": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "字段",
                    "$ref": "#/definitions/model.EnterpriseMuttable"
                },
                "parent_id": {
                    "description": "机构id 非必填",
                    "type": "string"
                },
                "uid": {
                    "description": "用户id",
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

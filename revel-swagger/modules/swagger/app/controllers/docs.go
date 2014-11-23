
package controllers
//This file is generated automatically. Do not try to edit it manually.

var resourceListingJson = `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "apis": [
        {
            "path": "/helloworld",
            "description": "helloworld"
        }
    ],
    "info": {
        "contact": "api@contact.me",
        "termsOfServiceUrl": "http://google.com/",
        "license": "BSD",
        "licenseUrl": "http://opensource.org/licenses/BSD-2-Clause"
    }
}`
var apiDescriptionsJson = map[string]string{"helloworld":`{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "127.0.0.1:9000",
    "resourcePath": "/helloworld",
    "apis": [
        {
            "path": "/j/helloworld/{key}",
            "description": "获取指定 key 数据",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "AjaxGetHelloWorld",
                    "type": ".github.com.busyStone.go-labs.revel-swagger.app.controllers.AjaxResult",
                    "items": {},
                    "summary": "获取指定 key 数据",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "key",
                            "description": "id 必须 \u003e 0",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "",
                            "responseModel": ".github.com.busyStone.go-labs.revel-swagger.app.controllers.AjaxResult"
                        },
                        {
                            "code": 400,
                            "message": "",
                            "responseModel": ".github.com.busyStone.go-labs.revel-swagger.app.controllers.AjaxResult"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/j/helloworld",
            "description": "设置指定 key 数据",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "AjaxPostHelloWorld",
                    "type": ".github.com.busyStone.go-labs.revel-swagger.app.controllers.AjaxResult",
                    "items": {},
                    "summary": "设置指定 key 数据",
                    "parameters": [
                        {
                            "paramType": "form",
                            "name": "key",
                            "description": "id 必须 \u003e 0",
                            "dataType": "int",
                            "type": "int",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "form",
                            "name": "data",
                            "description": "数据",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "",
                            "responseModel": ".github.com.busyStone.go-labs.revel-swagger.app.controllers.AjaxResult"
                        },
                        {
                            "code": 400,
                            "message": "",
                            "responseModel": ".github.com.busyStone.go-labs.revel-swagger.app.controllers.AjaxResult"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        ".github.com.busyStone.go-labs.revel-swagger.app.controllers.AjaxResult": {
            "id": ".github.com.busyStone.go-labs.revel-swagger.app.controllers.AjaxResult",
            "properties": {
                "data": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "error": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "success": {
                    "type": "bool",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`,}

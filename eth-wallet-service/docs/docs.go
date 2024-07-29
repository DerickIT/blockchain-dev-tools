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
        "/eth/burn-nft": {
            "post": {
                "description": "Burn (destroy) a specified NFT",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ethereum"
                ],
                "summary": "Burn an NFT",
                "parameters": [
                    {
                        "description": "NFT burn details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.BurnNFTRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/eth/contract-balance": {
            "get": {
                "description": "Get the balance of a specified contract address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ethereum"
                ],
                "summary": "Get the balance of a contract",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Contract address",
                        "name": "address",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/eth/create-wallet": {
            "post": {
                "description": "Create a new Ethereum wallet and return the address and private key",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ethereum"
                ],
                "summary": "Create a new Ethereum wallet",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/eth/gas-balance": {
            "get": {
                "description": "Get the gas (ETH) balance of a specified address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ethereum"
                ],
                "summary": "Get the gas balance of an address",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ethereum address",
                        "name": "address",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/eth/multi-sig-transfer": {
            "post": {
                "description": "Perform a multi-signature transfer of Ethereum",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ethereum"
                ],
                "summary": "Perform a multi-signature transfer",
                "parameters": [
                    {
                        "description": "Multi-signature transfer details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.MultiSigTransferRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/eth/token-supply": {
            "get": {
                "description": "Get the total supply of a specified ERC20 token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ethereum"
                ],
                "summary": "Get the total supply of a token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token contract address",
                        "name": "address",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/eth/withdraw": {
            "post": {
                "description": "Withdraw tokens from the service to a specified address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ethereum"
                ],
                "summary": "Withdraw tokens",
                "parameters": [
                    {
                        "description": "Withdrawal details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.WithdrawRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "services.BurnNFTRequest": {
            "type": "object"
        },
        "services.MultiSigTransferRequest": {
            "type": "object"
        },
        "services.WithdrawRequest": {
            "type": "object"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Ethereum Wallet Service API",
	Description:      "This is a sample server for Ethereum wallet operations.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

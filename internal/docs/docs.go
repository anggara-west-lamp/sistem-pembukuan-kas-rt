package docs

import "github.com/swaggo/swag"

type s struct{}

func (s *s) ReadDoc() string { return doc }

func init() { swag.Register(swag.Name, &s{}) }

const doc = `{
  "swagger": "2.0",
  "info": {
    "title": "Sistem Pembukuan Kas RT API",
    "version": "1.0",
    "description": "API untuk autentikasi, pengguna/role, kas, transaksi, dan laporan."
  },
  "basePath": "/api/v1",
  "schemes": ["http"],
  "securityDefinitions": {
    "BearerAuth": {"type": "apiKey", "name": "Authorization", "in": "header"}
  },
  "paths": {
    "/healthz": {
      "get": {"summary": "Health check", "responses": {"200": {"description": "OK"}}}
    },
    "/auth/login": {
      "post": {
        "summary": "Login",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "parameters": [
          {"in": "body", "name": "body", "schema": {"type": "object", "required": ["email","password"], "properties": {"email": {"type": "string"}, "password": {"type": "string"}}}}
        ],
        "responses": {"200": {"description": "OK"}, "401": {"description": "Unauthorized"}}
      }
    },
    "/users": {
      "get": {"summary": "List users", "security": [{"BearerAuth": []}], "responses": {"200": {"description": "OK"}}},
      "post": {"summary": "Create user", "security": [{"BearerAuth": []}], "parameters": [{"in":"body","name":"body","schema": {"type":"object","required":["email","password"],"properties":{"name":{"type":"string"},"email":{"type":"string"},"password":{"type":"string"}}}}], "responses": {"201": {"description": "Created"}}}
    },
    "/kas": {
      "get": {"summary": "List kas", "security": [{"BearerAuth": []}], "responses": {"200": {"description": "OK"}}},
      "post": {"summary": "Create kas", "security": [{"BearerAuth": []}], "parameters": [{"in":"body","name":"body","schema": {"type":"object","required":["type","category"],"properties":{"type":{"type":"string"},"category":{"type":"string"},"amount":{"type":"number"},"description":{"type":"string"}}}}], "responses": {"201": {"description": "Created"}}}
    },
    "/transaksi": {
      "post": {"summary": "Create transaksi", "security": [{"BearerAuth": []}], "parameters": [{"in":"body","name":"body","schema": {"type":"object","required":["user_id","kas_id","amount"],"properties":{"user_id":{"type":"integer"},"kas_id":{"type":"integer"},"amount":{"type":"number"},"status":{"type":"string"},"proof_url":{"type":"string"}}}}], "responses": {"201": {"description": "Created"}}}
    },
    "/laporan": {
      "get": {"summary": "Laporan bulanan", "security": [{"BearerAuth": []}], "parameters": [{"in":"query","name":"month","type":"string","required":true}], "responses": {"200": {"description": "OK"}}}
    }
  }
}`


// Package classification CRUD
//
// Documentation of our awesome API.
//
//     Schemes: http
//     BasePath: /api/v1/
//     Version: 1.0.0
//     Host: localhost:10000
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
// swagger:meta
package docs

import (
	"REST_soft/StructUser"
)

// swagger:route GET /users/{id} users getUser
// returns user by specified id
// Produces:
//     - application/json
// responses:
//      200: userGetResponse
//      400: badRequest

// swagger:parameters getUser
type UserIdParam struct {
	// Specifies uuid for a user
	//
	// unique: true
	// in: path
	// example: 3ca4ce84-ed71-42aa-8d1a-c0e001d3b8b4
	Id string `json:"id"`
}

// Error struct with error explanation string
// swagger:response badRequest
type BadRequestResponseWrapper struct {
	// in:body
	Body string
}

// swagger:response userGetResponse
type userGetResponse struct {
	// Specifies uuid for a user
	//
	// in: body
	//
	Body StructUser.User
}

// swagger:route GET /users users getUsers
// returns users
// Produces:
//     - application/json
// responses:
//      200: UsersGetResponse
//      400: badRequest

// swagger:response UsersGetResponse
type UsersGetResponse struct {
	//in: body
	Body []StructUser.User
}

// swagger:route POST /users users createUser
// Produces:
//     - application/json
// responses:
//      400: badRequest

// swagger:parameters createUser
type UserPostParam struct {
	// in: body
	Body StructUser.User
}

// swagger:route PUT /users/{id} users updateUser
// Produces:
//     - application/json
// responses:
//      400: badRequest

// swagger:parameters updateUser
type UserPutParam struct {
	//in: body
	Body StructUser.User
}

// swagger:route DELETE /users/{id} users deleteUser
// Produces:
//     - application/json
// responses:
//      400: badRequest

// swagger:parameters deleteUser
type UserDelParam struct {
	// Specifies uuid for a user
	//
	// unique: true
	// in: path
	// example: 3ca4ce84-ed71-42aa-8d1a-c0e001d3b8b4
	Id string `json:"id"`
}


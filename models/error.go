package models

//Error is a type which has error code and message
type Error struct {
	Code    int    `json:"code" example:"33"`
	Message string `json:"message" example:"this is a error of type 33"`
}

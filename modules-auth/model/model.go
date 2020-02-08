package model

import "github.com/dgrijalva/jwt-go"

// User models a REST API user
type User struct {
	ID       int    `json:"id" validate:"numeric,gte=0"`
	Name     string `json:"name" validate:"required,min=5,max=30"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5,max=15"`
	Age      int    `json:"age" validate:"required,numeric,gte=0,lte=130"`
}

type UserToken struct {
	UserID string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.StandardClaims
}
type Token struct {
}

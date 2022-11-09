package domain

import "time"

type Code struct {
	ID           uint
	Code         string
	AccessToken  string
	TokenType    string
	RefreshToken string
	Expiry       time.Time
}

type PesanCore struct {
	ID     uint
	Email  string
	Status string
}
type Repository interface {
	InsertCode(code string) error
	UpdateCode(code Code) error
	GetCode() (Code, error)
	GetPesan() PesanCore
}

type Services interface {
	AddCode(Code string) error
	UpdateCode(code Code) error
	GetCode() (Code, error)
	GetPesan() PesanCore
}

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
	ID       uint
	IdRanger uint
	Email    string
	Status   string
}

type RangerCore struct {
	ID          uint
	UserID      uint
	Docs        string
	Price       uint
	Detail      string
	Status      string
	StatusApply string
}

type Repository interface {
	InsertCode(code string) error
	UpdateCode(code Code) error
	GetCode() (Code, error)
	GetPesan() (PesanCore, RangerCore)
}

type Services interface {
	AddCode(Code string) error
	UpdateCode(code Code) error
	GetCode() (Code, error)
	GetPesan() (PesanCore, RangerCore)
}

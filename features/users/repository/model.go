package repository

import (
	"GunTour/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName    string
	Email       string
	Password    string
	Role        string
	Address     string
	Dob         string
	Gender      string
	UserPicture string
}

func FromCore(uc users.Core) User {
	return User{
		Model:       gorm.Model{ID: uint(uc.ID), CreatedAt: uc.CreatedAt, UpdatedAt: uc.UpdatedAt},
		FullName:    uc.FullName,
		Email:       uc.Email,
		Password:    uc.Password,
		Role:        uc.Role,
		Address:     uc.Address,
		Dob:         uc.Dob,
		Gender:      uc.Gender,
		UserPicture: uc.UserPicture,
	}
}

func ToCore(u User) users.Core {
	return users.Core{
		ID:          int(u.ID),
		FullName:    u.FullName,
		Email:       u.Email,
		Password:    u.Password,
		Role:        u.Role,
		Address:     u.Address,
		Dob:         u.Dob,
		Gender:      u.Gender,
		UserPicture: u.UserPicture,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

func ToCoreArray(ua []User) []users.Core {
	var arr []users.Core
	for _, val := range ua {
		arr = append(arr, users.Core{
			ID:          int(val.ID),
			FullName:    val.FullName,
			Email:       val.Email,
			Password:    val.Password,
			Role:        val.Role,
			Address:     val.Address,
			Dob:         val.Dob,
			Gender:      val.Gender,
			UserPicture: val.UserPicture,
			CreatedAt:   val.CreatedAt,
			UpdatedAt:   val.UpdatedAt,
		})
	}
	return arr
}

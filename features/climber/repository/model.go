package repository

import (
	"GunTour/features/climber/domain"

	"gorm.io/gorm"
)

type Climber struct {
	gorm.Model
	IsClimber     int
	MaleClimber   int
	FemaleClimber int
}

func ToDomain(c Climber) domain.Core {
	return domain.Core{
		ID:            c.ID,
		IsClimber:     c.IsClimber,
		MaleClimber:   c.MaleClimber,
		FemaleClimber: c.FemaleClimber,
	}
}

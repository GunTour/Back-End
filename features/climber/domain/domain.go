package domain

type Core struct {
	ID            uint
	IsClimber     int
	MaleClimber   int
	FemaleClimber int
}

type Repository interface {
	GetClimber() (Core, error)
}

type Services interface {
	ShowClimber() (Core, error)
}

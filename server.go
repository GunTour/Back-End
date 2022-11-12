package main

import (
	"GunTour/config"

	ad "GunTour/features/admin/delivery"
	ar "GunTour/features/admin/repository"
	as "GunTour/features/admin/services"
	bd "GunTour/features/booking/delivery"
	br "GunTour/features/booking/repository"
	bs "GunTour/features/booking/services"
	cd "GunTour/features/climber/delivery"
	cr "GunTour/features/climber/repository"
	cs "GunTour/features/climber/services"
	gd "GunTour/features/google/delivery"
	gr "GunTour/features/google/repository"
	gs "GunTour/features/google/services"
	pd "GunTour/features/product/delivery"
	pr "GunTour/features/product/repository"
	ps "GunTour/features/product/services"
	rd "GunTour/features/ranger/delivery"
	rr "GunTour/features/ranger/repository"
	rs "GunTour/features/ranger/services"
	ud "GunTour/features/users/delivery"
	ur "GunTour/features/users/repository"
	us "GunTour/features/users/services"
	"GunTour/utils/database"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	db := database.InitDB(cfg)
	database.MigrateDB(db)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	uRepo := ur.New(db)
	bRepo := br.New(db)
	aRepo := ar.New(db)
	gRepo := gr.New(db)
	pRepo := pr.New(db)
	cRepo := cr.New(db)
	uService := us.New(uRepo)
	bService := bs.New(bRepo)
	aService := as.New(aRepo)
	pService := ps.New(pRepo)
	gService := gs.New(gRepo)
	cService := cs.New(cRepo)
	ud.New(e, uService)
	bd.New(e, bService)
	ad.New(e, aService)
	pd.New(e, pService)
	gd.New(e, gService)
	cd.New(e, cService)

	rRepo := rr.New(db)
	rService := rs.New(rRepo)
	rd.New(e, rService)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.ServerPort)))
}

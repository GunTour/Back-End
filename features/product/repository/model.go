package repository

import (
	"GunTour/features/product/domain"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName    string
	RentPrice      int
	Detail         string
	Note           string
	ProductPicture string
}

func FromCore(pc domain.Core) Product {
	return Product{
		Model:          gorm.Model{ID: pc.ID, CreatedAt: pc.CreatedAt, UpdatedAt: pc.UpdatedAt},
		ProductName:    pc.ProductName,
		RentPrice:      pc.RentPrice,
		Detail:         pc.Detail,
		Note:           pc.Note,
		ProductPicture: pc.ProductPicture,
	}
}

func ToCore(p Product) domain.Core {
	return domain.Core{
		ID:             p.ID,
		ProductName:    p.ProductName,
		RentPrice:      p.RentPrice,
		Detail:         p.Detail,
		Note:           p.Note,
		ProductPicture: p.ProductPicture,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
	}
}

func ToCoreArray(ap []Product) []domain.Core {
	var arr []domain.Core
	for _, val := range ap {
		arr = append(arr, domain.Core{
			ID:             val.ID,
			ProductName:    val.ProductName,
			RentPrice:      val.RentPrice,
			Detail:         val.Detail,
			Note:           val.Note,
			ProductPicture: val.ProductPicture,
			CreatedAt:      val.CreatedAt,
			UpdatedAt:      val.UpdatedAt,
		})
	}
	return arr
}

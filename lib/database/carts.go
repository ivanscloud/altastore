package database

import (
	"altastore/config"
	"altastore/models"
)

func InsertCarts(cart models.Carts) (interface{}, error) {

	if err := config.DB.Save(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

package model

import (
	// "github.com/jinzhu/gorm"
	"gorm.io/gorm"
    "gorm.io/driver/mysql"
)

func getDB() (*gorm.DB, error) {
	dsn := "yutanj:spacebros224@tcp(localhost:3306)/allergen-menu?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func (u *MenuAllergen) Read(tableName string) (MenuAllergen, error) {
	menuAllergen := MenuAllergen{}
	db, err := getDB()
	if err != nil {
		return menuAllergen, err
	}
	sqlDb, err := db.DB()
	if err != nil {
		return menuAllergen, err
	}
	defer sqlDb.Close()
 
	result := []*MenuAllergen{}
	//err = db.Table(tableName).Select("user_id,user_name,password").Where("user_id = ?", u.UserId).Scan(&result).Error
	err = db.Table(tableName).Select("menu_id,menu_name,restaurant_id").Scan(&result).Error
	if err != nil {
		return menuAllergen, err
	}
	for _, r := range result {
		menuAllergen.MenuName = r.MenuName
		menuAllergen.MenuId = r.MenuId
		menuAllergen.RestaurantId = r.RestaurantId
	}
	return menuAllergen, nil
}

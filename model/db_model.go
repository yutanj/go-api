package model

import (
	// "github.com/jinzhu/gorm"
	// "gorm.io/gorm"
    // "gorm.io/driver/mysql"	
)

// func getDB() (*gorm.DB, error) {
// 	dsn := "yutanj:spacebros224@tcp(localhost:3306)/allergen-menu?charset=utf8mb4&parseTime=True&loc=Local"
// 	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
// }

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
	err = db.Table(tableName).Select("menu_id,menu_name,restaurant_id").Find(&result).Error
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

//指定したアレルゲンを含まないメニューを返す
func (u *MenuAllergen) ReadSelect() (MenuAllergen, error) {
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
	err = db.Table("am_menu_allergen").Where("crab <= ? AND shrimp <= ? AND egg <= ? AND dairy <= ? AND peanut <= ? AND wheat <= ? AND buckwheat <= ? AND abalone <= ? AND squid <= ? AND salmonroe <= ? AND orange <= ? AND cashewnut <= ? AND kiwi <= ? AND beef <= ? AND walnut <= ? AND sesame <= ? AND salmon <= ? AND mackerel <= ? AND soybean <= ? AND chicken <= ? AND banana <= ? AND pork <= ? AND matsutake <= ? AND peach <= ? AND yam <= ? AND apple <= ? AND gelatin <= ? AND almond <= ?", u.Crab, u.Shrimp, u.Egg, u.Dairy, u.Peanut, u.Wheat, u.BuckWheat, u.Abalone, u.Squid, u.SalmonRoe, u.Orange, u.CashewNut, u.Kiwi, u.Beef, u.Walnut, u.Sesame, u.Salmon, u.Mackerel, u.SoyBean, u.Chicken, u.Banana, u.Pork, u.Matsutake, u.Peach, u.Yam, u.Apple, u.Gelatin, u.Almond).Find(&result).Error
	//err = db.Table("am_menu_allergen").Where("crab <= ? AND shrimp <= ?", u.Crab, u.Shrimp).Find(&result).Error

	if err != nil {
		return menuAllergen, err
	}
	for _, r := range result {
		menuAllergen.MenuName = r.MenuName
		menuAllergen.MenuId = r.MenuId
		menuAllergen.RestaurantId = r.RestaurantId
		menuAllergen.Crab = r.Crab
		menuAllergen.Shrimp = r.Shrimp
		menuAllergen.Egg = r.Egg
		menuAllergen.Dairy = r.Dairy
		menuAllergen.Peanut = r.Peanut
		menuAllergen.Wheat = r.Wheat
		menuAllergen.BuckWheat = r.BuckWheat
		menuAllergen.Abalone = r.Abalone
		menuAllergen.Squid = r.Squid
		menuAllergen.SalmonRoe = r.SalmonRoe
		menuAllergen.Orange = r.Orange
		menuAllergen.CashewNut = r.CashewNut
		menuAllergen.Kiwi = r.Kiwi
		menuAllergen.Beef = r.Beef
		menuAllergen.Walnut = r.Walnut
		menuAllergen.Sesame = r.Sesame
		menuAllergen.Salmon = r.Salmon
		menuAllergen.Mackerel = r.Mackerel
		menuAllergen.SoyBean = r.SoyBean
		menuAllergen.Chicken = r.Chicken
		menuAllergen.Banana = r.Banana
		menuAllergen.Pork = r.Pork
		menuAllergen.Matsutake = r.Matsutake
		menuAllergen.Peach = r.Peach
		menuAllergen.Yam = r.Yam
		menuAllergen.Apple = r.Apple
		menuAllergen.Gelatin = r.Gelatin
		menuAllergen.Almond = r.Almond
	}
	//println(&menuAllergen)
	return menuAllergen, nil
}



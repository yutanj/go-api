package main

import (
	//"github.com/gin-gonic/gin"
	"github.com/yutanj/go-api/model"
	"fmt"
)

func main(){
	// r := gin.Default()
	// r.GET("/home", func(c *gin.Context){
        
	// })
    // r.Run()

	//tableName := "am_menu_allergen"
	u := &model.MenuAllergen{}

	// ユーザーのアレルゲンを定義
	// 0: アレルゲン
	// 1: アレルゲンだが、ライン混入はok
	// 2: アレルゲンではない
	u.Crab = 1
    u.Shrimp = 1 
	u.Egg = 1
	u.Dairy = 2
	u.Peanut = 2
	u.Wheat = 2
	u.BuckWheat = 2
	u.Abalone = 2
	u.Squid = 2
	u.SalmonRoe = 2
	u.Orange = 2
	u.CashewNut = 2
	u.Kiwi = 2
	u.Beef = 2
	u.Walnut = 2
	u.Sesame = 2
	u.Salmon = 2
	u.Mackerel = 2
	u.SoyBean = 2
	u.Chicken = 2
	u.Banana = 2
	u.Pork = 2
	u.Matsutake = 2
	u.Peach = 2
	u.Yam = 2
	u.Apple = 2
	u.Gelatin = 2
	u.Almond = 2


	//menuAllergen, err := u.Read(tableName)
	menuAllergen, err := u.ReadSelect()
	if err != nil {
		fmt.Println(err)
		return
	}
    fmt.Println(menuAllergen)
	
	
}
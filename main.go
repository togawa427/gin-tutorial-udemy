package main

import (
	"gin-tutorial-udemy/controllers"
	"gin-tutorial-udemy/models"
	"gin-tutorial-udemy/repositories"
	"gin-tutorial-udemy/services"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
		{ID: 2, Name: "商品2", Price: 2000, Description: "説明2", SoldOut: true},
		{ID: 3, Name: "商品3", Price: 3000, Description: "説明3", SoldOut: false},
	}

	itemRepository := repositories.NewItemMemoryRepository(items)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pog",
	// 	})
	// })
	r.GET("/items", itemController.FindAll)
	r.Run("0.0.0.0:8090") // 0.0.0.0:8080 でサーバーを立てます。
}
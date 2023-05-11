package main

import (
	"food-delivery/component/appctx"
	"food-delivery/middleware"
	restauranttransport "food-delivery/module/restaurant/transport/restaurant"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Restaurant struct {
	Id   int    `json:"id" grom:"column:id;"`
	Name string `json:"name" grom:"column:name;"`
	Addr string `json:"addr" grom:"column:addr;"`
}

func (Restaurant) tableName() string { return "restaurants" }

func main() {
	dsn := os.Getenv("MYSQL_CONV_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(db)
	r := gin.Default()

	appCtx := appctx.NewAppContext(db)

	//r.Use(middleware.Recover(appCtx))

	//v1 := r.Group("/v1", middleware.Recover(appCtx))
	v1 := r.Group("/v1", middleware.Recover(appCtx))

	restaurants := v1.Group("restaurants")
	restaurants.POST("/", restauranttransport.CreateRestaurant(appCtx))
	restaurants.DELETE("/:id", restauranttransport.DeleteRestaurant(appCtx))
	restaurants.GET("/", restauranttransport.ListRestaurant(appCtx))

	r.Run()

	//newRestaurant := Restaurant{Name: "A", Addr: "Washington"}
	//
	//if err := db.Create(&newRestaurant).Error; err != nil {
	//	log.Println(err)
	//}
	//
	//var myRestaurant Restaurant

}

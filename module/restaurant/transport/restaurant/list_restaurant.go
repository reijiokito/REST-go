package restauranttransport

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	restaurantbiz "food-delivery/module/restaurant/biz"
	restaurantmodel "food-delivery/module/restaurant/model"
	restaurantstorage "food-delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSONP(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		pagingData.Fullfill()

		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSONP(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var result []restaurantmodel.Restaurant

		//Initiate abstract
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewListRestaurantBiz(store)

		result, err := biz.ListDataWithCondition(c, &filter, &pagingData)
		if err != nil {
			c.JSONP(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSONP(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}

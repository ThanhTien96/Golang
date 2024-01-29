package ginItem

import (
	"net/http"
	"strconv"
	"todo-list-social/common"
	"todo-list-social/modules/item/biz"
	"todo-list-social/modules/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// get detail task by id
func GetItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	
		store := storage.NewSQLStore(db)
		business := biz.NewGetItemBiz(store)

		data, err := business.GetItemById(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}
}

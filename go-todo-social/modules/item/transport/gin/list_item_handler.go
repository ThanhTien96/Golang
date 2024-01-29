package ginItem

import (
	"net/http"
	"todo-list-social/common"
	"todo-list-social/modules/item/biz"
	"todo-list-social/modules/item/model"
	"todo-list-social/modules/item/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// get all items
func ListItems(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.Process()


		var filter model.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}


		store := storage.NewSQLStore(db)
		business := biz.ListItemStorage(store)

		result, err := business.ListItem(c.Request.Context(), &filter, &paging);
		if  err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}





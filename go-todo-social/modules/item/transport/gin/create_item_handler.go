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

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		var data model.TodoItemCreation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)

		business := biz.NewCreateItemBiz(store)

		if err := business.CreateNewItem(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.Id))
	}
}

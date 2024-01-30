package ginItem

import (
	"net/http"
	"strconv"
	"todo-list-social/common"
	"todo-list-social/modules/item/biz"
	"todo-list-social/modules/item/model"
	"todo-list-social/modules/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// update to do list
func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemUpdate

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewUpdateItemBiz(store)

		if err := business.UpdateItem(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}


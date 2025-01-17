package ginitem

import (
	"github.com/gin-gonic/gin"
	"golandtitorial/common"
	"golandtitorial/modules/item/biz"
	"golandtitorial/modules/item/model"
	"golandtitorial/modules/item/storage"
	"gorm.io/gorm"
	"net/http"
)

func UpdateItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemUpdate
		id := c.Param("id")
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := storage.NewSqlStore(db)
		business := biz.NewUpdateItemBiz(store)
		if err := business.UpdateItemById(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponce(true))
	}
}

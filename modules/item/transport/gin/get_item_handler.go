package ginitem

import (
	"golandtitorial/common"
	"golandtitorial/modules/item/biz"
	"golandtitorial/modules/item/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		store := storage.NewSqlStore(db)
		business := biz.NewGetItemBiz(store)
		data, err := business.GetItembyId(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponce(data))
	}
}

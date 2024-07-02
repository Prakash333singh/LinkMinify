package routes

import (
	"net/http"
	"time"

	"github.com/Prakash333singh/url_shotner/api/database"
	"github.com/Prakash333singh/url_shotner/api/models"
	"github.com/gin-gonic/gin"
)

func EditURl(c *gin.Context) {
	shortID := c.Param("shortid")

	var body models.Request

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot Parse Json",
		})
	}
	r := database.CreateClient(0)
	defer r.Close()

	//check if the shortId exits int the db or not
	val, err := r.Get(database.Ctx, shortID).Result()

	if err != nil || val == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "ShortId doesn't exists",
		})

	}
	//update the content of the url,expiry time with the shortId
	err = r.Set(database.Ctx, shortID, body.URL, body.Expiry*3600*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to update the shortend content",
		})

		c.JSON(http.StatusOK, gin.H{
			"message": "The Content Has been Updated!!!",
		})
	}
}

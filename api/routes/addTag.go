package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Prakash333singh/url_shotner/api/database"
	"github.com/gin-gonic/gin"
)

type TagRequest struct {
	ShortID string `json:"shortID"`
	Tag     string `json:"tag"`
}

func AddTag(c *gin.Context) {
	var tagRequest TagRequest
	if err := c.ShouldBindJSON(&tagRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request Body",
		})
		return
	}

	shortId := tagRequest.ShortID
	tag := tagRequest.Tag

	r := database.CreateClient(0)
	defer r.Close()

	val, err := r.Get(database.Ctx, shortId).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data not found for the given ShortID",
		})
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		//if the daya is not a json object
		data = make(map[string]interface{})
		data["data"] = val
	}
	//check if tags filed alredy exists and its a slice of strings
	var tags []string
	if existingTags, ok := data["tags"].([]interface{}); ok {
		for _, t := range existingTags {
			if strTag, ok := t.(string); ok {
				tags = append(tags, strTag)
			}
		}
	}

	// check for duplicate tags

	for _, existingTag := range tags {
		if existingTag == tag {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Tag already exist",
			})
			return
		}
	}

	//add the new tag to the tags slice
	tags = append(tags, tag)
	data["tags"] = tags

	//marshal the updated data back to json
	updatedData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to marshal updated data",
		})
		return
	}

	err = r.Set(database.Ctx, shortId, updatedData, 0).Err()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update the database",
		})
	}

	//response with the updated data
	c.JSON(http.StatusOK, data)

}

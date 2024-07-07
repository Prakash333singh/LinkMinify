package routes

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Prakash333singh/url_shotner/api/database"
	"github.com/Prakash333singh/url_shotner/api/models"
	"github.com/Prakash333singh/url_shotner/api/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func ShortenURL(c *gin.Context) {
	var body models.Request

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot parse JSON"})
		return
	}

	r2 := database.CreateClient(1)
	defer r2.Close()

	val, err := r2.Get(database.Ctx, c.ClientIP()).Result()
	if err == redis.Nil {
		apiQuota := os.Getenv("API_QUOTA")
		if apiQuota == "" {
			apiQuota = "10"
		}

		// log.Printf("Setting rate limit for %s to %s requests", c.ClientIP(), apiQuota)

		err = r2.Set(database.Ctx, c.ClientIP(), apiQuota, 30*60*time.Second).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to set rate limit"})
			return
		}
		val = apiQuota
	} else {
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			limit, _ := r2.TTL(database.Ctx, c.ClientIP()).Result()
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error":            "rate limit exceeded",
				"rate_limit_reset": limit / time.Minute,
			})

			// log.Printf("Rate limit exceeded for %s, reset in %v minutes", c.ClientIP(), limit/time.Minute)

			return
		}

		// log.Printf("Rate limit for %s: %d requests remaining", c.ClientIP(), valInt)

	}

	if !govalidator.IsURL(body.URL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid URL"})
		return
	}

	if !utils.IsDifferentDomain(body.URL) {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "you can't hack this system :)"})
		return
	}

	body.URL = utils.EnsureHttPPrefix(body.URL)

	var id string
	if body.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShort
	}

	r := database.CreateClient(0)
	defer r.Close()

	val, _ = r.Get(database.Ctx, id).Result()
	if val != "" {
		c.JSON(http.StatusForbidden, gin.H{"error": "URL custom short already exists"})
		return
	}

	if body.Expiry == 0 {
		body.Expiry = 24
	}

	err = r.Set(database.Ctx, id, body.URL, body.Expiry*3600*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to connect to the Redis server"})
		return
	}

	resp := models.Response{
		Expiry:          body.Expiry,
		XRateLimitReset: 30,
		XRateRemaining:  10,
		URL:             body.URL,
		CustomShort:     os.Getenv("DOMAIN") + "/" + id,
	}

	r2.Decr(database.Ctx, c.ClientIP())

	val, _ = r2.Get(database.Ctx, c.ClientIP()).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)

	ttl, _ := r2.TTL(database.Ctx, c.ClientIP()).Result()
	resp.XRateLimitReset = ttl / time.Minute

	//log.Printf("Rate limit for %s after request: %d requests remaining, reset in %v minutes", c.ClientIP(), resp.XRateRemaining, resp.XRateLimitReset)

	c.JSON(http.StatusOK, resp)
}

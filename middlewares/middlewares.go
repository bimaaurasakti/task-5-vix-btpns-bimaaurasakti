package middlewares

import (
	"net/http"
	"strings"
	"vix-btpns/helpers"
	"vix-btpns/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")
		if !strings.Contains(bearerToken, "Bearer") {
			response := helpers.APIResponse("unauthorized token", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		stringToken := ""
		splitedBearerToken := strings.Split(bearerToken, " ")
		if len(splitedBearerToken) == 2 {
			stringToken = splitedBearerToken[1]
		}

		token, err := helpers.ValidateToken(stringToken)
		if err != nil {
			response := helpers.APIResponse("unauthorized token", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helpers.APIResponse("unauthorized token", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		var user models.User

		err = db.Where("id = ?", userID).Find(&user).Error
		if err != nil {
			response := helpers.APIResponse("unauthorized token", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}

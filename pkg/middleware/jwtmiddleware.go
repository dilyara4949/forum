package middleware

import (
	"fmt"
	"forum/internal/domain"
	"forum/internal/tokenutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		fmt.Println(t)
		if len(t) == 2 {
			authToken := t[1]
			
			authorized, err := tokenutil.IsAuthorized(authToken, secret)
			if err != nil {

				c.JSON(http.StatusUnauthorized, domain.ErrorResponse{
						Message: "Authorization error",
					})
				c.Abort()
				return
			}
			if authorized {

				email, err := tokenutil.ExtractIDFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, domain.ErrorResponse{

							Message: "Error to extract email from token" + err.Error(),
						})
					c.Abort()
					return
				}
				c.Set("x-email", email)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{

					Message: "Not authorized",
				})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{
			
				Message: "Authorization is required Header",
			})
		c.Abort()
	}
}
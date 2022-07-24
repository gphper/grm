/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-11-07 19:17:00
 */
package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		auth := c.Request.Header.Get("Authorization")
		info, err := jwt.ParseWithClaims(auth, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("TyPbWNRjho"), nil
		})
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "Auth Error",
			})
			c.Abort()
			return
		}

		uInfo := info.Claims.(jwt.MapClaims)

		expire, err := time.ParseInLocation("2006-01-02 15:04:05", uInfo["expire"].(string), time.Local)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "Auth Error",
			})
			c.Abort()
			return
		}

		if time.Until(expire).Seconds() < 0 {
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "Jwt Token Expired",
			})
			c.Abort()
			return
		}
		c.Set("username", uInfo["user"].(string))
		c.Next()
	}
}

/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-11-20 18:42:49
 */
package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

/**
* 静态资源缓存中间件
 */
func StaticCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()

		if strings.Contains(path, "/statics") {
			c.Writer.Header().Set("Cache-Control", "max-age=7200")
		}

		c.Next()
	}
}

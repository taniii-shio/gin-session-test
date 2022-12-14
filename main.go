package main

import (
	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


func main() {
    r := gin.Default()

		r.GET("/login", func(c *gin.Context) {
			sessionID := uuid.New().String()

			// c.SetSameSite(http.SameSiteNoneMode)
			
			c.SetCookie("sessionID", sessionID, 3600, "/", "", true, true)
			c.JSON(200, gin.H{ "message": "cookieをセットしました。"  })
		})

		r.GET("/callback", func(c *gin.Context) {
			cookie, err := c.Cookie("sessionID") 

			if err != nil {
				c.JSON(200, gin.H{ "message": "sessionIDを取得できません。" })
				return
			}

			c.JSON(200, gin.H{ "message": "sessionIDを取得しました。" + cookie })
		})
		
    r.Run()
}
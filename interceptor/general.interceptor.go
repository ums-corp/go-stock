package interceptor

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// GeneralIntercept - call this methos add interceptor
func GeneralIntercept(c *gin.Context)  {
	token := c.Query("token")
	if(token == "1234"){ 
		c.Next()
	} else { 
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"}) 
		c.Abort()
	}
}
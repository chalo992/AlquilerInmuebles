package common

import "github.com/gin-gonic/gin"

func ObtenerIdUsuarioClaims(c *gin.Context) (uint, bool) {
	idClaim, exists := c.Get("id")
	if !exists {
		return 0, false
	}

	id, ok := idClaim.(uint)
	return id, ok
}

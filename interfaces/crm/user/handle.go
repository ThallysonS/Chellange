package user

import (
	app "challenge/application/crm/user"

	"github.com/gin-gonic/gin"
)

func listarcadeiras(c *gin.Context) {

	res, err := app.ListarCadeiras(c)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)

}

package order

import "github.com/gin-gonic/gin"

func CadastrarCadeira(c *gin.Context) (interface{}, error) {
	var cadeira Cadeira

	if err := c.ShouldBindJSON(&cadeira); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return nil, err
	}
	return nil, nil
}

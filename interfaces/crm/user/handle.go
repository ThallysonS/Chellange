package user

import (
	order "challenge/application/crm/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func cadastrarcadeira(c *gin.Context) {
	var cadeira order.Cadeira

	if err := c.BindJSON(&cadeira); err != nil {
		c.JSON(400, gin.H{"erro": err.Error()})
		return
	}

	if err := order.CadastrarCadeira(cadeira); err != nil {
		c.JSON(500, gin.H{"erro": "erro ao cadastrar a cadeira no banco de dados, verifique o seu ID"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"mensagem": "cadeira cadastrada com sucesso"})
}

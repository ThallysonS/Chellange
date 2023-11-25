package user

import (
	"challenge/application/entries/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func cadastrarprofessor(c *gin.Context) {
	var professor user.Professor

	if err := c.BindJSON(&professor); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := user.CadastrarProfessor(professor); err != nil {
		c.JSON(500, gin.H{"error": "Erro ao cadastrar professor no banco de dados"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"menssagem": "Professor cadastrado com sucesso"})
}

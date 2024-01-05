package user

import (
	"challenge/application/entries/user"
	"challenge/config"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// função que cadastra o professor no banco de dados e envia email com a senha cadastrada (token)

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

	//Enviando email de confirmação

	err := config.CarregarVariaveisAmbiente()
	if err != nil {
		log.Fatal(err)
	}

	user.Configurar(config.SmtpServidor, config.SmtpPorta, config.SmtpUsuario, config.SmtpSenha)

	mensagem := fmt.Sprintf("Cadastro realizado com sucesso! SENHA do professor: %v", *professor.Senha)
	if err := user.EnviarEmailConfirmacao(*professor.Email, mensagem); err != nil {
		c.JSON(500, gin.H{"error": "Erro ao enviar e-mail de confirmação"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Mensagem": "Professor cadastrado"})

}

// função que cadastra o aluno no banco de dados

func cadastraraluno(c *gin.Context) {
	var aluno user.Aluno

	if err := c.BindJSON(&aluno); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := user.CadastrarAluno(aluno); err != nil {
		c.JSON(500, gin.H{"error": "erro ao cadastrar o aluno no banco de dados"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"menssagem": "aluno cadastrado com sucesso"})
}

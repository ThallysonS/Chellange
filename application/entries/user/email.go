package user

import (
	"net/smtp"
)

var (
	smtpServidor string
	smtpPorta    string
	smtpUsuario  string
	smtpSenha    string
)

//função de configurar váriável de ambiente para email

func Configurar(smtpServer, smtpPort, smtpUser, smtpPassword string) {
	smtpServidor = smtpServer
	smtpPorta = smtpPort
	smtpUsuario = smtpUser
	smtpSenha = smtpPassword
}

//função para enviar email

func EnviarEmailConfirmacao(paraEmail, mensagem string) error {
	aut := smtp.PlainAuth("", smtpUsuario, smtpSenha, smtpServidor)

	tipoConteudo := "text/plain; charset=UTF-8"
	corpo := mensagem

	mensagemEmail := []byte("Para: " + paraEmail + "\r\n" +
		"Assunto: Cadastro Realizado com Sucesso\r\n" +
		"Tipo de Conteúdo: " + tipoConteudo + "\r\n\r\n" +
		corpo + "\r\n")

	err := smtp.SendMail(smtpServidor+":"+smtpPorta, aut, smtpUsuario, []string{paraEmail}, mensagemEmail)
	if err != nil {
		return err
	}

	return nil
}

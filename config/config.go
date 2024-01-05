package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	SmtpServidor string
	SmtpPorta    string
	SmtpUsuario  string
	SmtpSenha    string
)

//função para carregar váriável de ambiente

func CarregarVariaveisAmbiente() error {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
		return err
	}

	SmtpServidor = os.Getenv("SMTP_SERVIDOR")
	SmtpPorta = os.Getenv("SMTP_PORTA")
	SmtpUsuario = os.Getenv("SMTP_USUARIO")
	SmtpSenha = os.Getenv("SMTP_SENHA")

	return nil
}

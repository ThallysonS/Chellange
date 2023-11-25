package postgres

import (
	"challenge/domain/entries/user"
	"database/sql"
)

func CadastrarProfNoBanco(tx *sql.Tx, p user.Professor) error {
	_, err := tx.Exec("INSERT INTO professor (idprof, nome, telefone, email, cpf) VALUES ($1, $2, $3, $4, $5)",
		p.IDProf, p.Nome, p.Telefone, p.Email, p.CPF)
	if err != nil {
		return err
	}

	return nil
}

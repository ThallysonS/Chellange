package postgres

import (
	"challenge/domain/crm/order"
	"challenge/domain/entries/user"
	"database/sql"
	"log"
)

func CadastrarProfNoBanco(tx *sql.Tx, p user.Professor) error {
	_, err := tx.Exec("INSERT INTO professor (nome, telefone, email, cpf) VALUES ($1, $2, $3, $4)",
		p.Nome, p.Telefone, p.Email, p.CPF)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func CadastrarAlunoNoBanco(tx *sql.Tx, a user.Aluno) error {
	_, err := tx.Exec("INSERT INTO aluno (nome, cpf, email) VALUES ($1, $2, $3)",
		a.Nome, a.CPF, a.Email)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func CadastrarCadeiraNoBanco(tx *sql.Tx, c order.Cadeira) error {
	_, err := tx.Exec("INSERT INTO cadeira (nome, slug, data_inicio, data_fim, carga_horaria, idprof) VALUES ($1, $2, $3, $4, $5, $6)",
		c.Nome, c.Slug, c.DataInicio, c.DataFim, c.CargaHoraria, c.IDProf)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

package user

import (
	"challenge/config/database"
	"challenge/domain/entries/user"
	"challenge/infrastructure/crm/order/postgres"

	"fmt"
)

func CadastrarProfessor(professor Professor) error {
	tx, err := database.NovaTX()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	err = postgres.CadastrarProfNoBanco(tx, (user.Professor(professor)))
	if err != nil {
		return err
	}

	fmt.Println("Professor registrado no banco de dados")

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func CadastrarCadeira() {

}

func CadastrarAluno() {

}

func ListarCadeiras() {

}

func ListarAlunosCadeiras() {

}

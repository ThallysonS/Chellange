package order

import (
	"challenge/config/database"
	"challenge/domain/crm/order"
	"challenge/infrastructure/crm/order/postgres"
	"fmt"
)

func CadastrarCadeira(cadeira Cadeira) error {
	tx, err := database.NovaTX()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = postgres.CadastrarCadeiraNoBanco(tx, (order.Cadeira(cadeira)))
	if err != nil {
		return err
	}

	fmt.Println("cadeira cadastrada com sucesso")

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil

}

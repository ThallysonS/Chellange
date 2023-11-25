package order

import "time"

type Cadeira struct {
	Nomecadeira  *string    `json:"nomecadeira"`
	Slug         *string    `json:"slug"`
	DataInicio   *time.Time `json:"datainicio"`
	DataFinal    *time.Time `json:"datafinal"`
	CargaHoraria *string    `json:"cargahoraria"`
	IDCadeira    *int       `json:"idcadeira"`
	IDProf       *int64     `json:"idprof"`
}

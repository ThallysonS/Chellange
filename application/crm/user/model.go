package order

import (
	"time"
)

type Cadeira struct {
	ID           int
	Nome         string
	Slug         string
	DataInicio   time.Time
	DataFim      time.Time
	CargaHoraria int
	IDProf       int
}

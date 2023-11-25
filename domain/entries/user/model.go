package user

type Professor struct {
	IDProf   *int64  `json:"idprof,omitempty"`
	Nome     *string `json:"nome,omitempty"`
	Telefone *string `json:"telefone,omitempty"`
	Email    *string `json:"email,omitempty"`
	CPF      *string `json:"cpf,omitempty"`
}

type Aluno struct {
	Nome      *string `json:"nome"`
	CPF       *string `json:"cpf"`
	Email     *string `json:"email"`
	Matricula *int    `json:"matricula"`
}

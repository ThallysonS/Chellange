package user

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {

	//!aluno lista as cadeiras disponíveis
	r.POST("/cadastrarcadeira", cadastrarcadeira)
	//!professor lista as solicitações de alunos sobre cadeiras disponíveis
	//r.GET("/listarsolicitacoes", listarsolicitacoes)
	//!o sistema lista as cadeiras e seus alunos matriculados
	//r.GET("/listaralunoscadeiras", listarcadeirasealunos)
	//!o Professor cadastra a cadeira
	//	r.POST("/cadastrarcadeira", cadastrarcadeira)

}

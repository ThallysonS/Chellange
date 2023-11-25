package user

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	//!rota reposável por cadastrar professor
	r.POST("/cadastrarprofessor", cadastrarprofessor)
	//!rota responsável por deletar professor
	//r.DELETE("/:deletarprofessor", deletarprofessor)
	//!rota responsável por cadastrar professor
	//r.POST("/cadastraraluno", cadastraraluno)
	//!rota responsável por deletar aluno
	//r.DELETE("/:deletaraluno", deletaraluno)

}

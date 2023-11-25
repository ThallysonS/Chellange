package crm

import (
	"challenge/interfaces/crm/user"

	"github.com/gin-gonic/gin"
)

func Rotas(r *gin.RouterGroup) {
	user.Router(r.Group("/rotas"))
}

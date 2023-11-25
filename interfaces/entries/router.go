package entries

import (
	"challenge/interfaces/entries/user"

	"github.com/gin-gonic/gin"
)

func Rotas(r *gin.RouterGroup) {
	user.Router(r.Group("/rotas"))
}

package controller

import (
	"net/http"

	"github.com/JuliaTeles/golang/model"
	"github.com/JuliaTeles/golang/repository"
	"github.com/gin-gonic/gin"
)

func PostAluno(c *gin.Context) error {
	var usuario model.Usuario
	if err := c.BindJSON(&usuario); err != nil {
		return err
	}

	if err := repository.InsertUsuario(usuario); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.String(http.StatusOK, "Deu tudo certo")
	return nil
}

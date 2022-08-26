package controller

import (
	"github.com/JuliaTeles/golang/model"
	"github.com/gin-gonic/gin"
)

func PostAluno(c *gin.Context) error {
	var usuario model.Usuario
	if err := c.BindJSON(&usuario); err !: nil{
		return err
	}
}

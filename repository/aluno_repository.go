package repository

import (
	"github.com/JuliaTeles/golang/database"
	"github.com/JuliaTeles/golang/model"
)

func InsertUsuario(usuario model.Usuario) error {
	db := database.GrabDB()

	statement := db.MustBegin()
	statement.MustExec("INSERT INTO usuarios VALUES($1,$2)", usuario.Nome, usuario.RA)
	if err := statement.Commit(); err != nil {
		return err
	}
	return nil
}

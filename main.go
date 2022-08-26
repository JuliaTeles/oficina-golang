package main

import (
	"fmt"
	"os"

	"github.com/JuliaTeles/golang/database"
	"github.com/gin-gonic/gin"
)

func main() {

	if err := database.Connect(); err != nil {
		panic(err)
	}

	db := database.GrabDB()
	st := db.MustBegin() // prepara um comando
	sql := "CREATE TABLE IF NOT EXISTS usuarios (nome VARCHAR(128), ra VARCHAR(6))"
	if err := st.MustExec(sql); err != nil{
		panic(err)
	}

	r := gin.Default()
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}

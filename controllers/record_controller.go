package controllers

import (
	"servercg/db"
	"encoding/json"
	"github.com/kataras/iris/v12"
)

type Player struct {
	ID       	int 	`json:"id,omitempty"`
	Nivel       int 	`json:"nivel,omitempty"`
	Pontuacao   int 	`json:"pontuacao,omitempty"`
	Username 	string 	`json:"username,omitempty"`
	Password 	string	`json:"password,omitempty"`
	Nickname 	string	`json:"nickname,omitempty"`
}

func Insert() func(ctx iris.Context) {
	var p Player
	return func(ctx iris.Context) {
		body, _ := ctx.GetBody()
		a := json.Unmarshal(body, &p)
		if a != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"Status": "Fail", "Errors": "Erro na estrutura de envio da transação."})
			return
		}

		db := db.ConnectionBDPostgreSQL()
		const SCRIPT = `INSERT INTO player (nivel,pontuacao,login,senha,nickname)
								  VALUES ($1,$2,$3,$4,$5) RETURNING id`
		rows, err := db.Query(SCRIPT,
			p.Nivel, p.Pontuacao, p.Username, p.Password, p.Nickname)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"Status": "Fail", "Errors": "Erro na estrutura de envio da transação."})
			return
		}

		defer rows.Close()
		defer db.Close()

		ctx.StatusCode(iris.StatusAccepted)
		ctx.JSON(iris.Map{"Status": "Success", "Errors": ""})
		return

	}
}

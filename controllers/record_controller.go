package controllers

import (
	"servercg/utils"
	"encoding/json"
	"github.com/kataras/iris/v12"
)

type Player struct {
	ID       	int 	`json:"id"`
	Nivel       int 	`json:"nivel"`
	Pontuacao   int 	`json:"pontuacao"`
	Username 	string 	`json:"username"`
	Password 	string	`json:"password"`
	Nickname 	string	`json:"nickname"`
}

func Insert() func(ctx iris.Context) {
	var p Player
	return func(ctx iris.Context) {
		body, _ := ctx.GetBody()
		a := json.Unmarshal(body, &p)
		if a != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"Status": "Fail-S", "Errors": "Erro na estrutura de envio da transação."})
			return
		}

		// USER COM PELO MENOS 5 CARACT
		if len(p.Username) < 5{
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"Status": "Fail-U", "Errors": "Username precisa ter pelo menos 5 caracteres."})
			return
		}

		// PASS COM PELO MENOS 6 CARACT
		if len(p.Password) < 6{
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"Status": "Fail-P", "Errors": "Password precisa ter pelo menos 6 caracteres."})
			return
		}

		// NICK COM PELO MENOS 5 CARACT
		if len(p.Nickname) < 5{
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"Status": "Fail-N", "Errors": "Nickname precisa ter pelo menos 5 caracteres."})
			return
		}

		db := utils.ConnectionBDPostgreSQL()
		const SCRIPT = `INSERT INTO player (nivel,pontuacao,login,senha,nickname)
								  VALUES ($1,$2,$3,$4,$5) RETURNING id`
		rows, err := db.Query(SCRIPT,
			p.Nivel, p.Pontuacao, p.Username, p.Password, p.Nickname)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"Status": "Fail-C", "Errors": "Erro no insert Record.Insert()"})
			return
		}

		defer rows.Close()
		defer db.Close()

		ctx.StatusCode(iris.StatusAccepted)
		ctx.JSON(iris.Map{"Status": "Success", "Errors": ""})
		return

	}
}

func Update() func(ctx iris.Context) {
	var p Player
	return func(ctx iris.Context) {
		body, _ := ctx.GetBody()
		a := json.Unmarshal(body, &p)
		if a != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"Status": "Fail-S", "Errors": "Erro na estrutura de envio da transação."})
			return
		}

		// USER COM PELO MENOS 5 CARACT
		if p.Username == ""{
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"Status": "Fail-U", "Errors": "Username obrigatório."})
			return
		}

		db := utils.ConnectionBDPostgreSQL()
		const SCRIPT = `UPDATE player SET pontuacao = $1 
								  WHERE login = $2`
		rows, err := db.Query(SCRIPT,
			p.Pontuacao, p.Username)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"Status": "Fail-C", "Errors": "Erro no update Record.Update()"})
			return
		}

		defer rows.Close()
		defer db.Close()

		ctx.StatusCode(iris.StatusAccepted)
		ctx.JSON(iris.Map{"Status": "Success", "Errors": ""})
		return

	}
}

func List() func(ctx iris.Context) {
	return func(ctx iris.Context) {
		var list []Player
		db := utils.ConnectionBDPostgreSQL()
		const SCRIPT = `SELECT id, nivel, pontuacao, login, senha, nickname
							FROM public.player ORDER BY pontuacao DESC`
		rows, err := db.Query(SCRIPT)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"Status": "Fail-C", "Errors": "Erro no update Record.Update()"})
			return
		}

		for rows.Next() {
			var p Player
			err = rows.Scan(&p.ID,&p.Nivel,&p.Pontuacao,&p.Username,&p.Password,&p.Nickname)
			if err != nil {
				ctx.StatusCode(iris.StatusBadRequest)
				ctx.JSON(iris.Map{"Status": "Fail-C", "Errors": "Erro no update Record.Update()"})
				return
			}

			list = append(list, p)
		}

		defer rows.Close()
		defer db.Close()

		ctx.StatusCode(iris.StatusAccepted)
		ctx.JSON(iris.Map{"Status": "Success", "Errors": "", "Records" : list})
		return

	}
}
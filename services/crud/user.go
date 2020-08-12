package crud

import (
	"log"
	"web-zjh/models"
	. "web-zjh/settings"
)

func InsertOneUser(user models.User) (id int, err error) {
	row := DB.QueryRow(
		`INSERT INTO "user"(username,password,telephone,token)values($1,$2,$3,$4) RETURNING id`, user.Username,
		user.Password, user.Telephone, user.Token,
	)
	err = row.Scan(&id)
	if err != nil {
		log.Println("crud.InsertOneUser:", err)
	}
	return
}

func FindUserById(id int) (user models.User) {
	//rows, err := DB.Query(`SELECT id,username,token FROM "user" WHERE id=$1`, id)
	row := DB.QueryRow(`SELECT id,username,token FROM "user" WHERE id=$1`, id)
	err := row.Scan(&user.Id, &user.Username, &user.Token)
	if err != nil {
		log.Println("crud.FindUserById.row.Scan:", err)
	}
	return
}

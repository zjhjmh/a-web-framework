package sql

import (
	"log"
	"web-zjh/models"
	"web-zjh/settings"
)


func FindUserById(id int) (userList []models.User){
	rows, err := settings.DB.Query("SELECT id,name FROM auth_group WHERE id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	res := models.User{}
	for rows.Next() {
		err = rows.Scan(&res.Id, &res.Username)
		if err != nil {
			log.Fatal(err)
		}
		userList = append(userList, res)
	}
	log.Print(userList)
	if len(userList) >= 1 {
		log.Print(userList[0].Username)
	}
	return
}




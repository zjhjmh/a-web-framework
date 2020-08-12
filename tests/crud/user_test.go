package crud

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"web-zjh/models"
	"web-zjh/services/crud"
	"web-zjh/settings"
)

func TestMain(m *testing.M) {
	settings.ConnectTestDB()
	file, err := ioutil.ReadFile("C:\\Users\\zjh\\Desktop\\web-zjh\\main.sql")
	if err != nil {
		log.Println(err)
	}
	_, err = settings.DB.Exec(string(file))
	if err != nil {
		log.Println(err)
	}

	exitCode := m.Run()

	_ = settings.DB.Close()
	os.Exit(exitCode)
}

func TestFindUserById(t *testing.T) {
	userWant := models.User{
		Username:  "test",
		Password:  "password",
		Telephone: "110",
		Token:     "klfasjd;lfk",
	}
	userId, _ := crud.InsertOneUser(userWant)
	userReal := crud.FindUserById(userId)
	if userReal.Username != userWant.Username {
		t.Errorf("userReal:%q != userWant:%q", userReal, userWant)
	}
}

package crud

import (
	"testing"
	"web-zjh/models"
	"web-zjh/services/crud"
)

func TestFindUserById2(t *testing.T) {
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

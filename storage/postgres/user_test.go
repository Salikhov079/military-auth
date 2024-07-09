package postgres

import (
	"log"
	"testing"

	pb "root/genprotos"

	"github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {
	stg, err := NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	user := &pb.LoginUserRequest{
		UserName: "dilshod",
		Password: "1234",
	}
	assert.NoError(t, err)
	userLogin, err := stg.User().LoginUser(user)

	assert.NoError(t, err)

	assert.Equal(t, userLogin.UserName, user.UserName)

}

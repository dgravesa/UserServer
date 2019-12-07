package controller

import (
	"os"
	"testing"

	"github.com/dgravesa/WaterLogger-UserServer/data"
	"github.com/dgravesa/WaterLogger-UserServer/model"
)

var testUserData = []model.User{
	model.User{Name: "Jeremy1"},
	model.User{Name: "WaterDrinkerGuy"},
}

func TestMain(m *testing.M) {
	userData := data.NewInMemoryUserStore()

	for _, user := range testUserData {
		userData.Insert(user)
	}

	model.SetUserDataLayer(userData)

	os.Exit(m.Run())
}

func Test_GetUserByName_WhenNameExists_ReturnsCorrectUser(t *testing.T) {

}

func Test_GetUserByName_WhenNameDoesNotExist_ReturnsNotFound(t *testing.T) {

}

func Test_GetUserByID_WhenIDExists_ReturnsCorrectUser(t *testing.T) {

}

func Test_GetUserByID_WhenIDDoesNotExist_ReturnsNotFound(t *testing.T) {

}

func Test_GetUser_WithoutQuery_ReturnsBadRequest(t *testing.T) {

}

func Test_PostUser_WithNewName_ReturnsNewUserURL(t *testing.T) {

	// TODO verify user at URL matches expected
}

func Test_PostUser_WithExistingName_ReturnsConflict(t *testing.T) {

}

func Test_PostUser_WithoutName_ReturnsBadRequest(t *testing.T) {

}

func Test_DeleteUser_WhenIDExists_ReturnsNotFoundOnSubsequentGet(t *testing.T) {

}

func Test_DeleteUser_WhenIDDoesNotExist_ReturnsNotFound(t *testing.T) {

}

func Test_DeleteUser_WithoutID_ReturnsBadRequest(t *testing.T) {

}

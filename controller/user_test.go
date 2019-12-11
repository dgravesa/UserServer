package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgravesa/WaterLogger-UserServer/data"
	"github.com/dgravesa/WaterLogger-UserServer/model"
)

var testUserData = []model.User{
	model.User{Name: "Jeremy1"},
	model.User{Name: "WaterDrinkerGuy1000"},
}

func initTestUserData(testData []model.User) {
	userData := data.NewInMemoryUserStore()

	for _, user := range testUserData {
		userData.Insert(user)
	}

	model.SetUserDataLayer(userData)
}

func Test_GetUserByName_WhenNameExists_ReturnsCorrectUser(t *testing.T) {
	// Arrange
	initTestUserData(testUserData)
	expectedCode := http.StatusOK
	userName := "WaterDrinkerGuy1000"
	target := fmt.Sprintf("http://localhost/user?name=%s", userName)
	req := httptest.NewRequest("GET", target, nil)
	res := httptest.NewRecorder()

	// Act
	userHandler(res, req)

	// Assert
	if res.Code != expectedCode {
		t.Errorf("expected status code = %d, received status code = %d", expectedCode, res.Code)
	} else {
		var user model.User
		if err := json.Unmarshal(res.Body.Bytes(), &user); err != nil {
			t.Errorf("error decoding user: %s", err)
		} else if userName != user.Name {
			t.Errorf("expected user name = %s, received user name = %s", userName, user.Name)
		}
	}
}

func Test_GetUserByName_WhenNameDoesNotExist_ReturnsNotFound(t *testing.T) {
	// Arrange
	expectedCode := http.StatusNotFound
	userName := "NotARealUser"
	target := fmt.Sprintf("http://localhost/user?name=%s", userName)
	req := httptest.NewRequest("GET", target, nil)
	res := httptest.NewRecorder()

	// Act
	userHandler(res, req)

	// Assert
	if res.Code != expectedCode {
		t.Errorf("expected status code = %d, received status code = %d", expectedCode, res.Code)
	}
}

func Test_GetUserByID_WhenIDExists_ReturnsCorrectUser(t *testing.T) {
	// TODO implement
	t.SkipNow()
}

func Test_GetUserByID_WhenIDDoesNotExist_ReturnsNotFound(t *testing.T) {
	// TODO implement
	t.SkipNow()
}

func Test_GetUser_WithoutQuery_ReturnsBadRequest(t *testing.T) {
	// TODO implement
	t.SkipNow()
}

func Test_PostUser_WithNewName_ReturnsNewUserURL(t *testing.T) {
	// TODO implement
	t.SkipNow()

	// TODO verify user at URL matches expected
}

func Test_PostUser_WithExistingName_ReturnsConflict(t *testing.T) {
	// TODO implement
	t.SkipNow()
}

func Test_PostUser_WithoutName_ReturnsBadRequest(t *testing.T) {
	// TODO implement
	t.SkipNow()
}

func Test_DeleteUser_WhenIDExists_ReturnsNotFoundOnSubsequentGet(t *testing.T) {
	// TODO implement
	t.SkipNow()
}

func Test_DeleteUser_WhenIDDoesNotExist_ReturnsNotFound(t *testing.T) {
	// TODO implement
	t.SkipNow()
}

func Test_DeleteUser_WithoutID_ReturnsBadRequest(t *testing.T) {
	// TODO implement
	t.SkipNow()
}

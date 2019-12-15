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
	model.User{Name: "AlphaSquid13"},
	model.User{Name: "BetaDolphin2"},
}

func initTestUserData(testData []model.User) {
	userData := data.NewInMemoryUserStore()

	for _, user := range testUserData {
		userData.Insert(user)
	}

	model.SetUserDataLayer(userData)
}

func checkResponseCode(expected, received int, t *testing.T) {
	if expected != received {
		t.Errorf("expected status code = %d, received status code = %d", expected, received)
	}
}

func Test_GetUserByName_WhenNameExists_ReturnsCorrectUser(t *testing.T) {
	// Arrange
	initTestUserData(testUserData)
	expectedCode := http.StatusOK
	userName := "WaterDrinkerGuy1000"
	target := fmt.Sprintf("http://localhost/users?name=%s", userName)
	req := httptest.NewRequest("GET", target, nil)
	res := httptest.NewRecorder()

	// Act
	userHandler(res, req)

	// Assert
	checkResponseCode(expectedCode, res.Code, t)
	var user model.User
	if err := json.Unmarshal(res.Body.Bytes(), &user); err != nil {
		t.Errorf("error decoding user: %s", err)
	} else if userName != user.Name {
		t.Errorf("expected user name = %s, received user name = %s", userName, user.Name)
	}
}

func Test_GetUserByName_WhenNameDoesNotExist_ReturnsNotFound(t *testing.T) {
	// Arrange
	initTestUserData(testUserData)
	expectedCode := http.StatusNotFound
	userName := "NotARealUser"
	target := fmt.Sprintf("http://localhost/users?name=%s", userName)
	req := httptest.NewRequest("GET", target, nil)
	res := httptest.NewRecorder()

	// Act
	userHandler(res, req)

	// Assert
	checkResponseCode(expectedCode, res.Code, t)
	var user model.User
	if err := json.Unmarshal(res.Body.Bytes(), &user); err == nil {
		t.Errorf("user received when no response should have been given")
	}
}

func Test_GetUserByID_WhenIDExists_ReturnsCorrectUser(t *testing.T) {
	// Arrange
	initTestUserData(testUserData)
	expectedCode := http.StatusOK
	userName := "AlphaSquid13"
	userID := uint64(2) // expected index from in-memory user store
	target := fmt.Sprintf("http://localhost/users?id=%d", userID)
	req := httptest.NewRequest("GET", target, nil)
	res := httptest.NewRecorder()

	// Act
	userHandler(res, req)

	// Assert
	checkResponseCode(expectedCode, res.Code, t)
	var user model.User
	if err := json.Unmarshal(res.Body.Bytes(), &user); err != nil {
		t.Errorf("error decoding user: %s", err)
	} else if user.Name != userName || user.ID != userID {
		t.Errorf("expected name/id: %s/%d, received name/id: %s/%d", userName, userID, user.Name, user.ID)
	}
}

func Test_GetUserByID_WhenIDDoesNotExist_ReturnsNotFound(t *testing.T) {
	// Arrange
	initTestUserData(testUserData)
	expectedCode := http.StatusNotFound
	userID := uint64(101)
	target := fmt.Sprintf("http://localhost/users?id=%d", userID)
	req := httptest.NewRequest("GET", target, nil)
	res := httptest.NewRecorder()

	// Act
	userHandler(res, req)

	// Assert
	checkResponseCode(expectedCode, res.Code, t)
	var user model.User
	if err := json.Unmarshal(res.Body.Bytes(), &user); err == nil {
		t.Errorf("user received when no response should have been given")
	}
}

func Test_GetUserByIDAndName_WithNameAndIDMatch_ReturnsCorrectUser(t *testing.T) {
	// Arrange
	initTestUserData(testUserData)
	expectedCode := http.StatusOK
	userID := uint64(3)
	userName := "BetaDolphin2"
	target := fmt.Sprintf("http://localhost/users?id=%d&name=%s", userID, userName)
	req := httptest.NewRequest("GET", target, nil)
	res := httptest.NewRecorder()

	// Act
	userHandler(res, req)

	// Assert
	checkResponseCode(expectedCode, res.Code, t)
	var user model.User
	if err := json.Unmarshal(res.Body.Bytes(), &user); err != nil {
		t.Errorf("error decoding user: %s", err)
	} else if user.Name != userName || user.ID != userID {
		t.Errorf("expected name/id: %s/%d, received name/id: %s/%d", userName, userID, user.Name, user.ID)
	}
}

func Test_GetUserByIDAndName_WithNameAndIDMismatch_ReturnsNotFound(t *testing.T) {
	// Arrange
	initTestUserData(testUserData)
	expectedCode := http.StatusNotFound
	userID := uint64(3)
	userName := "Jeremy1"
	target := fmt.Sprintf("http://localhost/users?id=%d&name=%s", userID, userName)
	req := httptest.NewRequest("GET", target, nil)
	res := httptest.NewRecorder()

	// Act
	userHandler(res, req)

	// Assert
	checkResponseCode(expectedCode, res.Code, t)
	var user model.User
	if err := json.Unmarshal(res.Body.Bytes(), &user); err == nil {
		t.Errorf("user received when no response should have been given")
	}
}

func Test_GetUser_WithoutQuery_ReturnsBadRequest(t *testing.T) {
	// Arrange
	initTestUserData(testUserData)
	expectedCode := http.StatusBadRequest
	target := "http://localhost/users"
	req := httptest.NewRequest("GET", target, nil)
	res := httptest.NewRecorder()

	// Act
	userHandler(res, req)

	// Assert
	checkResponseCode(expectedCode, res.Code, t)
	var user model.User
	if err := json.Unmarshal(res.Body.Bytes(), &user); err == nil {
		t.Errorf("user received when no response should have been given")
	}
}

func Test_PostUser_WithNewName_ReturnsNewUserURL(t *testing.T) {
	// Arrange
	initTestUserData(testUserData)
	expectedCode := http.StatusCreated
	newUserName := "H2Oliver"
	target := fmt.Sprintf("http://localhost/users?name=%s", newUserName)
	req := httptest.NewRequest("POST", target, nil)
	rec := httptest.NewRecorder()

	// Act
	userHandler(rec, req)

	// Assert
	res := rec.Result()
	checkResponseCode(expectedCode, res.StatusCode, t)
	loc, err := res.Location()
	if err != nil {
		t.Fatal(err)
	}
	verifyReq := httptest.NewRequest("GET", loc.String(), nil)
	verifyRec := httptest.NewRecorder()
	userHandler(verifyRec, verifyReq)
	checkResponseCode(http.StatusOK, verifyRec.Code, t)
}

func Test_PostUser_WithExistingName_ReturnsConflict(t *testing.T) {
	// Arrange
	initTestUserData(testUserData)
	expectedCode := http.StatusConflict
	newUserName := "WaterDrinkerGuy1000"
	target := fmt.Sprintf("http://localhost/users?name=%s", newUserName)
	req := httptest.NewRequest("POST", target, nil)
	rec := httptest.NewRecorder()

	// Act
	userHandler(rec, req)

	// Assert
	res := rec.Result()
	checkResponseCode(expectedCode, res.StatusCode, t)
}

func Test_PostUser_WithoutName_ReturnsBadRequest(t *testing.T) {
	// Arrange
	initTestUserData(testUserData)
	expectedCode := http.StatusBadRequest
	target := "http://localhost/users"
	req := httptest.NewRequest("POST", target, nil)
	rec := httptest.NewRecorder()

	// Act
	userHandler(rec, req)

	// Assert
	res := rec.Result()
	checkResponseCode(expectedCode, res.StatusCode, t)
}

func Test_DeleteUser_WhenIDExists_ReturnsSuccessAndHasDeletedUser(t *testing.T) {
	// Arrange
	initTestUserData(testUserData)
	expectedCode := http.StatusOK
	userID := 2
	userName := "AlphaSquid13"
	target := fmt.Sprintf("http://localhost/users?id=%d", userID)
	req := httptest.NewRequest("DELETE", target, nil)
	res := httptest.NewRecorder()

	// Act
	userHandler(res, req)

	// Assert
	checkResponseCode(expectedCode, res.Code, t)
	if _, found := model.FindUserByName(userName); found {
		t.Error("user found in data after delete")
	}
}

func Test_DeleteUser_WhenIDDoesNotExist_ReturnsNotFound(t *testing.T) {
	// Arrange
	initTestUserData(testUserData)
	expectedCode := http.StatusNotFound
	userID := 101
	target := fmt.Sprintf("http://localhost/users?id=%d", userID)
	req := httptest.NewRequest("DELETE", target, nil)
	res := httptest.NewRecorder()

	// Act
	userHandler(res, req)

	// Assert
	checkResponseCode(expectedCode, res.Code, t)
}

func Test_DeleteUser_WithoutID_ReturnsBadRequest(t *testing.T) {
	// Arrange
	initTestUserData(testUserData)
	expectedCode := http.StatusBadRequest
	target := "http://localhost/users"
	req := httptest.NewRequest("DELETE", target, nil)
	res := httptest.NewRecorder()

	// Act
	userHandler(res, req)

	// Assert
	checkResponseCode(expectedCode, res.Code, t)
}

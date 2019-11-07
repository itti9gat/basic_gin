package handler_test

import (
	"bytes"
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"iiujapp.tech/basic-gin/model"
	"iiujapp.tech/basic-gin/server"
)

type mockTest struct {
	DBData *sql.DB
}

func (s *mockTest) QueryUser() (model.ListUser, error) {

	data := model.ListUser{model.User{UserID: 1, Username: "testuser", Password: "pass1234", Name: "test", Status: "ADMIN"}, model.User{UserID: 2, Username: "demo2", Password: "demo1234", Name: "demo", Status: "USER"}}
	return data, nil
}

func (s *mockTest) WriteData(m model.User) error {
	return nil
}

type mockFailTest struct {
	DBData *sql.DB
}

func (s *mockFailTest) QueryUser() (model.ListUser, error) {
	return model.ListUser{}, errors.New("Error 1049: Unknown database 'dbname'")
}

func (s *mockFailTest) WriteData(m model.User) error {
	return errors.New("Error 1049: Unknown database 'dbname'")
}

func TestUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db, _, _ := sqlmock.New()
	defer db.Close()

	s := &mockTest{
		DBData: db,
	}

	serv := server.NewServer(s)
	r := serv.SetupRouter()

	req, _ := http.NewRequest("GET", "/user", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expect := "{\"code\":200,\"message\":[{\"user_id\":1,\"username\":\"testuser\",\"password\":\"pass1234\",\"name\":\"test\",\"status\":\"ADMIN\"},{\"user_id\":2,\"username\":\"demo2\",\"password\":\"demo1234\",\"name\":\"demo\",\"status\":\"USER\"}]}"

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expect, w.Body.String())
}

func TestSaveUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db, _, _ := sqlmock.New()
	defer db.Close()

	s := &mockTest{
		DBData: db,
	}

	serv := server.NewServer(s)
	r := serv.SetupRouter()

	user := "{\"username\":\"9demo\",\"password\":\"112233\",\"name\":\"demo\",\"status\":\"A\"}"
	param := []byte(user)

	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(param))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expect := "{\"code\":200,\"message\":true}"

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expect, w.Body.String())
}

func TestSaveUserHandlerFail(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db, _, _ := sqlmock.New()
	defer db.Close()

	s := &mockTest{
		DBData: db,
	}

	serv := server.NewServer(s)
	r := serv.SetupRouter()

	user := "{\"username\":\"9demo\",\"password:\"112233\",\"name\":\"demo\",\"status\":\"A\"}"
	param := []byte(user)

	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(param))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expect := "{\"error\":\"invalid character '1' after object key\"}"

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, expect, w.Body.String())
}

func TestUserDBFailHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db, _, _ := sqlmock.New()
	defer db.Close()

	s := &mockFailTest{
		DBData: db,
	}

	serv := server.NewServer(s)
	r := serv.SetupRouter()

	req, _ := http.NewRequest("GET", "/user", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expect := "{\"error\":\"Error 1049: Unknown database 'dbname'\"}"

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, expect, w.Body.String())
}

func TestSaveUserDBFailHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db, _, _ := sqlmock.New()
	defer db.Close()

	s := &mockFailTest{
		DBData: db,
	}

	serv := server.NewServer(s)
	r := serv.SetupRouter()

	user := "{\"username\":\"9demo\",\"password\":\"112233\",\"name\":\"demo\",\"status\":\"A\"}"
	param := []byte(user)

	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(param))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	expect := "{\"error\":\"Error 1049: Unknown database 'dbname'\"}"

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, expect, w.Body.String())
}

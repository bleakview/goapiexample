package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bleakview/goapiexample/cmd/goapiexample/app"
	"github.com/bleakview/goapiexample/cmd/goapiexample/initdb"
	"github.com/bleakview/goapiexample/cmd/goapiexample/model"
	"github.com/stretchr/testify/require"
)

var a *app.App

//Main test function
func TestMain(m *testing.M) {
	a = &app.App{}
	a.Initialize()
	code := m.Run()
	os.Exit(code)
}

//Test if we get empty results with empty table
func TestEmptyTable(t *testing.T) {
	initdb.CleanDB(a.DB)
	req, _ := http.NewRequest("GET", "/books", nil)
	response := executeRequest(req, a)
	checkResponseCode(t, http.StatusOK, response.Code)
	body := response.Body.String()
	require.JSONEq(t, "[]", body)
}

//Check if we get default records with initilize table
func TestGetRecords(t *testing.T) {
	initdb.InitDB(a.DB)
	req, _ := http.NewRequest("GET", "/books", nil)
	response := executeRequest(req, a)
	checkResponseCode(t, http.StatusOK, response.Code)
	body := response.Body.String()
	sampleDbData := `[{
		"id": "5d596c01-e20b-4049-91e9-a0be77715260",
		"name": "name",
		"author": "author",
		"release_year": 1980,
		"isbn": "isbn"
	  }]`
	require.JSONEq(t, sampleDbData, body)
}

//Check if we get default record with id
func TestGetRecord(t *testing.T) {
	initdb.InitDB(a.DB)
	req, _ := http.NewRequest("GET", "/books/5d596c01-e20b-4049-91e9-a0be77715260", nil)
	response := executeRequest(req, a)
	checkResponseCode(t, http.StatusOK, response.Code)
	body := response.Body.String()
	sampleDbData := `{
		"id": "5d596c01-e20b-4049-91e9-a0be77715260",
		"name": "name",
		"author": "author",
		"release_year": 1980,
		"isbn": "isbn"
	  }`
	require.JSONEq(t, sampleDbData, body)
}

// check if we can update record
func TestUpdateRecord(t *testing.T) {
	initdb.InitDB(a.DB)
	var jsonData = []byte(`{
		"name": "name test",
		"author": "author test",
		"release_year": 1981,
		"isbn": "isbn test"
	}`)
	req, _ := http.NewRequest("PUT", "/books/5d596c01-e20b-4049-91e9-a0be77715260", bytes.NewBuffer(jsonData))
	response := executeRequest(req, a)
	checkResponseCode(t, http.StatusOK, response.Code)
	req, _ = http.NewRequest("GET", "/books/5d596c01-e20b-4049-91e9-a0be77715260", nil)
	response = executeRequest(req, a)
	checkResponseCode(t, http.StatusOK, response.Code)
	body := response.Body.String()
	sampleDbData := `{
		"id": "5d596c01-e20b-4049-91e9-a0be77715260",
		"name": "name test",
		"author": "author test",
		"release_year": 1981,
		"isbn": "isbn test"
	}`
	require.JSONEq(t, sampleDbData, body)
}

//check if we can create record
func TestCreateRecord(t *testing.T) {
	initdb.InitDB(a.DB)
	var jsonData = []byte(`{
		"name": "name test create",
		"author": "author test create",
		"release_year": 1982,
		"isbn": "isbn test create"
	}`)
	req, _ := http.NewRequest("POST", "/books/", bytes.NewBuffer(jsonData))
	response := executeRequest(req, a)
	checkResponseCode(t, http.StatusCreated, response.Code)
	book := model.Book{}
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&book); err != nil {
		t.Errorf("Empty response")
	}
	req, _ = http.NewRequest("GET", "/books/"+book.ID.String(), nil)
	response = executeRequest(req, a)
	checkResponseCode(t, http.StatusOK, response.Code)
	body := response.Body.String()
	sampleDbData := `{
		"id": "` + book.ID.String() + `",
		"name": "name test create",
		"author": "author test create",
		"release_year": 1982,
		"isbn": "isbn test create"
	}`
	require.JSONEq(t, sampleDbData, body)
}

func executeRequest(req *http.Request, a *app.App) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

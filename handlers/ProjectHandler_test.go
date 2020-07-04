package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"task-manager/adapters"
	"task-manager/models"
	"testing"
)

var srv *httptest.Server

func TestMain(m *testing.M) {
	connStr := "user=postgres password=manager dbname=test_task_manager host=localhost sslmode=disable"

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	postgresAdapter := adapters.PostgresAdapter{Connection: conn}
	projectHandler := NewProjectHandler(postgresAdapter)
	columnHandler := NewColumnHandler(postgresAdapter)
	taskHandler := NewTaskHandler(postgresAdapter)
	commentHandler := NewCommentHandler(postgresAdapter)
	r := mux.NewRouter()
	s := r.PathPrefix("/v1").Subrouter()
	s.HandleFunc("/projects", projectHandler.GetList).Methods(http.MethodGet)
	s.HandleFunc("/project/{id}", projectHandler.Get).Methods(http.MethodGet)
	s.HandleFunc("/project", projectHandler.Create).Methods(http.MethodPost)
	s.HandleFunc("/project", projectHandler.Update).Methods(http.MethodPut)
	s.HandleFunc("/project/{id}", projectHandler.Delete).Methods(http.MethodDelete)

	s.HandleFunc("/columns/{projectId}", columnHandler.GetByProjectId).Methods(http.MethodGet)
	s.HandleFunc("/column/{id}", columnHandler.Get).Methods(http.MethodGet)
	s.HandleFunc("/column", columnHandler.Create).Methods(http.MethodPost)
	s.HandleFunc("/column", columnHandler.Update).Methods(http.MethodPut)
	s.HandleFunc("/column/{id}", columnHandler.Delete).Methods(http.MethodDelete)

	s.HandleFunc("/tasks/{columnId}", taskHandler.GetByColumnId).Methods(http.MethodGet)
	s.HandleFunc("/task/{id}", taskHandler.Get).Methods(http.MethodGet)
	s.HandleFunc("/task", taskHandler.Create).Methods(http.MethodPost)
	s.HandleFunc("/task", taskHandler.Update).Methods(http.MethodPut)
	s.HandleFunc("/task/{id}", taskHandler.Delete).Methods(http.MethodDelete)

	s.HandleFunc("/comments/{taskId}", commentHandler.GetByTaskId).Methods(http.MethodGet)
	s.HandleFunc("/comment/{id}", commentHandler.Get).Methods(http.MethodGet)
	s.HandleFunc("/comment", commentHandler.Create).Methods(http.MethodPost)
	s.HandleFunc("/comment", commentHandler.Update).Methods(http.MethodPut)
	s.HandleFunc("/comment/{id}", commentHandler.Delete).Methods(http.MethodDelete)

	srv = httptest.NewServer(r)
	defer srv.Close()

	m.Run()
}

func TestProjectHandlerGetList(t *testing.T) {
	res, err := http.Get(fmt.Sprintf("%s/v1/projects", srv.URL))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var projects []*models.ProjectModel
	if err = json.Unmarshal(resBytes, &projects); err != nil {
		t.Fatal(err)
	}

	assert.Equal(
		t,
		&models.ProjectModel{Id: 1, Name: "test", Description: "test"},
		projects[0],
	)
}

func TestProjectHandlerGet(t *testing.T) {
	res, err := http.Get(fmt.Sprintf("%s/v1/project/1", srv.URL))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var project models.ProjectModel
	if err = json.Unmarshal(resBytes, &project); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "test", project.Name)
}

func TestProjectHandlerCreate(t *testing.T) {
	jsonStr := []byte(`{
 		"name": "string",
 		"description": "string"
	}`)

	res, err := http.Post(
		fmt.Sprintf("%s/v1/project", srv.URL),
		"application/json",
		bytes.NewBuffer(jsonStr),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var project models.ProjectModel
	if err = json.Unmarshal(resBytes, &project); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "string", project.Name)
}

func TestProjectHandlerUpdate(t *testing.T) {
	jsonStr := []byte(`{
		"id": 1,
 		"name": "update",
 		"description": "string"
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/v1/project", srv.URL), bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var project models.ProjectModel
	if err = json.Unmarshal(resBytes, &project); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "update", project.Name)
}

func TestProjectHandlerDelete(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/v1/project/2", srv.URL), nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var success string
	if err = json.Unmarshal(resBytes, &success); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "success", success)
}

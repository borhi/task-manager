package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"task-manager/models"
	"testing"
)

func TestTaskHandlerGetByColumnId(t *testing.T) {
	res, err := http.Get(fmt.Sprintf("%s/v1/tasks/1", srv.URL))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var tasks []*models.TaskModel
	if err = json.Unmarshal(resBytes, &tasks); err != nil {
		t.Fatal(err)
	}

	assert.Equal(
		t,
		&models.TaskModel{Id: 1, Name: "test", Description: "test", Position: 1, ColumnId: 1},
		tasks[0],
	)
}

func TestTaskHandlerGet(t *testing.T) {
	res, err := http.Get(fmt.Sprintf("%s/v1/task/1", srv.URL))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var task models.TaskModel
	if err = json.Unmarshal(resBytes, &task); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "test", task.Name)
}

func TestTaskHandlerCreate(t *testing.T) {
	jsonStr := []byte(`{
		"name": "string",
		"description": "string",
		"position": 1,
		"column_id": 1
	}`)

	res, err := http.Post(
		fmt.Sprintf("%s/v1/task", srv.URL),
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

	var task models.TaskModel
	if err = json.Unmarshal(resBytes, &task); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "string", task.Name)
}

func TestTaskHandlerUpdate(t *testing.T) {
	jsonStr := []byte(`{
		"id": 2,
		"name": "update",
		"description": "string",
		"position": 1,
		"column_id": 1
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/v1/task", srv.URL), bytes.NewBuffer(jsonStr))
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

	var task models.TaskModel
	if err = json.Unmarshal(resBytes, &task); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "update", task.Name)
}

func TestTaskHandlerDelete(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/v1/task/2", srv.URL), nil)
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

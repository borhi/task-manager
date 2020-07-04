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

func TestColumnHandlerGetByProjectId(t *testing.T) {
	res, err := http.Get(fmt.Sprintf("%s/v1/columns/1", srv.URL))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var columns []*models.ColumnModel
	if err = json.Unmarshal(resBytes, &columns); err != nil {
		t.Fatal(err)
	}

	assert.Equal(
		t,
		&models.ColumnModel{Id: 1, Name: "test", Position: 1, ProjectId: 1},
		columns[0],
	)
}

func TestColumnHandlerGet(t *testing.T) {
	res, err := http.Get(fmt.Sprintf("%s/v1/column/1", srv.URL))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var column models.ColumnModel
	if err = json.Unmarshal(resBytes, &column); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "test", column.Name)
}

func TestColumnHandlerCreate(t *testing.T) {
	jsonStr := []byte(`{
		"name": "string",
		"position": 1,
		"project_id": 1
	}`)

	res, err := http.Post(
		fmt.Sprintf("%s/v1/column", srv.URL),
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

	var column models.ColumnModel
	if err = json.Unmarshal(resBytes, &column); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "string", column.Name)
}

func TestColumnHandlerUpdate(t *testing.T) {
	jsonStr := []byte(`{
		"id": 2,
		"name": "update",
		"position": 1,
		"project_id": 1
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/v1/column", srv.URL), bytes.NewBuffer(jsonStr))
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

	var column models.ColumnModel
	if err = json.Unmarshal(resBytes, &column); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "update", column.Name)
}

func TestColumnHandlerDelete(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/v1/column/2", srv.URL), nil)
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

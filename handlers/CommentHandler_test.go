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

func TestCommentHandlerGetByTaskId(t *testing.T) {
	res, err := http.Get(fmt.Sprintf("%s/v1/comments/1", srv.URL))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var comments []*models.CommentModel
	if err = json.Unmarshal(resBytes, &comments); err != nil {
		t.Fatal(err)
	}

	assert.Equal(
		t,
		"test",
		comments[0].Text,
	)
}

func TestCommentHandlerGet(t *testing.T) {
	res, err := http.Get(fmt.Sprintf("%s/v1/comment/1", srv.URL))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var comment models.CommentModel
	if err = json.Unmarshal(resBytes, &comment); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "test", comment.Text)
}

func TestCommentHandlerCreate(t *testing.T) {
	jsonStr := []byte(`{
		"name": "string",
		"text": "string",
		"task_id": 1
	}`)

	res, err := http.Post(
		fmt.Sprintf("%s/v1/comment", srv.URL),
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

	var comment models.CommentModel
	if err = json.Unmarshal(resBytes, &comment); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "string", comment.Text)
}

func TestCommentHandlerUpdate(t *testing.T) {
	jsonStr := []byte(`{
		"id": 2,
		"text": "update",
		"task_id": 1
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/v1/comment", srv.URL), bytes.NewBuffer(jsonStr))
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

	var coment models.CommentModel
	if err = json.Unmarshal(resBytes, &coment); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "update", coment.Text)
}

func TestCommentHandlerDelete(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/v1/comment/2", srv.URL), nil)
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

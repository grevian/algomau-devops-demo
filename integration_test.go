//go:integration

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func Test_webServer(t *testing.T) {
	server := runWebserver()
	defer server.Shutdown(context.TODO())
	serverPath := fmt.Sprintf("http://%s", server.Addr)

	expectedResponse := `Hello, "!gnitset/"`
	response, err := http.DefaultClient.Get(serverPath + "/testing!")
	if err != nil {
		t.Error(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Error(response)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	if string(bodyBytes) != expectedResponse {
		t.Fail()
	}
}

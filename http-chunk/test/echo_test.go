package test

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestEcho(t *testing.T)  {
	resp, err := http.Get("http://127.0.0.1:8080/echo")
	if err!=nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	t.Log(resp.TransferEncoding)
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadString('\n')
		if len(line) > 0 {
			fmt.Print(line)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
	}
}

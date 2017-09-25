// 代码来自《GO语言实战》
// 这个示例程序展示如何测试内部服务端点的执行效果
package handlers_test

import (
	"encoding/json"
	"github.com/fengchunjian/goexamples/testing/endpoint/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

func init() {
	handlers.Routes()
}

// TestSendJSON测试/sendjson内部服务端点
func TestSendJSON(t *testing.T) {
	t.Log("Given the need to test the SendJSON endpoint.")
	{
		req, err := http.NewRequest("GET", "/sendjson", nil)
		if err != nil {
			t.Fatal("\tShould be able to create a request.", ballotX, err)
		}
		t.Log("\tShould be able to create a request.", checkMark)

		rw := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(rw, req)

		if rw.Code != 200 {
			t.Fatal("\tShould receive \"200\"", ballotX, rw.Code)
		}
		t.Log("\tShould receive \"200\"", checkMark)

		u := struct {
			Name  string
			Email string
		}{}

		if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
			t.Fatal("\tShould decode the response.", ballotX)
		}
		t.Log("\tShould decode the response.", checkMark)

		t.Logf("Name: %s\nEmail:%s\n", u.Name, u.Email)
	}
}

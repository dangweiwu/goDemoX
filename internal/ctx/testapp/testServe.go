package testapp

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
)

type TestServer struct {
	Request  *http.Request
	Response *httptest.ResponseRecorder
	Engine   *gin.Engine
}

/*
*
test http server
*/
func NewTestServer() *TestServer {
	a := &TestServer{}
	gin.SetMode(gin.ReleaseMode)
	a.Engine = gin.New()

	return a
}

// 注册路由
func (this *TestServer) RegRoute(f func(engine *gin.Engine)) {
	this.Engine = gin.New()
	f(this.Engine)
}

func (this *TestServer) SetToken(token string) *TestServer {
	this.Request.Header.Add("Authorization", token)
	return this
}

func (this *TestServer) SetBaseAuth(user, password string) *TestServer {
	base := user + ":" + password
	t := "Basic " + base64.StdEncoding.EncodeToString([]byte(base))
	this.Request.Header.Add("Authorization", t)
	return this
}

func (this *TestServer) Fetch(method string, target string, body io.Reader) *TestServer {
	this.Request = httptest.NewRequest(method, target, body)
	return this
}

func (this *TestServer) Get(target string) *TestServer {
	return this.Fetch(http.MethodGet, target, nil)
}

func (this *TestServer) Post(target string, obj interface{}) *TestServer {
	r, _ := json.Marshal(obj)

	return this.Fetch("POST", target, bytes.NewReader(r))
}

func (this *TestServer) Put(target string, obj interface{}) *TestServer {
	r, _ := json.Marshal(obj)
	return this.Fetch("PUT", target, bytes.NewReader(r))
}

func (this *TestServer) Delete(target string) *TestServer {

	return this.Fetch("DELETE", target, nil)
}

func (this *TestServer) Do() *TestServer {
	this.Response = httptest.NewRecorder()
	this.Engine.ServeHTTP(this.Response, this.Request)
	return this
}

func (this *TestServer) GetBody() string {
	return this.Response.Body.String()
}

func (this *TestServer) ResponseObj(obj interface{}) error {
	return json.Unmarshal([]byte(this.GetBody()), obj)
}
func (this *TestServer) GetCode() int {
	return this.Response.Code
}

package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/quangdangfit/gocommon/logger"
	"go.jumia.org/customers/app/migration"
	"go.uber.org/dig"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"go.jumia.org/customers/app"
)

var (
	engine    *gin.Engine
	container *dig.Container
)

type Results struct {
	Code string `json:"code"`
	Data []struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Phone       string `json:"phone"`
		Country     string `json:"country"`
		CountryCode string `json:"country_code"`
		State       bool   `json:"state"`
	} `json:"data"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func TestMain(m *testing.M) {
	logger.Initialize("testing")
	container = app.BuildContainer()
	engine = app.InitGinEngine(container)

	setup()
	code := m.Run()
	teardown()

	os.Exit(code)
}

func setup() {
	logger.Info("============> Setup for testing")
	removeUserData()

	migrate()
}

func teardown() {
	logger.Info("============> Teardown")
	removeUserData()
}

func migrate() {
	migration.Migrate(container)
}

func removeUserData() {
	//container.Invoke(func(
	//	customerRepo repositories.IUserRepository,
	//) error {
	//	for _, u := range customers {
	//		err := customerRepo.Delete(u.ID)
	//		if err != nil {
	//			return err
	//		}
	//	}
	//	return nil
	//})
}

func TestSuccessRequest(t *testing.T) {

	w := httptest.NewRecorder()
	req := newGetRequest("/api/v1/customers?state=true", nil)
	engine.ServeHTTP(w, req)

	var resp map[string]interface{}
	parseReader(w.Body, &resp)
	if resp["status"].(float64) != http.StatusOK {
		t.Fail()
	}

}

func TestBadRequest(t *testing.T) {

	w := httptest.NewRecorder()
	req := newGetRequest("/api/v1/customers?state=6", nil)
	engine.ServeHTTP(w, req)

	var resp Results
	json.NewDecoder(w.Body).Decode(&resp)
	if resp.Status != http.StatusBadRequest {
		t.Fail()
	}

}

func TestMaxLimitRequest(t *testing.T) {

	w := httptest.NewRecorder()
	req := newGetRequest("/api/v1/customers?page_size=10", nil)
	engine.ServeHTTP(w, req)

	var resp Results
	json.NewDecoder(w.Body).Decode(&resp)
	if len(resp.Data) != 10 {
		t.Fail()
	}

}

func TestInvalidMaxLimitParamRequest(t *testing.T) {

	w := httptest.NewRecorder()
	req := newGetRequest("/api/v1/customers?page_size=a", nil)
	engine.ServeHTTP(w, req)

	var resp Results
	json.NewDecoder(w.Body).Decode(&resp)
	if resp.Status != http.StatusBadRequest {
		t.Fail()
	}

}

func TestMaximumMaxLimitParamRequest(t *testing.T) {

	w := httptest.NewRecorder()
	req := newGetRequest("/api/v1/customers?page_size=2000", nil)
	engine.ServeHTTP(w, req)

	var resp Results
	json.NewDecoder(w.Body).Decode(&resp)
	if len(resp.Data) != 30 {
		t.Fail()
	}

}

func TestOneCountryOnlyRequest(t *testing.T) {

	w := httptest.NewRecorder()
	req := newGetRequest("/api/v1/customers?country=Cameroon", nil)
	engine.ServeHTTP(w, req)

	var resp Results
	json.NewDecoder(w.Body).Decode(&resp)
	for _, data := range resp.Data {
		if strings.ToLower(data.Country) != "cameroon" {
			t.Fail()
			return
		}
	}

}

func toReader(v interface{}) io.Reader {
	buf := new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(v)
	return buf
}

func parseReader(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func newGetRequest(formatRouter string, v interface{}, args ...interface{}) *http.Request {
	req, _ := http.NewRequest("GET", fmt.Sprintf(formatRouter, args...), toReader(v))
	return req
}

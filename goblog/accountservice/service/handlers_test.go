package service

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/callistanenterprise/goblog/accountservice/dbclient"
	"github.com/callistanenterprise/goblog/accountservice/model"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetAccountWrongPath(t *testing.T) {
	Convey("Given a HTTP request for /invalid/123", t, func() {
		req := httptest.NewRequest("GET", "/invalid/123", nil)
		resp := httptest.NewRecorder()
		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)
			Convey("Then the response should be a 404", func() {
				So(resp.Code, ShouldEqual, 404)
			})
		})
	})
}

func TestGetAccount(t *testing.T) {
	mockRepo := &dbclient.MockBoltClient{}
	mockRepo.On("QueryAccount", "123").Return(model.Account{Id: "123", Name: "Person_123"}, nil)
	mockRepo.On("QueryAccount", "456").Return(model.Account{}, fmt.Errorf("Some error"))
	DBClient = mockRepo

	Convey("Given a HTTP request for /accounts/123", t, func() {
		req := httptest.NewRequest("GET", "/accounts/123", nil)
		resp := httptest.NewRecorder()
		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)
			Convey("Then the response should be a 200", func() {
				So(resp.Code, ShouldEqual, 200)
				account := model.Account{}
				json.Unmarshal(resp.Body.Bytes(), &account)
				So(account.Id, ShouldEqual, "123")
				So(account.Name, ShouldEqual, "Person_123")
			})
		})
	})
	/*Convey("Given a HTTP request for /accounts/456", t, func() {
		req := httptest.NewRequest("GET", "/accounts/456", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 404", func() {
				So(resp.Code, ShouldEqual, 404)
			})
		})
	})*/
}

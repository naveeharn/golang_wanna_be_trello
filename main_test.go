package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

var (
	r = gin.Default()
)

// func SetUpRouter() *gin.Engine {
// 	router := gin.Default()
// 	// router.Run(":4011")
// 	return router
// }

func SetUpRequest(relativePath, jsonBody string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, relativePath, bytes.NewBuffer([]byte(jsonBody)))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return req
}

func TestRegisterMethod_BadRequest(t *testing.T) {
	r.POST("/api/auth/register", authController.Register)
	req := SetUpRequest("/api/auth/register", `{"email": "a@a.com","password": "aaa"}`)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// responseData, _ := ioutil.ReadAll(w.Body)
	// log.Println("\n", string(responseData))
	assert.Equal(t, w.Code, http.StatusBadRequest)

}

func TestLoginMethod_StatusOK(t *testing.T) {
	r.POST("/api/auth/login", authController.Login)
	req := SetUpRequest("/api/auth/login", `{"email": "a@a.com","password": "aaa"}`)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// responseData, _ := ioutil.ReadAll(w.Body)
	// log.Println("\n", string(responseData))
	assert.Equal(t, w.Code, http.StatusOK)
}

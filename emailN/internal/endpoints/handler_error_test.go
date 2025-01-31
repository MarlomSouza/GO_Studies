package endpoints

import (
	internalerrors "emailn/internal/internal-errors"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HandlerError_when_endpoint_returns_internal_error(t *testing.T) {
	assert := assert.New(t)
	endpointWithError := func(w http.ResponseWriter, r *http.Request) (EndpointStruct, error) {
		return EndpointStruct{}, internalerrors.ErrInternal
	}
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	sut := HandlerError(endpointWithError)

	sut.ServeHTTP(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internalerrors.ErrInternal.Error())
}

func Test_HandlerError_when_endpoint_returns_domain_error(t *testing.T) {
	assert := assert.New(t)
	endpointsWithDomainError := func(w http.ResponseWriter, r *http.Request) (EndpointStruct, error) {
		return EndpointStruct{}, errors.New("domain error")
	}
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	sut := HandlerError(endpointsWithDomainError)

	sut.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
	assert.Contains(res.Body.String(), "domain error")
}

func Test_HandlerError_when_endpoint_returns_ok(t *testing.T) {
	assert := assert.New(t)
	type bodyForTest struct {
		Id int
	}
	expectedObj := bodyForTest{Id: 2}
	endpointsWithDomainError := func(w http.ResponseWriter, r *http.Request) (EndpointStruct, error) {
		return EndpointStruct{Obj: expectedObj, Status: http.StatusOK}, nil
	}
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	sut := HandlerError(endpointsWithDomainError)

	sut.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	result := bodyForTest{}
	json.Unmarshal(res.Body.Bytes(), &result)
	assert.Equal(expectedObj, result)
}

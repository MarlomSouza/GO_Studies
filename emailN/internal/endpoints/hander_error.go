package endpoints

import (
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type EndpointStruct struct {
	Obj    interface{}
	Status int
}

type EndpointFunc func(w http.ResponseWriter, r *http.Request) (EndpointStruct, error)

func HandlerError(endpointFunc EndpointFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := endpointFunc(w, r)
		if err != nil {
			if errors.Is(err, internalerrors.ErrInternal) {
				render.Status(r, http.StatusInternalServerError)
			} else if errors.Is(err, internalerrors.ErrNotFound) {
				render.Status(r, http.StatusNotFound)
			} else {
				render.Status(r, http.StatusBadRequest)
			}

			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		render.Status(r, response.Status)

		if response.Obj != nil {
			render.JSON(w, r, response.Obj)
		}
	})
}

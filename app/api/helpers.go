package api

import (
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type request struct{}

func (s *request) Bind(_ *http.Request) error {
	return nil
}

func bindRequest(r *http.Request, data render.Binder) error {
	if err := render.Bind(r, data); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return err
	}

	return nil
}

type Response struct {
	HTTPStatusCode int         `json:"-"`
	Payload        interface{} `json:"payload"`
}

func (r *Response) Render(_ http.ResponseWriter, req *http.Request) error {
	render.Status(req, r.HTTPStatusCode)
	return nil
}

func response(w http.ResponseWriter, req *http.Request, code int, payload interface{}) {
	r := &Response{
		HTTPStatusCode: code,
		Payload:        payload,
	}

	if err := render.Render(w, req, r); err != nil {
		// todo переписать логер
		log.Println(err)
	}
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(_ http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func responseError(w http.ResponseWriter, req *http.Request, httpStatusCode int, err error, details string) {
	r := &ErrResponse{
		Err:            err,
		HTTPStatusCode: httpStatusCode,
		StatusText:     details,
		ErrorText:      err.Error(),
	}

	if err := render.Render(w, req, r); err != nil {
		// todo переписать логер
		log.Println(err)
	}
}

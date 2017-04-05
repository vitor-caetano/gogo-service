package service

import (
	"net/http"
	"github.com/unrolled/render"
)

func createMatchHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Location", fakeMatchLocationResult1)
		formatter.JSON(w,
			http.StatusCreated,
				struct{ Test string }{"This is a test"})
	}
}

const (
	fakeMatchLocationResult1 = "/matches/5a003b78-409e-4452-b456-a6f0dcee05bd"
)
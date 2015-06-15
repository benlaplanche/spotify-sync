package main_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/benlaplanche/spotify-sync"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("API", func() {

	Describe("/ root path tests", func() {
		makeRequest := func() *httptest.ResponseRecorder {
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", "/", nil)

			RootPathHandler(recorder, request)

			return recorder
		}
		It("is successful", func() {
			response := makeRequest()

			Expect(response.Body.String()).To(ContainSubstring("hello world"))
			Expect(response.Code).To(Equal(200))
		})

	})
})

package services_test

import (
	"exchange/modules"
	"exchange/routes"
	"exchange/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExchangeRate(t *testing.T) {
	if err := modules.LoadCurrencyConfig("../rate.json"); err != nil {
		t.Fatal(err)
	}

	r := routes.Router()
	r.GET("/exchange-rates", services.ExchangeRate)

	testCases := []struct {
		source   string
		target   string
		amount   string
		expected int
	}{
		{"USD", "TWD", `$1,000`, http.StatusOK},

		// Empty parameters
		{"", "TWD", `$1,000`, http.StatusBadRequest},

		// Non-integer amount
		{"USD", "TWD", "abc", http.StatusBadRequest},

		// Unsupported currency
		{"USD", "EUR", `$1,000`, http.StatusBadRequest},
	}

	for _, tc := range testCases {
		req, err := http.NewRequest("GET", "/exchange-rates", nil)
		if err != nil {
			t.Fatal(err)
		}

		// Set query parameters
		query := req.URL.Query()
		query.Add("source", tc.source)
		query.Add("target", tc.target)
		query.Add("amount", tc.amount)
		req.URL.RawQuery = query.Encode()

		// Create a response recorder
		w := httptest.NewRecorder()

		// Send the request
		r.ServeHTTP(w, req)

		if w.Code != tc.expected {
			t.Errorf("Expected status code %d, but got %d", tc.expected, w.Code)
		}
	}
}

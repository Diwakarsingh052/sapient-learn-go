package main

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDoubleHandler(t *testing.T) {
	tt := []struct {
		name           string
		queryParam     string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "OK",
			queryParam:     "10",
			expectedStatus: http.StatusOK,
			expectedBody:   "20",
		},
		{
			name:           "Fail_MissingValue",
			queryParam:     "", // Missing `v` parameter
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "missing value",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			//ResponseRecorder is an implementation of [http.ResponseWriter
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/double?v="+tc.queryParam, nil)
			doubleHandler(w, r)

			require.Equal(t, tc.expectedStatus, w.Code)
			require.Equal(t, tc.expectedBody, strings.TrimSpace(w.Body.String()))
		})
	}
}

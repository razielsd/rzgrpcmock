package mockserver

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestServer_handlerHealthProbe(t *testing.T) {
	w, r := createGetReqAndWriter()
	api := createServer(t)
	api.handlerHealthProbe(w, r)

	require.Equal(t, http.StatusOK, w.Code)
	exp := SuccessResponse{Result: NewSuccessOK()}
	require.JSONEq(t, exp.JSON(), w.Body.String())
}
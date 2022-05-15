package mockserver

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServer_handlerHealthProbe(t *testing.T) {
	w, r := createGetReqAndWriter()
	api := createServer(t)
	api.handlerHealthProbe(w, r)

	require.Equal(t, http.StatusOK, w.Code)
	exp := SuccessResponse{Result: NewSuccessOK()}
	require.JSONEq(t, exp.JSON(), w.Body.String())
}

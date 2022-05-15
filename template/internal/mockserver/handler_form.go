package mockserver

import (
	"net/http"
	"strings"
)

func (s *Server) handlerForm(w http.ResponseWriter, r *http.Request) {
	form := `<html lang="en">
<head>
    <meta charset="utf-8" />
    <style>
        div.form-row {
            width: 80%;
            margin: 10pt;
        }
        div.form-row input[type=text] {
            width: 80%;
        }
        div.form-row input[type=submit] {
            width: 10%;
            font-weight: bold;
        }
        div.form-row textarea {
            width: 80%;
            height: 80pt;
        }
    </style>
</head>
<body>
<h1>Mock Debug Form</h1>
<form method="post" action="http://{{.Host}}/api/mock/add">
	<input type="hidden" name="ref" value="form">
    <div class="form-row">
    <label for="method">Method</label><br>
    <input name="method"  id="method" type="text" value="">
    </div>
    <div class="form-row">
        <label for="method">Request</label><br>
        <textarea name="request" id="request" type="text" value=""></textarea>
    </div>
    <div class="form-row">
        <label for="response">Response</label><br>
        <textarea name="response" id="response" type="text" value=""></textarea>
    </div>

    <div class="form-row">
        <input type="submit" value="Send">
    </div>
</form>
</body>
</html>
`
	addr := s.Addr
	if strings.HasPrefix(addr, ":") {
		addr = "0.0.0.0" + addr
	}
	form = strings.Replace(form, "{{.Host}}", addr, 1)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(form))
}

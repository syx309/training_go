package err

import "net/http"

const (
	JwtKey = "JWTtestingdoangbisagasih"
)

func ErrorForbidden(writer http.ResponseWriter) {
	http.Error(writer, http.StatusText(http.StatusForbidden), http.StatusForbidden)
	writer.Write([]byte(http.StatusText(http.StatusForbidden)))
	writer.WriteHeader(403)
}

func ErrorNotFound(writer http.ResponseWriter) {
	http.Error(writer, "Data Not Found..", http.StatusUnauthorized)
	writer.Write([]byte("Data Not Found.."))
	writer.WriteHeader(404)
}

func ErrorInternal(writer http.ResponseWriter) {
	http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	writer.Write([]byte(http.StatusText(http.StatusInternalServerError)))
	writer.WriteHeader(500)
}

func ErrorUnauthorized(writer http.ResponseWriter) {
	http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	writer.Write([]byte(http.StatusText(http.StatusUnauthorized)))
	writer.WriteHeader(401)
}

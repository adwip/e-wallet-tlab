package stacktrace

import (
	"net/http"

	"github.com/adwip/aj-teknik-backend-admin/common-lib/shared/constant"
)

const (
	SUCCESS               stacktraceCode = "ESC2000"
	INTERNAL_SERVER_ERROR stacktraceCode = "ESC5001"
	BAD_PROCESSING        stacktraceCode = "ESC5002"
	SERVICE_UNAVAILABLE   stacktraceCode = "ESC5003"
	INVALID_INPUT         stacktraceCode = "ESC4001"
	DATA_NOT_FOUND        stacktraceCode = "ESC4002"
	FORBIDDEN             stacktraceCode = "ESC4003"
	UNAUTHENTICATED       stacktraceCode = "ESC4004"
	UNAUTHORIZED          stacktraceCode = "ESC4005"
	BAD_REQUEST           stacktraceCode = "ESC4006"
)

// stdline Color code
const (
	INFO    = 36
	WARNING = 33
	FATAL   = 31
)

const (
	MESSAGE_INTERNAL_SERVER_ERROR string = "Internal server error"
	MESSAGE_BAD_PROCESSING        string = "Bad processing"
	MESSAGE_SERVICE_UNAVAILABLE   string = "Service unavailable"
	MESSAGE_INVALID_INPUT         string = "Invalid input"
	MESSAGE_DATA_NOT_FOUND        string = "Data not found"
	MESSAGE_FORBIDDEN             string = "Forbidden"
	MESSAGE_UNAUTHENTICATED       string = "Unauthenticated"
	MESSAGE_UNAUTHORIZED          string = "Unauthorized"
	MESSAGE_BAD_REQUEST           string = "Bad requests"
)

var server_error = map[stacktraceCode]string{
	INTERNAL_SERVER_ERROR: MESSAGE_INTERNAL_SERVER_ERROR,
	BAD_PROCESSING:        MESSAGE_BAD_PROCESSING,
	SERVICE_UNAVAILABLE:   MESSAGE_SERVICE_UNAVAILABLE,
}

var client_error = map[stacktraceCode]string{
	INVALID_INPUT:   MESSAGE_INVALID_INPUT,
	DATA_NOT_FOUND:  MESSAGE_DATA_NOT_FOUND,
	FORBIDDEN:       MESSAGE_FORBIDDEN,
	UNAUTHENTICATED: MESSAGE_UNAUTHENTICATED,
	UNAUTHORIZED:    MESSAGE_UNAUTHORIZED,
	BAD_REQUEST:     MESSAGE_BAD_REQUEST,
}

var stacktraceToHttpCode = map[stacktraceCode]int{
	SUCCESS:               http.StatusOK,                  // 200: Berhasil
	INTERNAL_SERVER_ERROR: http.StatusInternalServerError, // 500: Kesalahan server
	BAD_PROCESSING:        http.StatusUnprocessableEntity, // 422: Kesalahan pemrosesan
	SERVICE_UNAVAILABLE:   http.StatusServiceUnavailable,  // 503: Layanan tidak tersedia
	INVALID_INPUT:         http.StatusBadRequest,          // 400: Input tidak valid
	DATA_NOT_FOUND:        http.StatusNotFound,            // 404: Data tidak ditemukan
	FORBIDDEN:             http.StatusForbidden,           // 403: Akses dilarang
	UNAUTHENTICATED:       http.StatusUnauthorized,
	UNAUTHORIZED:          http.StatusUnauthorized, // 401: Tidak terautentikasi
	BAD_REQUEST:           http.StatusBadRequest,   // 400: Permintaan tidak valid
}

var stacktraceToLevelCode = map[stacktraceCode]struct {
	code int
	name string
}{
	SUCCESS:               {INFO, constant.INFO},       // Berhasil → INFO
	INTERNAL_SERVER_ERROR: {FATAL, constant.FATAL},     // Kesalahan server → FATAL
	BAD_PROCESSING:        {WARNING, constant.WARNING}, // Kesalahan pemrosesan → WARNING
	SERVICE_UNAVAILABLE:   {FATAL, constant.FATAL},     // Layanan tidak tersedia → FATAL
	INVALID_INPUT:         {WARNING, constant.WARNING}, // Input tidak valid → WARNING
	DATA_NOT_FOUND:        {WARNING, constant.WARNING}, // Data tidak ditemukan → WARNING
	FORBIDDEN:             {WARNING, constant.WARNING}, // Akses dilarang → WARNING
	UNAUTHENTICATED:       {WARNING, constant.WARNING},
	UNAUTHORIZED:          {WARNING, constant.WARNING}, // Tidak terautentikasi → WARNING
	BAD_REQUEST:           {WARNING, constant.WARNING}, // Permintaan tidak valid → WARNING
}

var httpCodeToStacktrace = map[int]stacktraceCode{
	http.StatusOK:                  SUCCESS,
	http.StatusInternalServerError: INTERNAL_SERVER_ERROR,
	http.StatusUnprocessableEntity: BAD_PROCESSING,
	http.StatusServiceUnavailable:  SERVICE_UNAVAILABLE,
	http.StatusBadRequest:          INVALID_INPUT, // atau BAD_REQUEST, tergantung konteks
	http.StatusNotFound:            DATA_NOT_FOUND,
	http.StatusForbidden:           FORBIDDEN,
	http.StatusUnauthorized:        UNAUTHENTICATED,
}

func StacktraceMessageByCode(code string) string {
	out, isset := server_error[stacktraceCode(code)]
	if !isset {
		out = client_error[stacktraceCode(code)]
	}
	return out
}

func HttpStatusCodeByStacktrace(code stacktraceCode) int {
	httpCode := stacktraceToHttpCode[code]
	return httpCode
}

func LevelByStacktrace(code stacktraceCode) (int, string) {
	level := stacktraceToLevelCode[code]

	return level.code, level.name
}

func StacktraceFromHTTPCode(status int) stacktraceCode {
	if val, ok := httpCodeToStacktrace[status]; ok {
		return val
	}
	return INTERNAL_SERVER_ERROR
}

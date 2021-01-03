package logger

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Request struct contains items of request info log
type Request struct {
	Path        string `json:"path,omitempty"`
	Method      string `json:"method,omitempty"`
	Query       string `json:"query,omitempty"`
	ContentType string `json:"content-type,omitempty"`
	Scheme      string `json:"scheme,omitempty"`
	Protocol    string `json:"protocol,omitempty"`
	UserAgent   string `json:"userAgent,omitempty"`
}

// Response struct contains items of request info log
type Response struct {
	StatusCode string `json:"statusCode,omitempty"`
}

// Host struct contains item of host info log
type Host struct {
	Hostname          string `json:"hostname,omitempty"`
	ForwardedHostName string `json:"forwardedHostname,omitempty"`
	IP                string `json:"ip,omitempty"`
}

// HTTP is the struct of the log formatter
type HTTP struct {
	Request  *Request  `json:"request,omitempty"`
	Response *Response `json:"response,omitempty"`
}

func getRequestID(headers http.Header) string {
	if requestID := headers.Get("X-Request-Id"); requestID != "" {
		return requestID
	}

	requestID, err := uuid.NewRandom()

	if err != nil {

	}

	return requestID.String()
}

// LoggingMiddleware func
func LoggingMiddleware(logger *logrus.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			// start := time.Now()

			requestID := getRequestID(r.Header)
			ctx := WithLogger(r.Context(), logrus.NewEntry(logger).WithFields(logrus.Fields{
				"CorrelationId": requestID,
			}))

			writer := logginResponseWriter{writer: rw, statusCode: http.StatusOK}

			// Next Middleware
			next.ServeHTTP(&writer, r.WithContext(ctx))

			Get(ctx).WithFields(logrus.Fields{
				"http": HTTP{
					Request: &Request{
						Path: r.URL.RequestURI(),
					},
				},
			}).Info("Request Completed")
		})
	}
}

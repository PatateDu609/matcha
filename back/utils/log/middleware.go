package log

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

func NewRouterLogger(logger *logrus.Logger) func(http.Handler) http.Handler {
	return middleware.RequestLogger(
		&LogrusLogger{
			Logger: logger,
		},
	)
}

type LogrusLogger struct {
	Logger *logrus.Logger
}

type LogrusEntry struct {
	Logger logrus.FieldLogger
}

func (logger *LogrusLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &LogrusEntry{
		Logger: logrus.NewEntry(logger.Logger),
	}
	fields := logrus.Fields{}

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		fields["req_id"] = reqID
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	fields["http_scheme"] = scheme
	fields["http_proto"] = r.Proto
	fields["http_method"] = r.Method

	fields["remote_addr"] = r.RemoteAddr
	fields["user_agent"] = r.UserAgent()

	fields["uri"] = fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)

	entry.Logger = entry.Logger.WithFields(fields)
	entry.Logger.Infoln("request started")

	return entry
}

func (l *LogrusEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	l.Logger = l.Logger.WithFields(logrus.Fields{
		"resp_status": status, "resp_bytes_length": bytes,
		"resp_elapsed_ms": float64(elapsed.Nanoseconds()) / 1000000.0,
	})

	l.Logger.Infoln("request complete")
}

func (l *LogrusEntry) Panic(v interface{}, stack []byte) {
	l.Logger = l.Logger.WithFields(logrus.Fields{
		"panic": fmt.Sprintf("%+v", v),
	})
}

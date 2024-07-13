package bootstrap

import "github.com/sirupsen/logrus"

// SetJSONFormatter setup log json formatter
func SetJSONFormatter() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FieldMap:        logrus.FieldMap{logrus.FieldKeyMsg: "message"},
	})
	logrus.SetReportCaller(true)

}

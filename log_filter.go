package filters

import "net/http"
import "github.com/chakrit/go-bunyan"
import z "github.com/zaiuz/zaiuz"

const LogFilterPrefix = "z.log"

type logFilterData struct {
	logger bunyan.Log
}

func LogFilter(name string, parent bunyan.Log) z.Filter {
	if parent == nil {
		parent = bunyan.NewStdLogger(name, bunyan.StdoutSink())
	}
	if name != "" {
		parent = parent.Child()
	}

	data := &logFilterData{parent}

	return func(action z.Action) z.Action {
		return func(c *z.Context) z.Result {
			var sublogger bunyan.Log
			rid := GetRequestId(c)
			if rid != "" {
				sublogger = parent.Record("request_id", rid).Child()
			} else {
				sublogger = parent.Child()
			}

			sublogger.Record("headers", headerDigest(c.Request.Header)).
				Infof("%s %s", c.Request.Method, c.Request.URL.Path)
			c.Set(LogFilterPrefix, data)

			result := action(c)

			duration := GetDuration(c)
			if duration != 0 {
				sublogger.Infof("finish %s", duration.String())
			} else {
				sublogger.Infof("finish")
			}

			return result
		}
	}
}

func GetLogger(c *z.Context) bunyan.Log {
	logger_, ok := c.GetOk(LogFilterPrefix)
	if !ok {
		return nil
	}

	logger, ok := logger_.(bunyan.Log)
	if !ok {
		return nil
	}

	return logger
}

func headerDigest(headers http.Header) map[string]string {
	interests := []string{"Accept", "User-Agent", "Referer"}

	result := map[string]string{}
	for _, header := range interests {
		values, ok := headers[header]
		if ok && len(values) > 0 {
			result[header] = values[0]
		}
	}

	return result
}

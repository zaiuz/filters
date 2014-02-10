package filters

import "code.google.com/p/go-uuid/uuid"
import z "github.com/zaiuz/zaiuz"

const RequestIdPrefix = "z.requestid"

func RequestIdFilter() z.Filter {
	return func(action z.Action) z.Action {
		return func(c *z.Context) z.Result {
			c.Set(RequestIdPrefix, uuid.New())
			return action(c)
		}
	}
}

func GetRequestId(c *z.Context) string {
	str_, ok := c.GetOk(RequestIdPrefix)
	if !ok {
		return ""
	}

	str, ok := str_.(string)
	if !ok {
		return ""
	}

	return str
}

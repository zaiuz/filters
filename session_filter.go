package filters

import "github.com/gorilla/sessions"
import z "github.com/zaiuz/zaiuz"

const SessionFilterPrefix = "z.session"

type sessionFilterData struct {
	store sessions.Store
	name  string
}

func SessionFilter(name, secret string) z.Filter {
	store := sessions.NewCookieStore([]byte(secret))
	data := &sessionFilterData{store, name}

	return func(action z.Action) z.Action {
		return func(c *z.Context) z.Result {
			c.Set(SessionFilterPrefix+".data", data)
			result := action(c)

			dirty, ok := c.GetOk(SessionFilterPrefix+".dirty")
			if ok && dirty.(bool) {
				session := GetSession(c)
				e := session.Save(c.Request, c.ResponseWriter)
				noError(e)
			}

			return result
		}
	}
}

func GetSession(c *z.Context) *sessions.Session {
	data_, ok := c.GetOk(SessionFilterPrefix+".data")
	if !ok {
		return nil
	}

	data := data_.(*sessionFilterData)
	session, e := data.store.Get(c.Request, data.name)
	noError(e)

	c.Set(SessionFilterPrefix+".dirty", true)
	return session
}


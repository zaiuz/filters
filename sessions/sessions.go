package sessions

import "github.com/gorilla/sessions"
import "errors"
import "../../context"
import m ".."

const ContextKey = "zaius.SessionModule"

type SessionModule struct {
	store sessions.Store
	name  string
}

var _ m.Module = new(SessionModule)

func CookieSession(name, secret string) *SessionModule {
	store := sessions.NewCookieStore([]byte(secret))
	return &SessionModule{store, name}
}

func (s *SessionModule) Attach(c *context.Context) error {
	// TODO: Access sessions lazily
	session, e := s.store.Get(c.Request, s.name)
	if e != nil {
		return e
	}

	c.Set(ContextKey, session)
	return nil
}

func Get(c *context.Context) (session *sessions.Session, e error) {
	result, exists := c.GetOk(ContextKey)
	if !exists {
		return nil, errors.New("Get() requires session module in current route.")
	}

	// TODO: Soft type assert
	return result.(*sessions.Session), nil
}

func (s *SessionModule) Detach(c *context.Context) error {
	result, exists := c.GetOk(ContextKey)
	if !exists {
		return nil
	}

	session := result.(*sessions.Session)
	return session.Save(c.Request, c.ResponseWriter)
}

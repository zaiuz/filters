package sessions

import "testing"
import "net/http/httptest"
import a "github.com/stretchr/testify/assert"
import "../../context"
import "../../testutil"

const TestSessionName = "zaius.modules.sessions.Cookie"
const TestSessionSecret = TestSessionName

func TestCookieSession(t *testing.T) {
	m := newSessionModule()
	a.NotNil(t, m, "ctor function returns nil.")
	a.NotNil(t, m.store, "session store not initialized.")
	a.NotEmpty(t, m.name, "session name should not be nil.")
	a.Equal(t, m.name, TestSessionName, "session name is wrong.")
}

func TestGet(t *testing.T) {
	module := newSessionModule()
	context := newTestContext()

	session, e := Get(context)
	a.Nil(t, session, "incorrectly got a session even when not yet attached.")
	a.NotNil(t, e, "get should return errors if called before attaching.")

	module.Attach(context)
	session, e = Get(context)
	a.NoError(t, e, "error getting session.")
	a.NotNil(t, session, "return value nil.")
}

func TestDetach(t *testing.T) {
	const TestValue = "The quick brown fox jumps over the lazy dog."

	m := newSessionModule()
	context := newTestContext()
	m.Attach(context)

	session, e := Get(context)
	a.NoError(t, e, "error when getting session.")

	session.Values["hello"] = TestValue
	m.Detach(context)

	recorder := context.ResponseWriter.(*httptest.ResponseRecorder)
	headers := recorder.Header()
	cookies := headers["Set-Cookie"]
	a.True(t, len(cookies) > 0, "no cookie were set.")

	cookie := cookies[0]
	a.True(t, len(cookie) > 0, "cookie string is empty.")
	a.Contains(t, cookie, TestSessionName, "cookie does not contains session name.")
}

func newSessionModule() *SessionModule {
	return CookieSession(TestSessionName, TestSessionSecret)
}

func newTestContext() *context.Context {
	response, request := testutil.NewTestRequestPair()
	return context.NewContext(response, request)
}

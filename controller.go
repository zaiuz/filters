package filters

import "time"
import "github.com/gorilla/sessions"
import "github.com/chakrit/go-bunyan"
import z "github.com/zaiuz/zaiuz"

type Controller struct{}

func (ctr *Controller) GetLogger(c *z.Context) bunyan.Log { return GetLogger(c) }
func (ctr *Controller) GetRequestId(c *z.Context) string { return GetRequestId(c) }
func (ctr *Controller) GetSession(c *z.Context) *sessions.Session { return GetSession(c) }
func (ctr *Controller) GetDuration(c *z.Context) time.Duration { return GetDuration(c) }


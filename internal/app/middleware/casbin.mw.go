package middleware

import (
	"strconv"

	"bio-assay-ms/internal/app/config"
	"bio-assay-ms/internal/app/contextx"
	"bio-assay-ms/internal/app/ginx"
	"bio-assay-ms/pkg/errors"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// Valid use interface permission
func CasbinMiddleware(enforcer *casbin.SyncedEnforcer, skippers ...SkipperFunc) gin.HandlerFunc {
	cfg := config.C.Casbin
	if !cfg.Enable {
		return EmptyMiddleware()
	}

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		p := c.Request.URL.Path
		m := c.Request.Method
		userID := contextx.FromUserID(c.Request.Context())
		if b, err := enforcer.Enforce(strconv.FormatUint(userID, 10), p, m); err != nil {
			ginx.ResError(c, errors.WithStack(err))
			return
		} else if !b {
			ginx.ResError(c, errors.ErrNoPerm)
			return
		}
		c.Next()
	}
}

package provide

import (
	_ "embed"
	"ginfx-template/pkg/logger"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	stringadapter "github.com/casbin/casbin/v2/persist/string-adapter"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed rbac_model.conf
var rbacModel string

//go:embed policy.csv
var policy string

type Authorizer struct {
	middleware gin.HandlerFunc
	enforcer   *casbin.Enforcer
}

func NewEnforcer() *casbin.Enforcer {
	m, err := model.NewModelFromString(rbacModel)
	if err != nil {
		logger.Error(err)
		return nil
	}

	a := stringadapter.NewAdapter(policy)
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		logger.Error(err)
		return nil
	}
	return e
}

// NewAuthorizer returns the authorizer, uses a Casbin enforcer as input
func NewAuthorizer(e *casbin.Enforcer) *Authorizer {
	a := &Authorizer{
		enforcer: e,
	}
	a.middleware = func(c *gin.Context) {
		if !a.CheckPermission(c) {
			a.RequirePermission(c)
		}
	}
	return a

}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *Authorizer) CheckPermission(c *gin.Context) bool {
	user := GetUserName(c)
	method := c.Request.Method
	path := c.Request.URL.Path

	allowed, err := a.enforcer.Enforce(user, path, method)
	if err != nil {
		logger.Error(err)
		return false
	}

	return allowed
}

// RequirePermission returns the 403 Forbidden to the client
func (a *Authorizer) RequirePermission(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}

func (a *Authorizer) MiddlewareFunc() gin.HandlerFunc {
	return a.middleware
}

package session

import (
	"github.com/alexedwards/scs/v2"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Session struct {
	CookieName     string
	CookieLifetime string
	CookiePersist  string
	CookieSecure   string
	CookieDomain   string
	SessionType    string
}

func (s *Session) InitSession() *scs.SessionManager {
	var persist, secure bool
	minutes, err := strconv.Atoi(s.CookieLifetime)
	if err != nil {
		minutes = 60
	}

	if strings.ToLower(s.CookiePersist) == "true" {
		persist = true
	}

	if strings.ToLower(s.CookieSecure) == "true" {
		secure = true
	}

	session := scs.New()
	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Name = s.CookieName
	session.Cookie.Domain = s.CookieDomain
	session.Cookie.Persist = persist
	session.Cookie.Secure = secure
	session.Cookie.SameSite = http.SameSiteLaxMode

	// TODO: Add support to different storage
	switch strings.ToLower(s.SessionType) {
	case "redis":
	default:
	}

	return session
}

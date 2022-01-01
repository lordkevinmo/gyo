package session

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"reflect"
	"testing"
)

func TestSession_InitSession(t *testing.T) {
	s := &Session{
		CookieDomain:   "localhost",
		CookieName:     "Gyo_",
		CookieLifetime: "120",
		CookiePersist:  "true",
		SessionType:    "cookie",
	}

	var sessionManager *scs.SessionManager
	var sessionKind reflect.Kind
	var sessionType reflect.Type

	ses := s.InitSession()

	reflectValue := reflect.ValueOf(ses)
	for reflectValue.Kind() == reflect.Ptr || reflectValue.Kind() == reflect.Interface {
		fmt.Println("For Loop:", reflectValue.Kind(), reflectValue.Type(), reflectValue)
		sessionKind = reflectValue.Kind()
		sessionType = reflectValue.Type()
		reflectValue = reflectValue.Elem()
	}

	if !reflectValue.IsValid() {
		t.Error("Invalid type or kind; kind:", reflectValue.Kind(), "type:", reflectValue.Type())
	}

	if sessionKind != reflect.ValueOf(sessionManager).Kind() {
		t.Error("Wrong kind returned testing cookie session. Expected",
			reflect.ValueOf(sessionManager).Kind(), "and got", sessionKind)
	}

	if sessionType != reflect.ValueOf(sessionManager).Type() {
		t.Error("Wrong type returned testing cookie session. Expected",
			reflect.ValueOf(sessionManager).Type(), "and got", sessionType)
	}
}

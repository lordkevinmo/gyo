package gyo

import "net/http"

func (g *Gyo) LoadSession(next http.Handler) http.Handler {
	return g.Session.LoadAndSave(next)
}

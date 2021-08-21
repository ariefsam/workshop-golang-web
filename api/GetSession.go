package api

import "todo/model"

func GetSession(token string) (user model.User, err error) {
	m.Lock()
	user, _ = Session[token]
	m.Unlock()
	return
}

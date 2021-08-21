package model

import (
	uuid "github.com/satori/go.uuid"
)

func GetUID() (uid string) {
	uids := uuid.NewV4()
	uid = uids.String()
	return
}

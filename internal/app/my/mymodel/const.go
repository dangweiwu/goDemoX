package mymodel

import (
	"strconv"
)

func GetAdminRedisId(id int) string {
	return strconv.Itoa(id)
}

// redis login id
func GetAdminRedisLoginId(id int) string {
	return "lgn:" + GetAdminRedisId(id)
}

const (
	ROLE_STATUS = "role:status:"
	ROLE_AUTH   = "role:auth:"
)

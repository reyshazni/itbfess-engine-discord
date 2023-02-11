package anonRepository

import "ITBFess/Database/redisdb"

const (
	Female = "Female"
	Male   = "Male"
	All    = "All"
)

func Insert(utype string, uid string) error {
	err := redisdb.PushData(utype, uid)
	return err
}

func Get(utype string) (string, error) {
	return redisdb.PopData(utype)
}

package models

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
)

var engine *redis.Pool

func init() {
	engine = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.DialURL("redis://:8541539655@47.88.10.29:6379/0")
			if err != nil {
				logrus.Error(err)
			}
			return
		},
	}
}

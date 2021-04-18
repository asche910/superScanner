package main

import (
	"testing"
)

func TestRedisSetAndGet(t *testing.T) {
	key := "test_set"
	rawVal := "hello,world"
	RedisSet(key, rawVal)
	val := RedisGet(key)

	if rawVal != val{
		t.Fail()
	}
}

func TestRedisPopAll(t *testing.T) {
	key := "IP_LIST"

	res := RedisPopAll(key)
	t.Log(res)
}
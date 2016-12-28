package goatbot

import (
	"github.com/patrickmn/go-cache"
	"time"
)

// Cache with default expiration of 12 hours, purging every 10 minutes
var kvs = cache.New(12*time.Hour, 10*time.Minute)

func Confirm(group string, user string) {
	set(group, user, true)
}

// TODO func Tentative(group, user string) {}

func Cancel(group string, user string) {
	set(group, user, false)
}

func ConfirmedList(group string) []string {
	keys := []string{}
	data, found := kvs.Get(group)
	if found {
		for k := range data.(map[string]bool) {
			keys = append(keys, k)
		}
	}
	return keys
}

func set(group string, user string, confirmed bool) {
	data, found := kvs.Get(group)
	var values map[string]bool
	if found {
		values = data.(map[string]bool)
	} else {
		values = make(map[string]bool)
	}
	if confirmed {
		values[user] = confirmed
	} else {
		delete(values, user)
	}
	kvs.Set(group, values, cache.DefaultExpiration)
}

package goatbot

// TODO Use Redis with 12 hours of expiration
var kvs = make(map[string]map[string]bool)

func Confirm(group string, user string) {
	set(group, user, true)
}

// TODO func Tentative(group, user string) {}

func Cancel(group string, user string) {
	set(group, user, false)
}

func ConfirmedList(group string) []string {
	keys := []string{}
	for k := range kvs[group] {
		keys = append(keys, k)
	}
	return keys
}

func set(group string, user string, confirmed bool) {
	if kvs[group] == nil {
		kvs[group] = make(map[string]bool)
	}
	if confirmed {
		kvs[group][user] = confirmed
	} else {
		delete(kvs[group], user)
	}
}

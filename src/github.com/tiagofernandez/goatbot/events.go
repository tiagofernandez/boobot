package goatbot

// TODO Use Redis with 12 hours of expiration
var kvs = make(map[string]map[string]bool)

func Going(group string, user string) {
	set(group, user, true)
}

func NotGoing(group string, user string) {
	set(group, user, false)
}

func AttendingList(group string) []string {
	keys := []string{}
	for k := range kvs[group] {
		keys = append(keys, k)
	}
	return keys
}

func set(group string, user string, attending bool) {
	if kvs[group] == nil {
		kvs[group] = make(map[string]bool)
	}
	if attending {
		kvs[group][user] = attending
	} else {
		delete(kvs[group], user)
	}
}

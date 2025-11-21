package cache

import "time"

type Item struct {
	Value      any
	Expiration int64
}

func (i Item) IsExpired() bool {
	if i.Expiration == 0 {
		return false
	}

	return time.Now().Unix() > i.Expiration
}

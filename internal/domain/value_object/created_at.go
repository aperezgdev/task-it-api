package valueobject

import "time"

type CreatedAt time.Time

func NewCreatedAt() CreatedAt {
	return CreatedAt(time.Now())
}
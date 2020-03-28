package entity

import "time"

// Post entity
type Post struct {
	id        string
	title     string
	body      string
	status    bool
	createdBy time.Time
	createdAt string
	updatedBy time.Time
	updatedAt string
}

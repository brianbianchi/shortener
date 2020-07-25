package main

import "time"

// URL model
type URL struct {
	Link        string    `json:"link"`
	Code        string    `json:"code"`
	Created     time.Time `json:"created"`
	Visited     int       `json:"visited"`
	LastVisited time.Time `json:"last_visited"`
}

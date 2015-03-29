package photo

import (
	"time"
)

type Photo struct {
	Id          int64
	AlbumId     int64
	Name        string
	Uploaded    time.Time
	Status      int
}

package photo

import (
	"time"
)

type Photo struct {
	Id       int64     `db:"id"`
	AlbumId  int64     `db:"album_id"`
	Name     string    `db:"name"`
	Uploaded time.Time `db:"uploaded"`
	Status   int64     `db:"status"`
}

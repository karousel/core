package file

type File struct {
	Id         int64  `db:"id"`
	ParentId   int64  `db:"parent_id"`
	Checksum   string `db:"checksum"`
	Resolution string `db:"resolution"`
}

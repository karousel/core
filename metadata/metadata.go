package metadata

type Metadata struct {
	Id       int64  `db:"id"`
	ParentId int64  `db:"parent_id"`
	Raw      string `db:"raw"`
}

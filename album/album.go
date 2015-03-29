package album

type Album struct {
	Id           int64  `db:"id"`
	CollectionId int64  `db:"collection_id"`
	Name         string `db:"name"`
}

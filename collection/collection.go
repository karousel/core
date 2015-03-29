package collection

type Collection struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}

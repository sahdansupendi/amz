package refid

type RefId struct {
	ID      string
	Flgused int
}

func (RefId) TableName() string {
	return "ref_id"
}

package repositories

const linksCollection = "links"

type LinkItem struct {
	ID   int64  `bson:"id"`
	Link string `bson:"link"`
}

type RangeItem struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

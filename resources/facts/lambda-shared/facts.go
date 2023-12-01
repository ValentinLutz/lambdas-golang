package shared

type FactEntity struct {
	Id   int    `db:"fact_id"`
	Text string `db:"fact_text"`
}

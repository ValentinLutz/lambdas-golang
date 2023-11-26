package shared

type FactRequest struct {
	Text string `json:"text"`
}

type FactResponse struct {
	Text string `json:"text"`
}

type FactEntity struct {
	Id   int    `db:"fact_id"`
	Text string `db:"fact_text"`
}

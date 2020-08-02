package analysis

type TypeInfo struct {
	StartPos int      `json:"start_pos"`
	EndPos   int      `json:"end_pos"`
	Values   []*Value `json:"values"`
}

type Value struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

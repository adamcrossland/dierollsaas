// roller
package roller

type RollSpec struct {
	Sides    int64
	DieCount int64
	Modifier int64
	Times    int64
}

type SetResult struct {
	Total int   `json:"total"`
	Dies  []int `json:"dies"`
}

type RollResults struct {
	Count int         `json:"count"`
	Rolls []SetResult `json:"rolls"`
}

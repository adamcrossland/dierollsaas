// roller
package roller

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type RollSpec struct {
	Sides    int64
	BestOf   int64
	DieCount int64
	Modifier int64
	Times    int64
}

type SetResult struct {
	Total int   `json:"total"`
	Count int   `json:"count"`
	Dies  []int `json:"dice"`
}

type RollResults struct {
	Count int         `json:"count"`
	Rolls []SetResult `json:"rolls"`
}

func (results RollResults) ToJSON() []byte {
	coded, _ := json.Marshal(results)
	return coded
}

func (results RollResults) ToText() []byte {
	var buf bytes.Buffer

	for er := 0; er < results.Count; er++ {
		if er > 0 {
			buf.WriteString("; ")
		}
		buf.WriteString(fmt.Sprintf("%d", results.Rolls[er].Total))
		if results.Rolls[er].Count > 1 {
			for ed := 0; ed < results.Rolls[er].Count; ed++ {
				if ed == 0 {
					buf.WriteString(": ")
				} else {
					buf.WriteString(" ")
				}
				buf.WriteString(fmt.Sprintf("[%d]", results.Rolls[er].Dies[ed]))
			}
		}
	}

	return buf.Bytes()
}

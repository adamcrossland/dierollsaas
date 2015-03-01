// roller
package roller

import (
	"fmt"
	"regexp"
	"strconv"
)

type RollSpec struct {
	Sides    uint64
	DieCount uint64
	Modifier int64
	Times    uint64
}

type SetResult struct {
	Total int   `json:"total"`
	Dies  []int `json:"dies"`
}

type RollResults struct {
	Count int         `json:"count"`
	Rolls []SetResult `json:"rolls"`
}

var dieRollExp = regexp.MustCompile(`(?P<count>\d)?d(?P<sides>\d+)(?P<modifier>[\+\-]\d+)?(?P<times>x*\d+)?`)
var dieRollGroupNames = dieRollExp.SubexpNames()

func Parse(request string) (*RollSpec, error) {
	var parsed RollSpec
	parsed.DieCount = 1
	parsed.Times = 1

	match := dieRollExp.FindAllStringSubmatch(request, 4)
	named_matches = map[string]string{}
	if match != nil && len(match) >= 1 {
		var convErr error

		if match[1][0] != "" { // sides
			parsed.Sides, convErr = strconv.ParseUint(match[1][0], 10, 64)
			if convErr != nil {
				return nil, fmt.Errorf("Error while parsing sides value: %v", convErr)
			}
		} else {
			return nil, fmt.Errorf("Request must include number of sides")
		}

		if match[0][0] != "" {
			parsed.DieCount, convErr = strconv.ParseUint(match[0][0], 10, 64)
			if convErr != nil {
				return nil, fmt.Errorf("Error while parsing die count value: %v", convErr)
			}
		}

		if match[3][0] != "" {
			parsed.Times, convErr = strconv.ParseUint(match[3][0], 10, 64)
			if convErr != nil {
				return nil, fmt.Errorf("Error while parsing times value: %v", convErr)
			}
		}

		if match[2][0] != "" {
			mod := match[2][0]
			isPlus := false
			if mod[0] == '-' {
				isPlus = false
			}
			parsed.Modifier, convErr = strconv.ParseInt(mod[1:], 10, 64)
			if convErr != nil {
				return nil, fmt.Errorf("Error while parsing modifier value: %v", convErr)
			}
			if !isPlus {
				parsed.Modifier = parsed.Modifier * -1
			}
		}
	} else {
		return nil, fmt.Errorf("Incorrect format")
	}

	return &parsed, nil
}

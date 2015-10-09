// roller_test.go
package roller

import (
	"testing"
)

func TestParse01(t *testing.T) {
	result, err := Parse("d6")
	if err != nil {
		t.Errorf("err returned: %v", err)
	}
	if result.Sides != 6 {
		t.Errorf("Got %d sides instead of 6", result.Sides)
	}
	if result.DieCount != 1 {
		t.Errorf("Got %d count instead of 1", result.DieCount)
	}
	if result.Modifier != 0 {
		t.Errorf("Got modifier %d instead of 0", result.Modifier)
	}
	if result.Times != 1 {
		t.Errorf("Got %d times instead of 1", result.Times)
	}
}

func TestParse02(t *testing.T) {
	result, err := Parse("2d6")
	if err != nil {
		t.Errorf("err returned: %v", err)
	}
	if result.Sides != 6 {
		t.Errorf("Got %d sides instead of 6", result.Sides)
	}
	if result.DieCount != 2 {
		t.Errorf("Got %d count instead of 2", result.DieCount)
	}
	if result.Modifier != 0 {
		t.Errorf("Got modifier %d instead of 0", result.Modifier)
	}
	if result.Times != 1 {
		t.Errorf("Got %d times instead of 1", result.Times)
	}
}

func TestParse03(t *testing.T) {
	result, err := Parse("d4+1")
	if err != nil {
		t.Errorf("err returned: %v", err)
	}
	if result.Sides != 4 {
		t.Errorf("Got %d sides instead of 4", result.Sides)
	}
	if result.DieCount != 1 {
		t.Errorf("Got %d count instead of 1", result.DieCount)
	}
	if result.Modifier != 1 {
		t.Errorf("Got modifier %d instead of +1", result.Modifier)
	}
	if result.Times != 1 {
		t.Errorf("Got %d times instead of 1", result.Times)
	}
}

func TestParse04(t *testing.T) {
	result, err := Parse("2d10+2")
	if err != nil {
		t.Errorf("err returned: %v", err)
	}
	if result.Sides != 10 {
		t.Errorf("Got %d sides instead of 10", result.Sides)
	}
	if result.DieCount != 2 {
		t.Errorf("Got %d count instead of 2", result.DieCount)
	}
	if result.Modifier != 2 {
		t.Errorf("Got modifier %d instead of +2", result.Modifier)
	}
	if result.Times != 1 {
		t.Errorf("Got %d times instead of 1", result.Times)
	}
}

func TestParse05(t *testing.T) {
	result, err := Parse("d4-1")
	if err != nil {
		t.Errorf("err returned: %v", err)
	}
	if result.Sides != 4 {
		t.Errorf("Got %d sides instead of 4", result.Sides)
	}
	if result.DieCount != 1 {
		t.Errorf("Got %d count instead of 1", result.DieCount)
	}
	if result.Modifier != -1 {
		t.Errorf("Got modifier %d instead of -1", result.Modifier)
	}
	if result.Times != 1 {
		t.Errorf("Got %d times instead of 1", result.Times)
	}
}

func TestParse06(t *testing.T) {
	result, err := Parse("2d12-2")
	if err != nil {
		t.Errorf("err returned: %v", err)
	}
	if result.Sides != 12 {
		t.Errorf("Got %d sides instead of 12", result.Sides)
	}
	if result.DieCount != 2 {
		t.Errorf("Got %d count instead of 2", result.DieCount)
	}
	if result.Modifier != -2 {
		t.Errorf("Got modifier %d instead of -2", result.Modifier)
	}
	if result.Times != 1 {
		t.Errorf("Got %d times instead of 1", result.Times)
	}
}

func TestParse07(t *testing.T) {
	result, err := Parse("4d6x6")
	if err != nil {
		t.Errorf("err returned: %v", err)
	}
	if result.Sides != 6 {
		t.Errorf("Got %d sides instead of 6", result.Sides)
	}
	if result.DieCount != 4 {
		t.Errorf("Got %d count instead of 4", result.DieCount)
	}
	if result.Modifier != 0 {
		t.Errorf("Got modifier %d instead of 0", result.Modifier)
	}
	if result.Times != 6 {
		t.Errorf("Got %d times instead of 6", result.Times)
	}
}

func TestParse08(t *testing.T) {
	result, err := Parse("2d12-2x8")
	if err != nil {
		t.Errorf("err returned: %v", err)
	}
	if result.Sides != 12 {
		t.Errorf("Got %d sides instead of 12", result.Sides)
	}
	if result.DieCount != 2 {
		t.Errorf("Got %d count instead of 2", result.DieCount)
	}
	if result.Modifier != -2 {
		t.Errorf("Got modifier %d instead of -2", result.Modifier)
	}
	if result.Times != 8 {
		t.Errorf("Got %d times instead of 8", result.Times)
	}
}

func TestParse09(t *testing.T) {
	result, err := Parse("1d100")
	if err != nil {
		t.Errorf("err returned: %v", err)
	}
	if result.Sides != 100 {
		t.Errorf("Got %d sides instead of 100", result.Sides)
	}
	if result.DieCount != 1 {
		t.Errorf("Got %d count instead of 1", result.DieCount)
	}
	if result.Modifier != 0 {
		t.Errorf("Got modifier %d instead of 0", result.Modifier)
	}
	if result.Times != 1 {
		t.Errorf("Got %d times instead of 1", result.Times)
	}
}

func TestParse10(t *testing.T) {
	result, err := Parse("20d100")
	if err != nil {
		t.Errorf("err returned: %v", err)
	}
	if result.Sides != 100 {
		t.Errorf("Got %d sides instead of 100", result.Sides)
	}
	if result.DieCount != 20 {
		t.Errorf("Got %d count instead of 20", result.DieCount)
	}
	if result.Modifier != 0 {
		t.Errorf("Got modifier %d instead of 0", result.Modifier)
	}
	if result.Times != 1 {
		t.Errorf("Got %d times instead of 1", result.Times)
	}
}

func TestParse11(t *testing.T) {
	result, err := Parse("3,4d6")
	if err != nil {
		t.Errorf("err returned: %v", err)
	}
	if result.Sides != 6 {
		t.Errorf("Got %d sides instead of 6", result.Sides)
	}
	if result.DieCount != 4 {
		t.Errorf("Got %d count instead of 4", result.DieCount)
	}
	if result.Modifier != 0 {
		t.Errorf("Got modifier %d instead of 0", result.Modifier)
	}
	if result.Times != 1 {
		t.Errorf("Got %d times instead of 1", result.Times)
	}
	if result.BestOf != 3 {
		t.Errorf("Got %d best-of instead of 3", result.BestOf)
	}
}

func TestParse12(t *testing.T) {
	result, err := Parse("3,4d6x10")
	if err != nil {
		t.Errorf("err returned: %v", err)
	}
	if result.Sides != 6 {
		t.Errorf("Got %d sides instead of 6", result.Sides)
	}
	if result.DieCount != 4 {
		t.Errorf("Got %d count instead of 4", result.DieCount)
	}
	if result.Modifier != 0 {
		t.Errorf("Got modifier %d instead of 0", result.Modifier)
	}
	if result.Times != 10 {
		t.Errorf("Got %d times instead of 10", result.Times)
	}
	if result.BestOf != 3 {
		t.Errorf("Got %d best-of instead of 3", result.BestOf)
	}
}

func TestParse13(t *testing.T) {
	result, err := Parse("3,4d6+1x10")
	if err != nil {
		t.Errorf("err returned: %v", err)
	}
	if result.Sides != 6 {
		t.Errorf("Got %d sides instead of 6", result.Sides)
	}
	if result.DieCount != 4 {
		t.Errorf("Got %d count instead of 4", result.DieCount)
	}
	if result.Modifier != 1 {
		t.Errorf("Got modifier %d instead of 1", result.Modifier)
	}
	if result.Times != 10 {
		t.Errorf("Got %d times instead of 10", result.Times)
	}
	if result.BestOf != 3 {
		t.Errorf("Got %d best-of instead of 3", result.BestOf)
	}
}

func TestParse14(t *testing.T) {
	result, err := Parse("3,4D6+1X10")
	if err != nil {
		t.Errorf("err returned: %v", err)
	}
	if result.Sides != 6 {
		t.Errorf("Got %d sides instead of 6", result.Sides)
	}
	if result.DieCount != 4 {
		t.Errorf("Got %d count instead of 4", result.DieCount)
	}
	if result.Modifier != 1 {
		t.Errorf("Got modifier %d instead of 1", result.Modifier)
	}
	if result.Times != 10 {
		t.Errorf("Got %d times instead of 10", result.Times)
	}
	if result.BestOf != 3 {
		t.Errorf("Got %d best-of instead of 3", result.BestOf)
	}
}

func TestRolls01(t *testing.T) {
	spec, err := Parse("d6x100")
	if err != nil {
		t.Errorf("err returned: %v", err)
	}

	rolls := DoRolls(*spec)
	if rolls.Count != 100 {
		t.Errorf("Got %d rolls back instead of 100", rolls.Count)
	}
	if len(rolls.Rolls) != rolls.Count {
		t.Errorf("Mismatch between Count (%d) and len(%d)", rolls.Count, len(rolls.Rolls))
	}

	for testEach := 0; testEach < rolls.Count; testEach++ {
		dieRolled := rolls.Rolls[testEach].Dies[0]
		if dieRolled < 1 || dieRolled > int(spec.Sides) {
			t.Errorf("Got value %d on a d%d", dieRolled, spec.Sides)
		}
	}
}

func TestRolls02(t *testing.T) {
	spec, err := Parse("2d6x100")
	if err != nil {
		t.Errorf("err returned: %v", err)
	}

	rolls := DoRolls(*spec)
	if rolls.Count != 100 {
		t.Errorf("Got %d rolls back instead of 100", rolls.Count)
	}
	if len(rolls.Rolls) != rolls.Count {
		t.Errorf("Mismatch between Count (%d) and len(%d)", rolls.Count, len(rolls.Rolls))
	}

	for testEachSet := 0; testEachSet < rolls.Count; testEachSet++ {
		eachSet := rolls.Rolls[testEachSet]
		if eachSet.Count != 2 {
			t.Errorf("Set had Count of %d instead of 2", eachSet.Count)
		}
		if eachSet.Count != len(eachSet.Dies) {
			t.Errorf("Mismatch between Count (%d) and len(%d) in a set result", eachSet.Count, len(eachSet.Dies))
		}
		dieTotals := 0
		for eachDie := 0; eachDie < eachSet.Count; eachDie++ {
			dieRolled := eachSet.Dies[eachDie]
			if dieRolled < 1 || dieRolled > int(spec.Sides) {
				t.Errorf("Got value %d on a d%d", dieRolled, spec.Sides)
			}
			dieTotals += dieRolled
		}
		if dieTotals < 2 || dieTotals > 12 {
			t.Errorf("Got out-of-range(2-12) die total of %d", dieTotals)
		}
		if dieTotals != eachSet.Total {
			t.Errorf("Calculated die total %d did not match returned value %d", dieTotals, eachSet.Total)
		}
	}
}

func TestBadParse01(t *testing.T) {
	_, err := Parse("dfhjsdjh")
	if err == nil {
		t.Errorf("Bad spec format should have failed", err)
	}
}

func TestBadParse02(t *testing.T) {
	// Can't have a best-of modifier with a number of dice
	_, err := Parse("3,d6")
	if err == nil {
		t.Errorf("Bad spec format should have failed", err)
	}
}

func TestBadParse03(t *testing.T) {
	// Can't have a best-of slash without a number in front of it.
	_, err := Parse(",4d6")
	if err == nil {
		t.Errorf("Bad spec format should have failed", err)
	}
}

func TestBadParse04(t *testing.T) {
	_, err := Parse("1,2d,3")
	if err == nil {
		t.Errorf("Bad spec format should have failed", err)
	}
}

func TestBadParse05(t *testing.T) {
	// The best-of modifier must be smaller than the number of dice
	_, err := Parse("4,3d6")
	if err == nil {
		t.Errorf("Bad spec format should have failed", err)
	}
}

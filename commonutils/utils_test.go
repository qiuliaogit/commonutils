package commonutils

import "testing"

func Test_CalcMoneyByCent2(t *testing.T) {
	v := []int64{123, 456, 789, 120, 130, 131, 140, 151, 109}
	dest := []string{"1.23", "4.56", "7.89", "1.20", "1.30", "1.31", "1.40", "1.51", "1.09"}
	dest1 := []string{"1.23", "4.56", "7.89", "1.2", "1.3", "1.31", "1.4", "1.51", "1.09"}
	for i, val := range v {
		result2 := CalcMoneyByCent2(val)
		result := CalcMoneyByCent(val)
		if result2 != dest[i] {
			t.Errorf("CalcMoneyByCent2(%d) = %s, want %s", val, result2, dest[i])
		} else {
			t.Logf("CalcMoneyByCent2(%d) => %s", val, result2)
		}
		if result != dest1[i] {
			t.Errorf("CalcMoneyByCent(%d) = %s, want %s", val, result, dest1[i])
		}
	}
}

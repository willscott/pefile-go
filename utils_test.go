package pefile

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

var EPSILON float64 = 0.00000001

func floatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}

func TestGetEntropy(t *testing.T) {
	tests := map[string]float64{
		"00b6ea24092c43db96e4dec79dfcdafd301c78a3d0ebaa27d8d5e4934793876d": 5.97361676112134,
		"a30fc540b7237f64b3fc07afae610a8aa7160ca614ede882b5800075c98dfe20": 7.493161152118229,
		"1809035eb26fa063a9baba068417f2c1733c4531a2409a4f6ccdc27958d8dbf3": 6.489522264127895,
		"d5d75e2b10f252cddebd110c0b48fc15a87fa1ba4b8cc22646c46557eae377d4": 5.864521602606575,
		"6b153122ce91ef8897fb850ccaf54355f84ac8e4da9f23b197135fb7cccd9655": 6.326375779160963,
	}

	for pename, expect := range tests {
		data, err := ioutil.ReadFile(filepath.Join(TEST_EXE_PATH, pename))
		if err != nil {
			t.Fatal(err)
		}
		entropy := GetEntropy(data)
		if !floatEquals(entropy, expect) {
			t.Errorf("got=%g, expect=%g", entropy, expect)
		}
	}
}

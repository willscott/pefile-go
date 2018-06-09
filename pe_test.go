package pefile

import (
	"path/filepath"
	"testing"
)

const TEST_EXE_PATH = "exe_test_files"

func TestGetImpHash(t *testing.T) {
	tests := map[string]string{
		"00b6ea24092c43db96e4dec79dfcdafd301c78a3d0ebaa27d8d5e4934793876d": "02c3f6b4018572e7a446fe853ac114e2",
		"a30fc540b7237f64b3fc07afae610a8aa7160ca614ede882b5800075c98dfe20": "c86b02c21ff392ad6ffcf21dcd4a5588",
		"1809035eb26fa063a9baba068417f2c1733c4531a2409a4f6ccdc27958d8dbf3": "e45f42af6cb5a2d26bc86b5052850b0c",
		"d5d75e2b10f252cddebd110c0b48fc15a87fa1ba4b8cc22646c46557eae377d4": "02c3f6b4018572e7a446fe853ac114e2",
		"6b153122ce91ef8897fb850ccaf54355f84ac8e4da9f23b197135fb7cccd9655": "36c88d39e09db05c09a463536218ae23",
	}

	for pename, expect := range tests {
		pe, err := NewPEFile(filepath.Join(TEST_EXE_PATH, pename))
		if err != nil {
			t.Fatal(err)
		}
		impHash := pe.GetImpHash()
		if impHash != expect {
			t.Errorf("got=%s, expect=%s", impHash, expect)
		}
	}

}

package pefile

import (
	"path/filepath"
	"testing"
)

type Hash struct {
	MD5, SHA1, SHA256 string
}

const (
	EPSILON       = 0.00000001
	TEST_EXE_PATH = "exe_test_files"
)

func floatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}

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

func TestGetEntropy(t *testing.T) {
	tests := map[string]float64{
		"00b6ea24092c43db96e4dec79dfcdafd301c78a3d0ebaa27d8d5e4934793876d": 5.97361676112134,
		"a30fc540b7237f64b3fc07afae610a8aa7160ca614ede882b5800075c98dfe20": 7.493161152118229,
		"1809035eb26fa063a9baba068417f2c1733c4531a2409a4f6ccdc27958d8dbf3": 6.489522264127895,
		"d5d75e2b10f252cddebd110c0b48fc15a87fa1ba4b8cc22646c46557eae377d4": 5.864521602606575,
		"6b153122ce91ef8897fb850ccaf54355f84ac8e4da9f23b197135fb7cccd9655": 6.326375779160963,
	}

	for pename, expect := range tests {
		pe, err := NewPEFile(filepath.Join(TEST_EXE_PATH, pename))
		if err != nil {
			t.Fatal(err)
		}
		entropy := pe.GetEntropy()
		if !floatEquals(entropy, expect) {
			t.Errorf("got=%g, expect=%g", entropy, expect)
		}
	}
}

func TestHashFunctions(t *testing.T) {
	tests := map[string]Hash{
		"00b6ea24092c43db96e4dec79dfcdafd301c78a3d0ebaa27d8d5e4934793876d": Hash{
			MD5:    "5514b37c23a880f6c47830ad4fbad7e9",
			SHA1:   "92886f1b335da758374db43f9d0969697896beb9",
			SHA256: "00b6ea24092c43db96e4dec79dfcdafd301c78a3d0ebaa27d8d5e4934793876d",
		},
		"a30fc540b7237f64b3fc07afae610a8aa7160ca614ede882b5800075c98dfe20": Hash{
			MD5:    "06da6b63a499506538eb47d75db9414f",
			SHA1:   "9eea9a0ff0b5f7e40e16e4d731e4543c9aeacf80",
			SHA256: "a30fc540b7237f64b3fc07afae610a8aa7160ca614ede882b5800075c98dfe20",
		},
		"1809035eb26fa063a9baba068417f2c1733c4531a2409a4f6ccdc27958d8dbf3": Hash{
			MD5:    "f3090ee3bcf2c38630f2f9ae8ae30522",
			SHA1:   "c1a3dc90f4af5d487ee3419bafa9bed6b63a3457",
			SHA256: "1809035eb26fa063a9baba068417f2c1733c4531a2409a4f6ccdc27958d8dbf3",
		},
	}

	for pename, expect := range tests {
		pe, err := NewPEFile(filepath.Join(TEST_EXE_PATH, pename))
		if err != nil {
			t.Fatal(err)
		}

		t.Run("MD5", func(t *testing.T) {
			md5Hash := pe.GetMD5Hash()
			if md5Hash != expect.MD5 {
				t.Errorf("got=%s, expect=%s", md5Hash, expect.MD5)
			}
		})

		t.Run("SHA1", func(t *testing.T) {
			sha1Hash := pe.GetSHA1Hash()
			if sha1Hash != expect.SHA1 {
				t.Errorf("got=%s, expect=%s", sha1Hash, expect.SHA1)
			}
		})

		t.Run("SHA256", func(t *testing.T) {
			sha2Hash := pe.GetSHA256Hash()
			if sha2Hash != expect.SHA256 {
				t.Errorf("got=%s, expect=%s", sha2Hash, expect.SHA256)
			}
		})
	}
}

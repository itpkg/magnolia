package i18n_test

import (
	"os"
	"testing"

	"github.com/itpkg/magnolia/web/i18n"
)

func TestGenerate(t *testing.T) {
	for _, lng := range []string{"en-US", "zh-Hans"} {
		if err := i18n.Generate(lng); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.RemoveAll("locales"); err != nil {
		t.Fatal(err)
	}
}

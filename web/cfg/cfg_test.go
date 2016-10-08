package cfg_test

import(
	"testing"

	"github.com/itpkg/magnolia/web/cfg"
)

const file = "config.toml"

func TestConfig(t *testing.T){
	cfg.Int("int_v", 123)
	cfg.String("string_v", "hello")
	if err:=cfg.Write(file); err!=nil{
		t.Fatalf()
	}
}

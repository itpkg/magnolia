package cfg_test

import(
	"testing"
	"time"
	"os"

	"github.com/itpkg/magnolia/web/cfg"
)

const file = "config.toml"

func TestConfig(t *testing.T){
	os.Remove(file)

	cfg.SetDefault("int", 123)
	cfg.SetDefault("string", "hello")
	cfg.SetDefault("strings", []string{"aaa", "bbb", "ccc"})
	cfg.SetDefault("map", map[string]interface{}{"aaa":111, "bbb":time.Now(), "ccc":"aaa"})
	if err:=cfg.Write(file); err!=nil{
		t.Fatal(err)
	}
	if err:=cfg.Read(file);err!=nil{
		t.Fatal(err)
	}
	t.Log(cfg.GetString("string"))
	t.Log(cfg.GetInt("int"))
	t.Logf("%+v", cfg.GetArray("strings"))
	tmp := cfg.GetMap("map")
	t.Logf("%+v", tmp)
	t.Log(tmp["aaa"])
}

package cfg

import(
	"os"

	"github.com/BurntSushi/toml"
)

func IsProd()bool{
	return GetString("env") == "prod"
}


func Read(file string) error{
	_, err:= toml.DecodeFile(file, &vars)
	return err
}

func Write(file string) error{
	fd, err:= os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_EXCL,0600)
	if err!=nil{
		return err
	}
	defer fd.Close()
	return toml.NewEncoder(fd).Encode(vars)
}

//-----------------------------------------------------------------------------

var vars map[string]interface{}

func init(){
  vars = make(map[string]interface{})
}

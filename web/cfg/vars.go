package cfg

//-----------------------------------------------------------------------------

var vars map[string]interface{}

func IsProd()bool{
	return String("env", "dev") == "prod"
}


func Load(file string) error{
	return nil
}

func Write(file string) error{
	return nil
}

func init(){
  vars = make(map[string])interface{}
}

package cfg

func Int(key string, def int){
	if val, ok := vars[key]; ok{
		return ok.(int)
	}
	vars[key] = def
	return def
}

package cfg

func GetString(key string) string {
	if val, ok := vars[key]; ok {
		return val.(string)
	}
	return ""
}

func GetInt(key string) int64 {
	if val, ok := vars[key]; ok {
		return val.(int64)
	}
	return 0
}

func GetArray(key string) []interface{} {
	if val, ok := vars[key]; ok {
		return val.([]interface{})
	}
	return nil
}

func GetMap(key string) map[string]interface{} {
	if val, ok := vars[key]; ok {
		return val.(map[string]interface{})
	}
	return nil
}

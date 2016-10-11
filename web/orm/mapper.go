package orm

//-----------------------------------------------------------------------------
var queries map[string]string

func init() {
	queries = make(map[string]string)
}

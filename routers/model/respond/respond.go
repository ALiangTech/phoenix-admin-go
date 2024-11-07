package respond

type Response struct {
	Data interface{}
	Err  string
	Msg  []interface{}
}

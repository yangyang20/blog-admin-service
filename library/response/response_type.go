package response

const (
	SUCCESS = 200
	Warning = 300
	ERROR   = 400
	FAIL    = 500
)

type ResultType struct {
	Code int
	Msg  string
	Data interface{}
}

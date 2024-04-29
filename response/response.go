package response

type Response interface {
	StatusCode() int
	GetBody() ([]byte, error)
	GetData() interface{}
	Error() string
}

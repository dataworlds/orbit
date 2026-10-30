package orbit

type server struct {
	Allow string
}

func New() *server {
	return &server{}
}

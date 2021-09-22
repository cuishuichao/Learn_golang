package mock

import "fmt"

type Receiver struct{
	Contents string
}

func(r *Receiver) String() string{
	return fmt.Sprintf("retriever: {Contents: %s}", r.Contents)
}

func (r *Receiver) Post(url string, from map[string]string) string {
	r.Contents = from["contents"]
	return "ok"
}

func (r *Receiver) Get(url string) string {
	return r.Contents
}


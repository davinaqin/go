package mock
type MockRetriver struct{
	Contents string
}

func (r MockRetriver) Get(url string) string {
	return r.Contents
}
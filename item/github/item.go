package github

type StarRepository struct {
	Author string `csv:"author"`
	Name   string `csv:"name"`
	Link   string `csv:"link"`
	Desc   string `csv:"desc"`
}

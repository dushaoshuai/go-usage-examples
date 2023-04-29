package third_party

type FooErr struct{}

func (e FooErr) ok() bool {
	return e != FooErr{}
}

func (e FooErr) Ok() bool {
	return e.ok()
}

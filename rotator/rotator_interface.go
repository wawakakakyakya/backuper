package rotator

type actRotator interface {
	_remove(fpath string) error
	_find() ([]string, error)
	_run(ar actRotator) error
}

type rotatorInterface interface {
	actRotator
	find(ar actRotator) ([]string, error)
	remove(files []string, ar actRotator) error
	Run() error
}

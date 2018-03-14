package isp

type Worker interface {
	Work(app *Application) error
}

type Working func(app *Application) error
func (s Working) Work(app *Application) error {
	return s(app)
}

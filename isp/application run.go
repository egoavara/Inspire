package isp

type Worker interface {
	Work(app *Application, dt int64) error
}

type Working func(app *Application) error
func (s Working) Work(app *Application) error {
	return s(app)
}

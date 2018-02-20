package egn

import "github.com/iamGreedy/Inspire/igl"

type Engine struct {
	App *igl.Application
	Draw func(engine *Engine) error

}


package igl

type Model interface {
	Render()
	Act(motion Motion)
}

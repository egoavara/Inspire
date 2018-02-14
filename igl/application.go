package igl

type Application interface {
	Run()
	Stop()
	//
	UseShader(shader Shader)
}
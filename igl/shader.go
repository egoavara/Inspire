package igl

type Shader interface {
	Use()
	Release()
	Close()

}


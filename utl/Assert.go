package utl

func Must(err error)  {
	if err != nil {
		panic(err)
	}
}

func MustValue(value interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return value
}
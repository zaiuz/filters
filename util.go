package filters

func noError(e error) {
	if e != nil {
		panic(e)
	}
}

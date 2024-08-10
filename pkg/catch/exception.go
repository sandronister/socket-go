package catch

func Exception(err error) {
	if err != nil {
		panic(err)
	}
}

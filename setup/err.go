package setup

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
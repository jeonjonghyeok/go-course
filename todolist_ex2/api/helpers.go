package api

func must(err error) {
	if err == db.ErrNotFound {
		return
	} else if err != nil {
		panic(internalError)
	}
}

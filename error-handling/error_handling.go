package erratum

const testVersion = 2

func Use(o ResourceOpener, input string) (err error) {
	var res Resource
	res, err = openResource(o)
	if err != nil {
		return
	}
	defer res.Close()

	defer func() {
		if e := recover(); e != nil {
			if fe, ok := e.(FrobError); ok {
				res.Defrob(fe.defrobTag)
				err = fe
			}
			err = e.(error)
		}
	}()

	// may panic
	res.Frob(input)
	return
}

// openResource keep trying if TransientError occurred
func openResource(o ResourceOpener) (Resource, error) {
	// here we use for loop, recursive call is better,
	// but since Go doesn't implement Tail Call Optimization,
	// recursive call may cause Stack Overflow
	for {
		res, err := o()
		if err == nil {
			return res, nil
		}
		if _, ok := err.(TransientError); ok {
			continue
		}
		return nil, err
	}
}

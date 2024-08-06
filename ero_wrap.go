package ero

func (eme *eroMsgError) Wrap(err error, text string) error {
	if err == nil {
		return nil
	}
	return &eroMsgError{
		msg: text,
		lt:  getLineTrace(),
		c:   err,
	}
}

func Wrap(err error, text string) error {
	e := mustAs[interface{ Wrap(error, string) error }](err)
	return e.Wrap(err, text)
}

func (eme *eroMsgError) UnwrapOnce(err error) error {
	if err == nil {
		return nil
	}
	e, ok := as[interface{ Child() error }](err)
	if !ok {
		return err
	}
	return e.Child()
}

func UnwrapOnce(err error) error {
	e := mustAs[interface{ UnwrapOnce(error) error }](err)
	return e.UnwrapOnce(err)
}

func UnwrapAll(err error) (errs []error) {
	if err == nil {
		return
	}
	walkUnwrap(err, func(err error, _ bool) bool {
		errs = append(errs, err)
		return true
	})
	return errs
}

func walkUnwrap(err error, fun func(err error, isLast bool) bool) {
	if err != nil {
		e := mustAs[interface{ Child() error }](err)
		child := e.Child()
		isLast := child == nil
		if !fun(err, isLast) {
			return
		}
		walkUnwrap(child, fun)
	}
}

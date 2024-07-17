package async_wrapper

import (
	"context"
	"time"
)

type AsyncWrapperCvt[oldT, newT any] struct {
	old   AsyncResult[oldT]
	cvtFn func(od oldT, oe error) (nd newT, ne error)
}

func (w *AsyncWrapperCvt[oldT, newT]) SetContext(ctx context.Context) AsyncResult[newT] {
	w.old.SetContext(ctx)
	return w
}

func (w *AsyncWrapperCvt[oldT, newT]) SetTimeout(timeout time.Duration) AsyncResult[newT] {
	w.old.SetTimeout(timeout)
	return w
}

func (w *AsyncWrapperCvt[oldT, newT]) AsyncGetResult(callback func(ret newT, err error)) {
	w.old.AsyncGetResult(func(od oldT, oe error) {
		callback(w.cvtFn(od, oe))
	})
}

func (w *AsyncWrapperCvt[oldT, newT]) BlockGetResult() (ret newT, err error) {
	od, oe := w.old.BlockGetResult()
	return w.cvtFn(od, oe)
}

func (w *AsyncWrapperCvt[oldT, newT]) RedirectResult(reciver chan struct {
	ret newT
	err error
}, block bool) {
	w.old.AsyncGetResult(func(od oldT, oe error) {
		ret, err := w.cvtFn(od, oe)
		reciver <- struct {
			ret newT
			err error
		}{
			ret, err,
		}
	})
}

func TransResultType[oldT, newT any](
	old AsyncResult[oldT],
	cvtFn func(oldVal oldT, oldErr error) (newVal newT, newErr error),
) AsyncResult[newT] {
	return &AsyncWrapperCvt[oldT, newT]{
		old, cvtFn,
	}
}

package test

//下面的迭代会有什么问题
//没有缓冲的chan，读写都会阻塞，如果外部调用了这个方案，但米有从chan读出，会产生一个阻塞的goroutine
//func (set *threadSafeSet) Iter() <-chan interface{} {
//	ch := make(chan interface{})
//	go func() {
//		set.RLock()
//
//		for elem := range set.s {
//			ch <- elem
//		}
//
//		close(ch)
//		set.RUnlock()
//
//	}()
//	return ch
//}

package queue

//
//LazyQueue struct ，使用channel 实现异步操作
//
type LazyQueue struct {
	qChan         chan interface{}
	DequeueMethod DequeueMethod
}

//DequeueMethod is
//出队列的操作
//
type DequeueMethod func(interface{})

//NewLazyQueue is
//new 一个lazyqueue 指定channel的长度 和 dequeue 的操作
//grcount 执行dequeue 的 goruntine 的总数
func NewLazyQueue(qsize int, dequeueMethod DequeueMethod, grcount int) *LazyQueue {
	qInstance := &LazyQueue{}
	qInstance.qChan = make(chan interface{}, qsize)
	qInstance.DequeueMethod = dequeueMethod
	//
	if grcount <= 0 {
		grcount = 1
	}
	for i := 0; i < grcount; i++ {
		go qInstance.startDequeue()
	}

	return qInstance
}

//Enqueue is
// 入队列的操作
func (lq *LazyQueue) Enqueue(qinfo interface{}) {
	lq.qChan <- qinfo
}

func (lq *LazyQueue) startDequeue() {
	for {
		select {
		case bm := <-lq.qChan:
			lq.DequeueMethod(bm)
		}
	}
}

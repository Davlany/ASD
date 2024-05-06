package queue

type QueueArray struct {
	pull  [255]int
	start int
	end   int
}

func (qa *QueueArray) Push(num int) {
	if qa.end < 0 {
		qa.end = 254
		qa.pull[qa.end] = num
	} else {
		qa.pull[qa.end] = num
	}
	qa.end--
}

func (qa *QueueArray) Get() int {
	if qa.start < 0 {
		qa.start = 254
		res := qa.pull[qa.start]
		qa.pull[qa.start] = 0
		qa.start--
		return res
	} else {
		res := qa.pull[qa.start]
		qa.pull[qa.start] = 0
		qa.start--
		return res
	}
}

func NewQueueArray() *QueueArray {
	return &QueueArray{
		start: 254,
		end:   254,
	}
}



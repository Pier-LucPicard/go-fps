package base

import 
(
	"time"
)

const  SECOND int64 = 1000000000;

type Time struct {
	delta int64
}

func GetTime() int64 {
	return time.Now().UnixNano()
}

func (t *Time) GetDelta() int64{
	return t.delta
}

func (t *Time) SetDelta(delta int64){
	t.delta = delta
}
package task2

const (
	FirstIntoSecond = "you can put the first envelop into the second"
	SecondIntoFirst = "you can put the second envelop into the first"
	CantCompare     = "you can't put one envelop into another"
)

type Envelop struct {
	height, width float32
}

func (it *Envelop) IsBiggerThan(o *Envelop) bool {
	return it.height > o.height && it.width > o.width
}

func (it *Envelop) MakeHeightBigger() *Envelop {
	if it.height < it.width {
		it.height, it.width = it.width, it.height
	}

	return it
}

func NewEnvelop(h, w float32) *Envelop {
	e := Envelop{h, w}
	return e.MakeHeightBigger()
}

func EnvelopComparisonStr(env1, env2 *Envelop) string {
	if env1.IsBiggerThan(env2) {
		return SecondIntoFirst
	} else if env2.IsBiggerThan(env1) {
		return FirstIntoSecond
	}
	return CantCompare
}

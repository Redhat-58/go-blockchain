package types

type NewBlock struct {
	Data string
}

type NewTranasaction struct {
	To     string
	From   string
	Amount float64
}

package currency

func Debit(amount int32) int32 {
	if amount > 0 {
		return -amount
	}
	return amount
}

func Credit(amount int32) int32 {
	if amount < 0 {
		return -amount
	}
	return amount
}

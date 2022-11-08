package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return  3.213
	case balance >= 0 && balance < 1000:
		return 0.5
	case balance >= 1000 && balance < 5000:
		return 1.621
	case balance >= 5000:
		return 2.475
	default:
		return 0
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	rate := InterestRate(balance)

	return balance * float64(rate / 100.00)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	// @todo refactor to use Future Value Formula
	// https://www.forex.in.rs/solve-for-n-in-present-value-formula-and-future-value-formula/

	newBalance := balance
	year := 0

	for {
		if newBalance >= targetBalance {
			break
		}

		newBalance = AnnualBalanceUpdate(newBalance)
		year++
	}

	return year
}

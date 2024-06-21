package domain

type Balance struct {
	AverageDebit   string
	AverageCredit  string
	BalanceValue   string
	MonthlyBalance []MonthlyBalance
}

type MonthlyBalance struct {
	Month    string
	Quantity int
}

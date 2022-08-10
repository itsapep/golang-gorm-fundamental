package dto

type TotalProductEachCustomer struct {
	Id    string
	Name  string
	Total int
}

type CustomerWithNoProduct struct {
	Id   string
	Name string
}

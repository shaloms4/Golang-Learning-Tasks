package models

type Member struct {
	ID            int
	Name          string
	BorrowedBooks []Book
}

var Members = []Member{
	{ID: 1, Name: "Shalom Habtamu", BorrowedBooks: []Book{}},
	{ID: 2, Name: "Abebe Kebede", BorrowedBooks: []Book{}},
	{ID: 3, Name: "Mehari Tesfaye", BorrowedBooks: []Book{}},
}

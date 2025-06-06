package data

import "time"

type Book struct {
	ID       int64     `json:"id"`
	CreateAt time.Time `json:"-"`
	Title    string    `json:"title"`
	Year     int32     `json:"year,omitempty"`
	// Используйте тип Pages вместо int32. Обратите внимание, что omitempty
	// по-прежнему будет работать, и если поле Pages будет равно 0, то оно будет считаться пустым
	// и не будет включено в JSON, а метод MarshalJSON(), который мы создали,
	// вообще не будет вызван.
	Pages   Pages    `json:"pages,omitempty"`
	Genres  []string `json:"genres,omitempty"`
	Edition int32    `json:"edition"`
}

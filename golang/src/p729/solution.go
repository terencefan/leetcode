package main

type Book struct {
	start, end int
}

type MyCalendar struct {
	books []Book
}

func Constructor() MyCalendar {
	return MyCalendar{
		[]Book{},
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	book := Book{start, end}

	if len(this.books) == 0 {
		this.books = append(this.books, book)
		return true
	} else if start >= this.books[len(this.books)-1].end {
		this.books = append(this.books, book)
		return true
	}

	lo, hi := 0, len(this.books)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if this.books[mid].start <= start {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	if this.books[lo].start < end {
		return false
	} else if lo > 0 && this.books[lo-1].end > start {
		return false
	}

	this.books = append(this.books, book)
	for i := len(this.books) - 1; i > lo; i-- {
		this.books[i] = this.books[i-1]
	}
	this.books[lo] = book
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */

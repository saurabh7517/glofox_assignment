package booking

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"

	"os"
	"sync"
)

const filename string = "booking.json"

//DB - in-memory database of key - id of the booking and value - Book object
//RWMutex is used to lock map object while writing so that the read threads are blocked at that time.
type DB struct {
	mu         sync.RWMutex
	bookingmap map[int]Book
}

var db *DB

func init() {
	db = &DB{bookingmap: make(map[int]Book)}
	if err := loadData(db); err != nil {
		log.Print("Panic in loading data")
	}
	log.Print("Booking Data from file is loaded!!")
}

func loadData(db *DB) error {
	if _, err := os.Stat(filename); err != nil {
		return fmt.Errorf("cannot load booking file name %s", filename)
	}
	if file, err := ioutil.ReadFile(filename); err != nil {
		return fmt.Errorf("Unable read data from file %s", filename)
	} else {
		bookingList := make([]Book, 10)
		if err := json.Unmarshal([]byte(file), &bookingList); err != nil {
			return fmt.Errorf("Improper format of JSON structure")
		}

		for i := 0; i < len(bookingList); i++ {
			db.bookingmap[bookingList[i].BookingID] = bookingList[i]
		}
	}
	return nil
}

// Get all bookings
func getBookings() []Book {
	db.mu.RLock()
	defer db.mu.RUnlock()
	bookingList := make([]Book, 0, len(db.bookingmap))
	for _, v := range db.bookingmap {
		bookingList = append(bookingList, v)
	}

	return bookingList

}

// Delete an existing booking
func removeBooking(id int) (int, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.bookingmap[id]; ok {
		delete(db.bookingmap, id)
		return id, nil
	}
	return -1, fmt.Errorf("Book not found in database")

}

//AddBooking - Checking if the booking already exists
//Criteria - if start date and name are equal then the booking exists
func addbooking(b *bookingdto) (int, error) {

	db.mu.Lock()
	defer db.mu.Unlock()
	ID := -1
	if book, err := findBookingByNameandDate(b); err != nil {
		if ID = getMaxID(); ID == -1 {
			db.bookingmap[1] = Book{BookingID: 1, Name: b.Name, Date: b.Date}
			return 1, nil
		}

		db.bookingmap[ID+1] = Book{BookingID: ID + 1, Name: b.Name, Date: b.Date}
		return ID + 1, nil
	} else {
		fmt.Printf("booking already exists for id %v", book.BookingID)
		return book.BookingID, fmt.Errorf("booking already exists for id %v", book.BookingID)
	}

}

//updateBooking of the incoming PUT request
func updatebooking(id int, b *bookingdto) (int, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	//Checking if the booking alredy exists with the same Name and type
	newBook := &Book{}
	if oldbook, ok := db.bookingmap[id]; ok {
		newBook.BookingID = oldbook.BookingID
		newBook.Name = b.Name
		newBook.Date = b.Date
		db.bookingmap[id] = *newBook
		return newBook.BookingID, nil
	}
	return -1, fmt.Errorf("Book not found in database")

}

func findBookingByNameandDate(b *bookingdto) (Book, error) {

	flag := false
	existingBook := &Book{}
	if len(db.bookingmap) != 0 {
		for _, book := range db.bookingmap {
			if book.Name == b.Name && book.Date == b.Date {
				flag = true // a person can either update an existing booking or add a new booking with an new date
				existingBook = &book
				break
			}
		}
		if flag {
			return *existingBook, nil
		}
	}
	return *existingBook, fmt.Errorf("Book not available")
}

func getMaxID() int {

	if len(db.bookingmap) != 0 {
		keys := make([]int, 0, len(db.bookingmap))
		for k := range db.bookingmap {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		return keys[len(keys)-1]
	}

	return -1

}

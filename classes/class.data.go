package classes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
)

const filename string = "classes.json"

// A map of class objects are created with id as the key
// RWMutex is used to lock map object while writing so that the read threads are blocked at that time.

type classDB struct {
	mu       sync.RWMutex
	classMap map[int]Class
}

var cdb *classDB

//Class data is loaded in-memory before the server startup
func init() {
	cdb = &classDB{classMap: make(map[int]Class)}
	if err := loadData(cdb); err != nil {
		log.Print("Panic in loading data")
	}
	log.Print("Class Data from file is loaded!!")

}

func loadData(db *classDB) error {
	if _, err := os.Stat(filename); err != nil {
		return fmt.Errorf("cannot load classes file name %s", filename)
	}
	if file, err := ioutil.ReadFile(filename); err != nil {
		return fmt.Errorf("Unable read data from file %s", filename)
	} else {
		classList := make([]Class, 10)
		if err := json.Unmarshal([]byte(file), &classList); err != nil {
			return fmt.Errorf("Improper format of JSON structure")
		}

		for i := 0; i < len(classList); i++ {
			db.classMap[classList[i].ClassID] = classList[i]
		}
	}
	return nil

}

//Get all classes
func getClasses() []Class {
	cdb.mu.RLock()
	defer cdb.mu.RUnlock()
	classList := make([]Class, 0, len(cdb.classMap))
	for _, v := range cdb.classMap {
		classList = append(classList, v)
	}

	return classList

}

//Delete an existing class
func removeClass(id int) (int, error) {
	cdb.mu.Lock()
	defer cdb.mu.Unlock()

	if _, ok := cdb.classMap[id]; ok {
		delete(cdb.classMap, id)
		return id, nil
	}
	return -1, fmt.Errorf("Class not found in database")

}

//Update an existing class
func updateclass(id int, b *classdto) (int, error) {
	cdb.mu.Lock()
	defer cdb.mu.Unlock()
	//Checking if the class alredy exists with the same name and type
	// var class *Class
	newClass := &Class{}
	if oldClass, ok := cdb.classMap[id]; ok {
		newClass.ClassID = oldClass.ClassID
		newClass.Name = b.Name
		newClass.StartDate = b.StartDate
		newClass.EndDate = b.EndDate
		cdb.classMap[id] = *newClass

		return newClass.ClassID, nil
	}
	return -1, fmt.Errorf("Class not found in database")

}

//Add a new class
func addclass(cd *classdto) (int, error) {

	cdb.mu.Lock()
	defer cdb.mu.Unlock()
	ID := -1
	if class, err := findClassByNameandDate(cd); err != nil {
		if ID = getMaxID(); ID == -1 {
			cdb.classMap[1] = Class{ClassID: 1, Name: cd.Name, StartDate: cd.StartDate, EndDate: cd.EndDate}
			return 1, nil
		}

		cdb.classMap[ID+1] = Class{ClassID: ID + 1, Name: cd.Name, StartDate: cd.StartDate, EndDate: cd.EndDate}
		return ID + 1, nil
	} else {
		fmt.Printf("class already exists for id %v\n", class.ClassID)
		return class.ClassID, fmt.Errorf("class already exists for id %v", class.ClassID)
	}

}

//Find a class by name and date as specified by front end
func findClassByNameandDate(cd *classdto) (Class, error) {

	flag := false
	existingClass := &Class{}
	if len(cdb.classMap) != 0 {
		for _, class := range cdb.classMap {
			if class.Name == cd.Name && class.StartDate == cd.StartDate && class.EndDate == cd.EndDate {
				flag = true // a person can either update an existing class or add a new class with an new date
				existingClass = &class
				break
			}
		}
		if flag {
			return *existingClass, nil
		}
	}
	return *existingClass, fmt.Errorf("Class not available")
}

//Get maximum id existing in in-memory database and return an id in sequence
func getMaxID() int {

	if len(cdb.classMap) != 0 {
		keys := make([]int, 0, len(cdb.classMap))
		for k := range cdb.classMap {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		return keys[len(keys)-1]
	}

	return -1

}

package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

// Dbcon is connecton type
var Dbcon *bolt.DB
var tablename = []byte("todo")

// Task is used to initialize
type Task struct {
	ID   int
	Task string
}

func itob(i int) []byte {
	bt := make([]byte, 8)
	binary.BigEndian.PutUint64(bt, uint64(i))

	return bt
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

// CreateTask will create a new task
func CreateTask(task string) (int, error) {
	var ID int
	err := Dbcon.Update(func(t *bolt.Tx) error {
		buc := t.Bucket(tablename)
		ID64, _ := buc.NextSequence()
		ID := int(ID64)
		key := itob(ID)
		return buc.Put(key, []byte(task))
	})

	return ID, err
}

// DeleteTask will delete existing task
func DeleteTask(ID int) error {
	return Dbcon.Update(func(t *bolt.Tx) error {
		buc := t.Bucket(tablename)
		return buc.Delete(itob(ID))
	})
}

// AllTasks will show list of all task
func AllTasks() ([]Task, error) {
	var todolist []Task
	err := Dbcon.View(func(t *bolt.Tx) error {
		buc := t.Bucket(tablename)
		cur := buc.Cursor()
		for key, val := cur.First(); key != nil; key, val = cur.Next() {
			todolist = append(todolist, Task{
				ID:   btoi(key),
				Task: string(val),
			})
		}
		return nil
	})
	return todolist, err
}

// Init is used to make db connection
func Init(dbpath string) error {
	var err error
	Dbcon, err = bolt.Open(dbpath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return Dbcon.Update(func(t *bolt.Tx) error {
		_, err := t.CreateBucketIfNotExists(tablename)
		return err
	})

}

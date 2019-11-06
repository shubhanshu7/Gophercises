package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var DbCon *bolt.DB
var Tablename = []byte("todo")

type Task struct {
	Id   int
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

func CreateTask(task string) (int, error) {
	var id int
	err := DbCon.Update(func(t *bolt.Tx) error {
		buc := t.Bucket(Tablename)
		id64, _ := buc.NextSequence() //handle error
		id := int(id64)
		key := itob(id)
		return buc.Put(key, []byte(task))
	})

	return id, err
}

func DeleteTask(id int) error {
	return DbCon.Update(func(t *bolt.Tx) error {
		buc := t.Bucket(Tablename)
		return buc.Delete(itob(id))
	})
}

func AllTasks() ([]Task, error) {
	var todolist []Task
	err := DbCon.View(func(t *bolt.Tx) error {
		buc := t.Bucket(Tablename)
		cur := buc.Cursor()
		for key, val := cur.First(); key != nil; key, val = cur.Next() {
			todolist = append(todolist, Task{
				Id:   btoi(key),
				Task: string(val),
			})
		}
		return nil
	})
	return todolist, err
}
func Init(dbpath string) error {
	var err error
	DbCon, err = bolt.Open(dbpath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return DbCon.Update(func(t *bolt.Tx) error {
		_, err := t.CreateBucketIfNotExists(Tablename)
		return err
	})

}

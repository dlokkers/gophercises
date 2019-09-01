package tasklist

import (
	"fmt"
	"strconv"
	"path/filepath"

	"github.com/boltdb/bolt"
)

var todo = []byte("todo")
var done = []byte("done")
var db *bolt.DB

func Init(folder string) error {
	var err error
	db, err = bolt.Open( filepath.Join(folder, "tasks.db"), 0600, nil)
	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(todo)
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists(done)
		if err != nil {
			return err
		}
		return nil
	})
}

func Add(task string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(todo)

		i, err := b.NextSequence()
		if err != nil {
			return err
		}

		key := []byte(strconv.Itoa(int(i)))

		return b.Put(key, []byte(task))
	})
}

func List() error {
	return db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(todo)
		if b == nil {
			return fmt.Errorf("Bucket %q not found!", todo)
		}
		return b.ForEach(func(k, v []byte) error {
			fmt.Printf("%s - %s\n", k, v)
			return nil
		})
	})
}

func Do(key string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(todo)
		if b == nil {
			return fmt.Errorf("Bucket %q doesn't exist!", todo)
		}

		return b.Delete([]byte(key))
	})
}

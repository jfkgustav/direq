package store

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/jfkgustav/direq/model"
	bolt "go.etcd.io/bbolt"
	"log"
	"time"
)

type dbInfo struct {
	LastUpdate time.Time
}

func NewBoltDB(filepath string) *bolt.DB {
	log.Printf("Opening bolt DB %s ...\n", filepath)
	db, err := bolt.Open(filepath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		log.Printf("with buckets:\n")
		_, err := tx.CreateBucketIfNotExists([]byte("repertoire"))
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\t- repertoire\n")
		_, err = tx.CreateBucketIfNotExists([]byte("sessions"))
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\t- sessions\n")
		_, err = tx.CreateBucketIfNotExists([]byte("dbinfo"))
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\t- dbinfo\n")
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("... successfully!")

	return db
}

func UpdateRepertoire(db *bolt.DB, songs []model.Song) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("repertoire"))
		for _, song := range songs {
			data, err := json.Marshal(song)
			if err != nil {
				return fmt.Errorf("update repertoire failed marshalling song (id: %d) data: %v", song.ID, err)
			}
			err = b.Put(itob(song.ID), data)
			if err != nil {
				return fmt.Errorf("update repertoire failed db update of song (id: %d) data: %v", song.ID, err)
			}
		}
		data, _ := json.Marshal(dbInfo{LastUpdate: time.Now()})
		b = tx.Bucket([]byte("dbinfo"))
		return b.Put([]byte("dbinfo"), data)

	})
	return err
}

func GetSongs(db *bolt.DB) ([]model.Song, error) {
	var songs []model.Song
	err := db.View(func(tx *bolt.Tx) error {
		repertoire := tx.Bucket([]byte("repertoire"))
		c := repertoire.Cursor()
		var song model.Song

		for k, v := c.First(); k != nil; k, v = c.Next() {
			err := json.Unmarshal(v, &song)
			if err != nil {
				return fmt.Errorf("get songs failed unmarhalling data %v", err)
			}
			songs = append(songs, song)
		}

		return nil
	})
	return songs, err
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

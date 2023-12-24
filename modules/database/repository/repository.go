package repository

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/boltdbgui/modules/database/model"
	bolt "go.etcd.io/bbolt"
	"golang.org/x/xerrors"
)

var (
	db *bolt.DB
)

type Config struct {
}

func Init(dbPath string) (err error) {
	db, err = bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return xerrors.Errorf("failed to open db: %w", err)
	}
	return nil
}
func Close() error {
	// Skip closing the database if the connection is not established.
	if db == nil {
		return nil
	}
	if err := db.Close(); err != nil {
		return xerrors.Errorf("failed to close DB: %w", err)
	}
	return nil
}

func ListElement(input model.ListElemReqBody) (elem model.ListedElem, err error) {
	var resultFullSet []model.Result
	err = db.View(func(tx *bolt.Tx) error {
		var rootBkt *bolt.Bucket
		if len(input.LevelStack) > 0 {
			rootBkt = tx.Bucket([]byte(input.LevelStack[0]))
			if rootBkt == nil {
				return xerrors.New(fmt.Sprintf("No Root Bucket found by the name : %s", input.LevelStack[0]))
			}
			for i, val := range input.LevelStack[1:] {
				rootBkt = rootBkt.Bucket([]byte(val))
				if rootBkt == nil {
					return xerrors.New(fmt.Sprintf("No Bucket found by the name : %s under the level : %s", val, input.LevelStack[:i]))
				}
			}
			_ = rootBkt.ForEach(func(k []byte, v []byte) error {
				if input.SearchKey != "" && !strings.Contains(string(k), input.SearchKey) {
					return nil
				}
				if v == nil {
					elBkt := rootBkt.Bucket(k)
					stats := elBkt.Stats()
					resultFullSet = append(resultFullSet,
						model.Result{
							Name:          string(k),
							IsBucket:      true,
							Value:         "",
							NoOfChildBkts: stats.InlineBucketN,
							NoOfPairs:     stats.KeyN,
						})
				} else {
					resultFullSet = append(resultFullSet,
						model.Result{
							Name:     string(k),
							IsBucket: false,
							Value:    string(v),
						})
				}
				return nil
			})
		} else {
			_ = tx.ForEach(func(name []byte, b *bolt.Bucket) error {
				if input.SearchKey != "" && !strings.Contains(string(name), input.SearchKey) {
					return nil
				}
				stats := b.Stats()
				resultFullSet = append(resultFullSet,
					model.Result{
						Name:          string(name),
						IsBucket:      true,
						Value:         "",
						NoOfChildBkts: stats.InlineBucketN,
						NoOfPairs:     stats.KeyN,
					})

				return nil
			})
		}

		elem.LevelStack = input.LevelStack
		elem.SearchKey = input.SearchKey
		elem.Results = resultFullSet
		return nil
	})
	if err != nil {
		return model.ListedElem{}, err
	}
	return elem, nil
}

func AddBuckets(input model.BucketsToAdd) (err error) {
	err = db.Update(func(tx *bolt.Tx) error {
		var rootBkt *bolt.Bucket
		if len(input.LevelStack) > 0 {
			rootBkt = tx.Bucket([]byte(input.LevelStack[0]))
			if rootBkt == nil {
				return xerrors.New(fmt.Sprintf("No Root Bucket found by the name : %s", input.LevelStack[0]))
			}
			for i, val := range input.LevelStack[1:] {
				rootBkt = rootBkt.Bucket([]byte(val))
				if rootBkt == nil {
					return xerrors.New(fmt.Sprintf("No Bucket found by the name : %s under the level : %s", val, input.LevelStack[:i]))
				}
			}
			for _, bkt := range input.Buckets {
				_, err = rootBkt.CreateBucket([]byte(bkt))
				if err != nil {
					return errors.Wrapf(err, "Unable to create bucket %s under stack %s", bkt, input.LevelStack)
				}
			}
		} else {
			for _, bkt := range input.Buckets {
				_, err = tx.CreateBucket([]byte(bkt))
				if err != nil {
					return errors.Wrapf(err, "Unable to create bucket %s under root", bkt)
				}
			}
		}
		return nil
	})
	return err
}

func AddPairs(input model.PairsToAdd) (err error) {
	err = db.Update(func(tx *bolt.Tx) error {
		var rootBkt *bolt.Bucket
		if len(input.LevelStack) == 0 {
			return errors.New("Cannot create key/value pairs without parent bucket, levelstack missing")
		}
		rootBkt = tx.Bucket([]byte(input.LevelStack[0]))
		if rootBkt == nil {
			return xerrors.New(fmt.Sprintf("No Root Bucket found by the name : %s", input.LevelStack[0]))
		}
		for i, val := range input.LevelStack[1:] {
			rootBkt = rootBkt.Bucket([]byte(val))
			if rootBkt == nil {
				return xerrors.New(fmt.Sprintf("No Bucket found by the name : %s under the level : %s", val, input.LevelStack[:i]))
			}
		}
		for _, pair := range input.Pairs {
			val, err1 := json.Marshal(pair.Value)
			if err1 != nil {
				return errors.Wrapf(err1, "Unable to marshal the value %v", pair.Value)
			}
			err = rootBkt.Put([]byte(pair.Key), val)
			if err != nil {
				return errors.Wrapf(err, "Unable to create pair %s under stack %s", pair, input.LevelStack)
			}
		}
		return nil
	})
	return err
}

func DeleteElement(input model.ItemToDelete) (err error) {
	err = db.Update(func(tx *bolt.Tx) error {
		var rootBkt *bolt.Bucket
		if len(input.LevelStack) > 0 {
			rootBkt = tx.Bucket([]byte(input.LevelStack[0]))
			if rootBkt == nil {
				return xerrors.New(fmt.Sprintf("No Root Bucket found by the name : %s", input.LevelStack[0]))
			}
			for i, val := range input.LevelStack[1:] {
				rootBkt = rootBkt.Bucket([]byte(val))
				if rootBkt == nil {
					return xerrors.New(fmt.Sprintf("No Bucket found by the name : %s under the level : %s", val, input.LevelStack[:i]))
				}
			}
			if bkt := rootBkt.Bucket([]byte(input.Key)); bkt != nil {
				err = rootBkt.DeleteBucket([]byte(input.Key))
				if err != nil {
					return errors.Wrapf(err, "Unable to delete bucket %s under stack %s", input.Key, input.LevelStack)
				}
			} else {
				err = rootBkt.Delete([]byte(input.Key))
				if err != nil {
					return errors.Wrapf(err, "Unable to delete key %s under stack %s", input.Key, input.LevelStack)
				}
			}
		} else {
			err = tx.DeleteBucket([]byte(input.Key))
			if err != nil {
				return errors.Wrapf(err, "Unable to delete bucket %s under root", input.Key)
			}
		}
		return nil
	})
	return err
}

func RenameElement(input model.ItemToRename) (err error) {
	err = db.Update(func(tx *bolt.Tx) error {
		var rootBkt *bolt.Bucket
		if len(input.LevelStack) > 0 {
			rootBkt = tx.Bucket([]byte(input.LevelStack[0]))
			if rootBkt == nil {
				return xerrors.New(fmt.Sprintf("No Root Bucket found by the name : %s", input.LevelStack[0]))
			}
			for i, val := range input.LevelStack[1:] {
				rootBkt = rootBkt.Bucket([]byte(val))
				if rootBkt == nil {
					return xerrors.New(fmt.Sprintf("No Bucket found by the name : %s under the level : %s", val, input.LevelStack[:i]))
				}
			}

			val := rootBkt.Get([]byte(input.Key))
			if val != nil {
				err = rootBkt.Put([]byte(input.NewKey), val)
				if err != nil {
					return errors.Wrapf(err, "Unable to put key %s under stack %s", input.NewKey, input.LevelStack)
				}
				err = rootBkt.Delete([]byte(input.Key))
				if err != nil {
					return errors.Wrapf(err, "Unable to delete key %s under stack %s", input.Key, input.LevelStack)
				}
			} else {
				return errors.New("No Key found to be replaced")
			}
		} else {
			return errors.New("Buckets cannot be renamed, please provide level stack")
		}
		return nil
	})
	return err
}

func UpdatePairValue(input model.ItemToUpdate) (err error) {
	err = db.Update(func(tx *bolt.Tx) error {
		var rootBkt *bolt.Bucket
		if len(input.LevelStack) > 0 {
			rootBkt = tx.Bucket([]byte(input.LevelStack[0]))
			if rootBkt == nil {
				return xerrors.New(fmt.Sprintf("No Root Bucket found by the name : %s", input.LevelStack[0]))
			}
			for i, val := range input.LevelStack[1:] {
				rootBkt = rootBkt.Bucket([]byte(val))
				if rootBkt == nil {
					return xerrors.New(fmt.Sprintf("No Bucket found by the name : %s under the level : %s", val, input.LevelStack[:i]))
				}
			}
			if val := rootBkt.Get([]byte(input.Key)); val != nil {
				val1, err1 := json.Marshal(input.NewValue)
				if err1 != nil {
					return errors.Wrapf(err1, "Unable to marshal the value %v", input.NewValue)
				}
				err = rootBkt.Put([]byte(input.Key), val1)
				if err != nil {
					return errors.Wrap(err, "Unable to put new value")
				}
			} else {
				return errors.New("Given key not found")
			}
		} else {
			return errors.New("Please provide level stack")
		}
		return nil
	})
	return err
}

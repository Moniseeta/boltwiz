package repository

import (
	"fmt"

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
		//siblingBkts := make(map[string]*bolt.Bucket)
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
				if input.SearchKey != "" && string(k) != input.SearchKey {
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
				if input.SearchKey != "" && string(name) != input.SearchKey {
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
			}) // TODO revisit
		}
		//if input.SearchKey != "" {
		//	if rootBkt != nil {
		//		if elBkt := rootBkt.Bucket([]byte(input.SearchKey)); elBkt != nil {
		//			elem.LevelStack = append(input.LevelStack, input.SearchKey)
		//			elem.IsBucket = true
		//			stats := elBkt.Stats()
		//			elem.NoOfChildBkts = stats.InlineBucketN
		//			elem.NoOfPairs = stats.KeyN
		//			// TODO need to write code for details of child buckets
		//			/*elBkt.ForEachBucket(func(bkt []byte) error {
		//				elem.ChildBkts = append(elem.ChildBkts, string(bkt))
		//			})*/
		//		} else if elVal := rootBkt.Get([]byte(input.Key)); elVal != nil {
		//			elem.LevelStack = append(input.LevelStack, input.Key)
		//			elem.IsBucket = false
		//			elem.Value = string(elVal)
		//		} else {
		//			return xerrors.New(fmt.Sprintf("No Bucket or Key found by the name : %s under the level : %s", input.Key, input.LevelStack))
		//		}
		//	} else {
		//		elem.LevelStack = []string{siblingBkts[0]}
		//		elem.IsBucket = true
		//		elem.SiblingBkts = siblingBkts[1:]
		//	}
		//
		//} else {
		//	elem.LevelStack = []string{siblingBkts[0]}
		//	elem.IsBucket = true
		//	elem.SiblingBkts = siblingBkts[1:]
		//}
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

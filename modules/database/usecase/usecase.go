package usecase

import (
	"github.com/boltdbgui/modules/database/model"
	"github.com/boltdbgui/modules/database/repository"
)

func ListElement(input model.ListElemReqBody) (elem model.ListedElem, err error) {
	return repository.ListElement(input)
}
func AddBuckets(input model.BucketsToAdd) (err error) {
	return repository.AddBuckets(input)
}
func AddPairs(input model.PairsToAdd) (err error) {
	return repository.AddPairs(input)
}
func DeleteElement(input model.ItemToDelete) (err error) {
	return repository.DeleteElement(input)
}
func RenameElement(input model.ItemToRename) (err error) {
	return repository.RenameElement(input)
}
func UpdatePairValue(input model.ItemToUpdate) (err error) {
	return repository.UpdatePairValue(input)
}

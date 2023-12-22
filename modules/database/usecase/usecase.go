package usecase

import (
	"github.com/boltdbgui/modules/database/model"
	"github.com/boltdbgui/modules/database/repository"
)

func ListElement(input model.ListElemReqBody) (elem model.ListedElem) {

	return repository.ListElement(input)
}

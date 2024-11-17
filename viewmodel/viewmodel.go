package viewmodel

import (
	"time"

	"leanmngconcept/model"
	"leanmngconcept/repository"
)

type ActionViewModel struct {
	Actions    []*model.Action
	Repository *repository.ActionRepository
}

type NextAction struct {
	Index  int
	Action *model.Action
}

func NewActionViewModel(repo *repository.ActionRepository) *ActionViewModel {
	actions, _ := repo.LoadActions()
	return &ActionViewModel{
		Actions:    actions,
		Repository: repo,
	}
}

func (vm *ActionViewModel) AddAction(name string) {
	action := model.NewAction(name)
	vm.Actions = append(vm.Actions, action)
	vm.SaveActions()
}

func (vm *ActionViewModel) MarkActionDone(index int) {
	if index >= 0 && index < len(vm.Actions) {
		vm.Actions[index].MarkDone()
		vm.SaveActions()
	}
}

func (vm *ActionViewModel) GetNextReviewActions() []NextAction {
	NextReviewActions := []NextAction{}

	currentDay := time.Now()
	for index, action := range vm.Actions {
		if action.NextReview.Before(currentDay) {
			NextReviewActions = append(
				NextReviewActions,
				NextAction{
					Index:  index,
					Action: action,
				},
			)
		}
	}

	return NextReviewActions
}

func (vm *ActionViewModel) SaveActions() {
	vm.Repository.SaveActions(vm.Actions)
}

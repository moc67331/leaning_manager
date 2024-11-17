package main

import (
	"leanmngconcept/repository"
	"leanmngconcept/viewmodel"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func loadResource(name string) []byte {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	data := make([]byte, info.Size())
	_, err = file.Read(data)
	if err != nil {
		panic(err)
	}

	return data
}

func main() {
	repo := repository.NewActionRepository("actionList.save")
	viewModel := viewmodel.NewActionViewModel(repo)

	myApp := app.New()
	myWindow := myApp.NewWindow("Review Manager")

	logo := theme.NewThemedResource(fyne.NewStaticResource("logo.png", loadResource("logo.png")))
	myApp.SetIcon(logo)

	actionListContainer := container.NewVBox()

	updateActionList := func() {}
	resetActionList := func() {
		for _, action := range viewModel.GetNextReviewActions() {
			actionListContainer.Add(container.NewHBox(
				widget.NewLabel("・"+action.Action.Name),
				widget.NewButton("Done", func() {
					viewModel.MarkActionDone(action.Index)
					updateActionList()
				}),
			))
		}
	}
	resetActionList()

	updateActionList = func() {
		actionListContainer.RemoveAll()
		resetActionList()
		actionListContainer.Refresh()
	}

	newActionEntry := widget.NewEntry()
	newActionButton := widget.NewButton("Add", func() {
		viewModel.AddAction(newActionEntry.Text)
		newActionEntry.SetText("")
		updateActionList()
	})

	updateButton := widget.NewButton("Update", func() {
		updateActionList()
	})

	actionListFrame := container.NewVBox(
		widget.NewLabel("Actions"),
		actionListContainer,
	)
	controlFrame := container.NewVBox(
		newActionEntry,
		newActionButton,
		updateButton,
	)

	mainFrame := container.NewGridWithRows(2,
		container.NewScroll(actionListFrame),
		controlFrame,
	)

	myWindow.Resize(fyne.NewSize(300, 400))
	myWindow.SetContent(mainFrame)
	myWindow.ShowAndRun()
}

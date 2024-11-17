package repository

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"leanmngconcept/model"
)

type ActionRepository struct {
	FilePath string
}

func NewActionRepository(filePath string) *ActionRepository {
	return &ActionRepository{FilePath: filePath}
}

func (r *ActionRepository) LoadActions() ([]*model.Action, error) {
	file, err := os.Open(r.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var actions []*model.Action
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "\x1f")
		if len(parts) < 3 {
			continue
		}

		dateInt, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			continue
		}

		reviewPeriod, err := strconv.Atoi(parts[2])
		if err != nil {
			continue
		}

		actions = append(actions, &model.Action{
			Name:         parts[0],
			NextReview:   time.Unix(dateInt, 0),
			ReviewPeriod: reviewPeriod,
		})
	}

	return actions, scanner.Err()
}

func (r *ActionRepository) SaveActions(actions []*model.Action) error {
	file, err := os.Create(r.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, action := range actions {
		_, err := fmt.Fprintf(writer, "%s\x1f%d\x1f%d\n", action.Name, action.NextReview.Unix(), action.ReviewPeriod)
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

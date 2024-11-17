package model

import "time"

type Action struct {
	Name         string    // 項目名
	NextReview   time.Time // 次回復習日
	ReviewPeriod int       // 復習間隔（日数）
}

// 新規アクションの生成
func NewAction(name string) *Action {
	return &Action{
		Name:         name,
		NextReview:   time.Now(),
		ReviewPeriod: 1, // 初期復習期間は1日
	}
}

// 復習が完了したときの処理
func (a *Action) MarkDone() {
	a.ReviewPeriod *= 2
	a.NextReview = time.Now().AddDate(0, 0, a.ReviewPeriod)
}

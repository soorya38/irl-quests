package quest

import "qgend/entity"

type QuestUsecase interface {
	CreateQuest(title, desc string, level int, tags []string) error
	FetchQuests(page, limit, level int, tags []string) ([]*entity.Quest, error)
	UpdateQuest(id int64, title, desc string, level int, tags []string) error
	DeleteQuest(id int64) error
}

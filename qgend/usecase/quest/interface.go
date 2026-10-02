package quest

import "qgend/entity"

// QuestRepository defines the quest repository functionalities.
type QuestRepository interface {
	CreateQuest(quest *entity.Quest) error
	FetchQuests(page, limit, level int, tags []string) ([]*entity.Quest, error)
	UpdateQuest(quest *entity.Quest) error
	DeleteQuest(id int64) error
}

// QuestUsecase defines the quest functionalities.
type QuestUsecase interface {
	CreateQuest(quest *entity.Quest) error
	FetchQuests(page, limit, level int, tags []string) ([]*entity.Quest, error)
	UpdateQuest(quest *entity.Quest) error
	DeleteQuest(id int64) error
}

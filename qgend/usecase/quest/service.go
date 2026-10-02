package quest

import "qgend/entity"

// QuestService implements the QuestUsecase.
type QuestService struct{}

// NewQuestService creates a QuestService instance.
func NewQuestService() *QuestService {
	return &QuestService{}
}

// CreateQuest validates if the user can create a quest and creates a quest.
func (q *QuestService) CreateQuest(title, desc string, level int, tags []string) error {
	return nil
}

// FetchQuests fetches the quests based on the provided level, tags.
func (q *QuestService) FetchQuests(page, limit, level int, tags []string) ([]*entity.Quest, error) {
	return nil, nil
}

// UpdateQuest updates an existing quest, only the person who created the quest can update it.
func (q *QuestService) UpdateQuest(id int64, title, desc string, level int, tags []string) error {
	return nil
}

// DeleteQuest deletes the quest based on the provided id. Only the person who created
// the quest must be able to delete the quest.
func (q *QuestService) DeleteQuest(id int64) error {
	return nil
}

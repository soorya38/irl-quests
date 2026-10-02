package repository

import "qgend/entity"

// QuestRepository implements the QuestRepository interface.
type QuestRepository struct{}

// CreateQuest inserts a record in the quest table.
func (r *QuestRepository) CreateQuest(quest *entity.Quest) error {
	return nil
}

// FetchQuests fetches the records from quest table where the provided filters match.
func (r *QuestRepository) FetchQuests(page, limit, level int, tags []string) ([]*entity.Quest, error) {
	return nil, nil
}

// UpdateQuest updates the quest based on the provided quest id.
func (r *QuestRepository) UpdateQuest(quest *entity.Quest) error {
	return nil
}

// DeleteQuest deletes the quest based on the provided quest id.
func (r *QuestRepository) DeleteQuest(id int64) error {
	return nil
}

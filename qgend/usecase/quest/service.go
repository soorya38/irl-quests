package quest

import (
	"fmt"
	"qgend/entity"
)

// QuestService implements the QuestUsecase.
type QuestService struct {
	repo QuestRepository
}

// NewQuestService creates a QuestService instance.
func NewQuestService(repo QuestRepository) *QuestService {
	return &QuestService{
		repo: repo,
	}
}

// CreateQuest validates if the user can create a quest and creates a quest.
func (s *QuestService) CreateQuest(quest *entity.Quest) error {
	if err := s.repo.CreateQuest(quest); err != nil {
		return fmt.Errorf("unable to create quest: %v", err)
	}
	return nil
}

// FetchQuests fetches the quests based on the provided level, tags.
func (s *QuestService) FetchQuests(page, limit, level int, tags []string) ([]*entity.Quest, error) {
	quests, err := s.repo.FetchQuests(page, limit, level, tags)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch quest: %v", err)
	}
	return quests, nil
}

// UpdateQuest updates an existing quest, only the person who created the quest can update it.
func (s *QuestService) UpdateQuest(quest *entity.Quest) error {
	if err := s.repo.UpdateQuest(quest); err != nil {
		return fmt.Errorf("unable to update quest: %v", err)
	}
	return nil
}

// DeleteQuest deletes the quest based on the provided id. Only the person who created
// the quest must be able to delete the quest.
func (s *QuestService) DeleteQuest(id int64) error {
	if err := s.repo.DeleteQuest(id); err != nil {
		return fmt.Errorf("unable to delete quest: %v", err)
	}
	return nil
}

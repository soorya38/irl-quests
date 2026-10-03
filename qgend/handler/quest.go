package handler

import "net/http"

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// Verifies if the person is capable of creating a quest.
// Creates a quest.
func CreateQuest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

// Lists the quests based on the provided filters.
func GetQuests(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Verifies if the person is capable of updating a quest.
// Updates an existing quest.
func UpdateQuest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Verifies if the person is capable of deleting a quest.
// Deletes an existing quest.
func DeleteQuest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

package service

import (
	agoratix "github.com/bellatrijuliana/agoratix-app/features/event"
)

type service struct {
	repo agoratix.RepositoryInterface
}

func NewService(repo agoratix.RepositoryInterface) agoratix.ServiceInterface {
	return &service{repo: repo}
}

func (s *service) GetEventList() ([]agoratix.Event, error) {
	return s.repo.GetEventList()
}

// PERBAIKAN: Hapus 'input' karena fungsi ini hanya butuh 'id'
func (s *service) GetEventByID(id int) (agoratix.Event, error) {
	return s.repo.GetEventByID(id)
}

func (s *service) InsertEvent(input agoratix.Event) (agoratix.Event, error) {
	// Di sini Anda bisa menambahkan validasi atau logika bisnis sebelum menyimpan
	return s.repo.InsertEvent(input)
}

func (s *service) UpdateEvent(id int, input agoratix.Event) (agoratix.Event, error) {
	return s.repo.UpdateEvent(id, input)
}

func (s *service) DeleteEvent(id int) error {
	return s.repo.DeleteEvent(id)
}

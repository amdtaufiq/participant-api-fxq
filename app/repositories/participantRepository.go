package repositories

import (
	"fmt"
	"participant-api/app/entities"

	"gorm.io/gorm"
)

type IParticipantRepository interface {
	GetAll() ([]entities.Participant, error)
	GetByID(ID string) (entities.Participant, error)
	Create(participant entities.Participant) (bool, error)
	Update(participant entities.Participant) (bool, error)
	Delete(ID string) (bool, error)
}

type participantRepository struct {
	db *gorm.DB
}

func ParticipantRepository(db *gorm.DB) *participantRepository {
	return &participantRepository{db}
}

func (r *participantRepository) GetAll() ([]entities.Participant, error) {
	var participants []entities.Participant

	err := r.db.Find(&participants).Error

	if err != nil {
		return participants, err
	}
	return participants, nil
}

func (r *participantRepository) GetByID(ID string) (entities.Participant, error) {
	var participant entities.Participant

	err := r.db.Where("id = ?", ID).Find(&participant).Error

	if err != nil {
		return participant, err
	}
	return participant, nil
}

func (r *participantRepository) Create(participant entities.Participant) (bool, error) {
	fmt.Print("create participant")
	err := r.db.Create(&participant).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *participantRepository) Update(participant entities.Participant) (bool, error) {
	err := r.db.Save(&participant).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *participantRepository) Delete(ID string) (bool, error) {
	err := r.db.Where("id = ?", ID).Delete(&entities.Participant{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

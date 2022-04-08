package services

import (
	"errors"
	"participant-api/app/entities"
	"participant-api/app/inputs"
	"participant-api/app/repositories"

	"github.com/google/uuid"
)

type IParticipantService interface {
	GetAll() ([]entities.Participant, error)
	GetByID(request inputs.GetIDParticipantInput) (entities.Participant, error)
	Create(request inputs.CreateParticipantInput) (bool, error)
	Update(requestID inputs.GetIDParticipantInput, requestData inputs.UpdateParticipantInput) (bool, error)
	Delete(request inputs.GetIDParticipantInput) (bool, error)
}

type participantService struct {
	participantRepository repositories.IParticipantRepository
}

func ParticipantService(participantRepository repositories.IParticipantRepository) *participantService {
	return &participantService{participantRepository}
}

func (s *participantService) GetAll() ([]entities.Participant, error) {
	participants, err := s.participantRepository.GetAll()
	if err != nil {
		return participants, err
	}
	return participants, nil
}

func (s *participantService) GetByID(request inputs.GetIDParticipantInput) (entities.Participant, error) {
	participant := entities.Participant{}
	participant, err := s.participantRepository.GetByID(request.ID)
	if err != nil {
		return participant, err
	}
	return participant, nil
}

func (s *participantService) Create(request inputs.CreateParticipantInput) (bool, error) {
	participantFormat := inputs.FormatCreateParticipant(request)

	result, err := s.participantRepository.Create(participantFormat)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *participantService) Update(requestID inputs.GetIDParticipantInput, requestData inputs.UpdateParticipantInput) (bool, error) {
	participant, err := s.participantRepository.GetByID(requestID.ID)
	if err != nil {
		return false, err
	}

	if participant.ID != requestData.ID {
		return false, errors.New("participant data not match")
	}

	participantFormat := inputs.FormatUpdateParticipant(participant, requestData)

	result, err := s.participantRepository.Update(participantFormat)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *participantService) Delete(inputs inputs.GetIDParticipantInput) (bool, error) {
	data, err := s.participantRepository.GetByID(inputs.ID)
	if err != nil || data.ID == uuid.Nil {
		return false, errors.New("participant data not available")
	}

	result, err := s.participantRepository.Delete(inputs.ID)
	if err != nil {
		return result, err
	}
	return result, nil
}

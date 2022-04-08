package inputs

import (
	"participant-api/app/entities"

	"github.com/google/uuid"
)

type GetIDParticipantInput struct {
	ID string `uri:"id" binding:"required"`
}

type CreateParticipantInput struct {
	FullName     string `json:"full_name" binding:"required"`
	BusinessName string `json:"business_name" binding:"required"`
	Email        string `json:"email" binding:"required"`
	PhoneNumber  string `json:"phone_number" binding:"required"`
}

type UpdateParticipantInput struct {
	ID           uuid.UUID `json:"id" binding:"required"`
	FullName     string    `json:"full_name" binding:"required"`
	BusinessName string    `json:"business_name" binding:"required"`
	Email        string    `json:"email" binding:"required"`
	PhoneNumber  string    `json:"phone_number" binding:"required"`
}

func FormatCreateParticipant(input CreateParticipantInput) entities.Participant {
	data := entities.Participant{}
	data.FullName = input.FullName
	data.BusinessName = input.BusinessName
	data.Email = input.Email
	data.PhoneNumber = input.PhoneNumber
	data.IsPrintCertificate = false
	data.IsPrintNameTag = false
	return data
}

func FormatUpdateParticipant(data entities.Participant, input UpdateParticipantInput) entities.Participant {
	data.FullName = input.FullName
	data.BusinessName = input.BusinessName
	data.Email = input.Email
	data.PhoneNumber = input.PhoneNumber
	return data
}

package formatters

import (
	"participant-api/app/entities"
	"time"

	"github.com/google/uuid"
)

type ParticipantFormatter struct {
	ID                 uuid.UUID `json:"id"`
	FullName           string    `json:"full_name"`
	BusinessName       string    `json:"business_name"`
	Email              string    `json:"email"`
	PhoneNumber        string    `json:"phone_number"`
	IsPrintCertificate bool      `json:"is_print_certificate"`
	IsPrintNameTag     bool      `json:"is_print_name_tag"`
	CreatedAt          time.Time `json:"created_at"`
}

func FormatParticipant(data entities.Participant) ParticipantFormatter {
	dataFormatter := ParticipantFormatter{}
	dataFormatter.ID = data.ID
	dataFormatter.FullName = data.FullName
	dataFormatter.BusinessName = data.BusinessName
	dataFormatter.Email = data.Email
	dataFormatter.PhoneNumber = data.PhoneNumber
	dataFormatter.IsPrintCertificate = data.IsPrintCertificate
	dataFormatter.IsPrintNameTag = data.IsPrintNameTag
	dataFormatter.CreatedAt = data.CreatedAt

	return dataFormatter
}

func FormatParticipants(datas []entities.Participant) []ParticipantFormatter {
	if len(datas) == 0 {
		return []ParticipantFormatter{}
	}

	datasFormattter := []ParticipantFormatter{}

	for _, data := range datas {
		dataFormatter := FormatParticipant(data)
		datasFormattter = append(datasFormattter, dataFormatter)
	}

	return datasFormattter
}

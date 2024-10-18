package card

import domaincard "lesson-1/internal/domain/card"

type CreateResponse struct {
}

type DefaultResponse struct {
	domaincard.Card
}

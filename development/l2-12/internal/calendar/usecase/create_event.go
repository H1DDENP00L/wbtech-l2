package usecase

import (
	"context"

	"github.com/H1DDENP00L/wbtech-l2/development/l2-12/internal/calendar/domain"
)

type CreateEventUseCase struct {
	eventRepository domain.Repository
}

func NewCreateEventUseCase(
	eventRepository domain.Repository,
) *CreateEventUseCase {
	return &CreateEventUseCase{
		eventRepository: eventRepository,
	}
}

func (uc *CreateEventUseCase) Execute(ctx context.Context, event domain.Event) (int, error) {
	return uc.eventRepository.CreateEvent(ctx, event)
}

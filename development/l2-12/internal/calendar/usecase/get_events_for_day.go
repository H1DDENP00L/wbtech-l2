package usecase

import (
	"context"
	"time"

	"github.com/H1DDENP00L/wbtech-l2/development/l2-12/internal/calendar/domain"
)

type GetEventsForDayUseCase struct {
	eventRepository domain.Repository
}

func NewGetEventsForDayUseCase(
	eventRepository domain.Repository,
) *GetEventsForDayUseCase {
	return &GetEventsForDayUseCase{
		eventRepository: eventRepository,
	}
}

func (uc *GetEventsForDayUseCase) Execute(ctx context.Context, date time.Time) ([]domain.Event, error) {
	return uc.eventRepository.GetEventsForDay(ctx, date)
}

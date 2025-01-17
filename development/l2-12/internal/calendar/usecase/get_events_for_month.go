package usecase

import (
	"context"
	"time"

	"github.com/H1DDENP00L/wbtech-l2/development/l2-12/internal/calendar/domain"
)

type GetEventsForMonthUseCase struct {
	eventRepository domain.Repository
}

func NewGetEventsForMonthUseCase(
	eventRepository domain.Repository,
) *GetEventsForMonthUseCase {
	return &GetEventsForMonthUseCase{
		eventRepository: eventRepository,
	}
}

func (uc *GetEventsForMonthUseCase) Execute(ctx context.Context, date time.Time) ([]domain.Event, error) {
	return uc.eventRepository.GetEventsForMonth(ctx, date)
}

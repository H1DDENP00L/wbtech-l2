package usecase

import (
	"context"
	"time"

	"github.com/H1DDENP00L/wbtech-l2/development/l2-12/internal/calendar/domain"
)

type GetEventsForWeekUseCase struct {
	eventRepository domain.Repository
}

func NewGetEventsForWeekUseCase(
	eventRepository domain.Repository,
) *GetEventsForWeekUseCase {
	return &GetEventsForWeekUseCase{
		eventRepository: eventRepository,
	}
}

func (uc *GetEventsForWeekUseCase) Execute(ctx context.Context, date time.Time) ([]domain.Event, error) {
	return uc.eventRepository.GetEventsForWeek(ctx, date)
}

package builder

import (
	"context"

	"github.com/H1DDENP00L/wbtech-l2/development/l2-12/internal/calendar/adapters"
	"github.com/H1DDENP00L/wbtech-l2/development/l2-12/internal/calendar/usecase"
)

type Application struct {
	CreateEvent       *usecase.CreateEventUseCase
	UpdateEvent       *usecase.UpdateEventUseCase
	DeleteEvent       *usecase.DeleteEventUseCase
	GetEventByID      *usecase.GetEventByIDUseCase
	GetEventsForDay   *usecase.GetEventsForDayUseCase
	GetEventsForWeek  *usecase.GetEventsForWeekUseCase
	GetEventsForMonth *usecase.GetEventsForMonthUseCase
}

func NewApplication(ctx context.Context) *Application {
	eventRepository := adapters.NewCacheEventRepository(200)

	return &Application{
		CreateEvent:       usecase.NewCreateEventUseCase(eventRepository),
		UpdateEvent:       usecase.NewUpdateEventUseCase(eventRepository),
		DeleteEvent:       usecase.NewDeleteEventUseCase(eventRepository),
		GetEventByID:      usecase.NewGetEventByIDUseCase(eventRepository),
		GetEventsForDay:   usecase.NewGetEventsForDayUseCase(eventRepository),
		GetEventsForWeek:  usecase.NewGetEventsForWeekUseCase(eventRepository),
		GetEventsForMonth: usecase.NewGetEventsForMonthUseCase(eventRepository),
	}
}

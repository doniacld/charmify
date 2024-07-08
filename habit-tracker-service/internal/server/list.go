package server

import (
	"context"

	"github.com/doniacld/charmify/habit-tracker-service/api"
	"github.com/doniacld/charmify/habit-tracker-service/internal/habit"
)

func (s *Server) ListHabits(ctx context.Context, _ *api.ListHabitsRequest) (*api.ListHabitsResponse, error) {
	s.lgr.Logf("ListHabits request received")
	habits, err := habit.ListHabits(ctx, s.db)
	if err != nil {
		return nil, err // todo wrap
	}

	return convertHabitsToAPI(habits), nil
}

func convertHabitsToAPI(habits []habit.Habit) *api.ListHabitsResponse {
	hts := make([]*api.Habit, len(habits))

	for i := range habits {
		hts[i] = &api.Habit{
			Id:              string(habits[i].ID),
			Name:            string(habits[i].Name),
			Description:     habits[i].Description,
			WeeklyFrequency: int32(habits[i].WeeklyFrequency),
		}
	}

	return &api.ListHabitsResponse{
		Habits: hts,
	}
}

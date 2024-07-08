package habit_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/doniacld/charmify/habit-tracker-service/internal/habit/mocks"

	"github.com/stretchr/testify/assert"
)

func TestListHabits(t *testing.T) {
	ctx := context.Background()

	habits := []Habit{
		{
			ID:              "123",
			Name:            "walk",
			WeeklyFrequency: 5,
			CreationTime:    time.Now(),
		},
		{
			ID:              "456",
			Name:            "sleep",
			WeeklyFrequency: 7,
			CreationTime:    time.Now(),
		},
	}

	dbErr := fmt.Errorf("db unavailable")

	tests := map[string]struct {
		db             func(ctl *minimock.Controller) *mocks.HabitListerMock
		expectedErr    error
		expectedHabits []Habit
	}{
		"empty": {
			db: func(ctl *minimock.Controller) *mocks.HabitListerMock {
				db := mocks.NewHabitListerMock(ctl)
				db.FindAllMock.Expect(ctx).Return(nil, nil)
				return db
			},
			expectedErr:    nil,
			expectedHabits: nil,
		},
		"2 items": {
			db: func(ctl *minimock.Controller) *mocks.HabitListerMock {
				db := mocks.NewHabitListerMock(ctl)
				db.FindAllMock.Expect(ctx).Return(habits, nil)
				return db
			},
			expectedErr:    nil,
			expectedHabits: habits,
		},
		"error case": {
			db: func(ctl *minimock.Controller) *mocks.HabitListerMock {
				db := mocks.NewHabitListerMock(ctl)
				db.FindAllMock.Expect(ctx).Return(nil, dbErr)
				return db
			},
			expectedErr:    dbErr,
			expectedHabits: nil,
		},
	}

	for name, tc := range tests {
		name, tc := name, tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := minimock.NewController(t)

			db := tc.db(ctrl)

			got, err := ListHabits(context.Background(), db)
			assert.ErrorIs(t, err, tc.expectedErr)
			assert.ElementsMatch(t, tc.expectedHabits, got)
		})
	}
}

/******* Test List Habits without minimock *******/

// MockList is a mock for FindAll method response.
type MockList struct {
	Items []Habit
	Err   error
}

// FindAll is a mock which returns the passed list of items and error.
func (list MockList) FindAll(context.Context) ([]Habit, error) { return list.Items, list.Err }

// TestListHabitsWithoutMinimock uses a simple mock of the db.
func TestListHabitsWithoutMinimock(t *testing.T) {
	habits := []Habit{
		{
			ID:              "123",
			Name:            "walk",
			WeeklyFrequency: 5,
			CreationTime:    time.Now(),
		},
		{
			ID:              "456",
			Name:            "sleep",
			WeeklyFrequency: 7,
			CreationTime:    time.Now(),
		},
	}
	dbErr := fmt.Errorf("db unavailable")

	tests := map[string]struct {
		db             MockList
		expectedErr    error
		expectedHabits []Habit
	}{
		"empty": {
			db:             MockList{Items: nil, Err: nil},
			expectedErr:    nil,
			expectedHabits: nil,
		},
		"2 items": {
			db:             MockList{Items: habits, Err: nil},
			expectedErr:    nil,
			expectedHabits: habits,
		},
		"error case": {
			db:             MockList{Items: nil, Err: dbErr},
			expectedErr:    dbErr,
			expectedHabits: nil,
		},
	}

	for name, tc := range tests {
		name, tc := name, tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := ListHabits(context.Background(), tc.db)
			assert.ErrorIs(t, err, tc.expectedErr)
			assert.ElementsMatch(t, tc.expectedHabits, got)
		})
	}
}

package client

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/doniacld/charmify/habit-tracker-service/api"
	"github.com/doniacld/charmify/pkg/habit"
)

func New(serverAddress string) (api.HabitsClient, error) {
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient(serverAddress, creds)
	if err != nil {
		return nil, err
	}

	log.Println("connected to habit-tracker backend")
	return api.NewHabitsClient(conn), nil
}

func AddHabits(ctx context.Context, cli api.HabitsClient, habits []habit.Habit) ([]habit.Habit, error) {
	createdHabits := make([]habit.Habit, len(habits))
	for i, h := range habits {
		created, err := Add(ctx, cli, h.Name, h.Desc, h.Target)
		if err != nil {
			log.Println("failed to add habits")
		}

		createdHabits[i] = created

		for i := 0; i < h.TicksCount; i++ {
			_, err = TickHabit(ctx, cli, created.ID)
			if err != nil {
				log.Printf("couldn't tick habit %v: %s", h, err)
			}
		}

		createdHabits[i].TicksCount = created.TicksCount

		log.Printf("habit %v created\n", h)
	}

	return createdHabits, nil
}

func Add(ctx context.Context, cli api.HabitsClient, name, desc string, freq int) (habit.Habit, error) {
	f := int32(freq)
	resp, err := cli.CreateHabit(ctx, &api.CreateHabitRequest{
		Name:            name,
		Description:     desc,
		WeeklyFrequency: &f,
	})
	if err != nil {
		return habit.Habit{}, err
	}

	status, err := cli.GetHabitStatus(ctx, &api.GetHabitStatusRequest{HabitId: resp.Habit.Id})
	if err != nil {
		return habit.Habit{}, err
	}

	h := habit.Habit{
		ID:         resp.Habit.Id,
		Name:       resp.Habit.Name,
		Desc:       resp.Habit.Description,
		Target:     int(resp.Habit.WeeklyFrequency),
		TicksCount: int(status.TicksCount),
	}

	log.Printf("habit %s (%d) created\n", name, freq)

	return h, nil
}

func ListHabits(ctx context.Context, cli api.HabitsClient) ([]habit.Habit, error) {
	resp, err := cli.ListHabits(ctx, &api.ListHabitsRequest{})
	if err != nil {
		return nil, err
	}

	habits := make([]habit.Habit, len(resp.Habits))
	for i, h := range resp.Habits {
		getHabitResp, err := GetHabit(ctx, cli, h.Id)
		if err != nil {
			log.Printf("failed to retrieve ticks count for habit id %s and name %s\n", h.Id, h.Name)
		}

		habits[i] = habit.Habit{
			ID:         h.Id,
			Name:       h.Name,
			Desc:       h.Description,
			Target:     int(h.WeeklyFrequency),
			TicksCount: getHabitResp.TicksCount,
		}
	}

	log.Printf("list habits returned %d entries\n", len(habits))

	return habits, nil
}

func TickHabit(ctx context.Context, cli api.HabitsClient, id string) (habit.Habit, error) {
	_, err := cli.TickHabit(ctx, &api.TickHabitRequest{
		HabitId: id,
	})
	if err != nil {
		return habit.Habit{}, err
	}

	h, err := GetHabit(ctx, cli, id)
	if err != nil {
		return habit.Habit{}, err
	}

	log.Printf("habit %s ticked (now at %d out of %d)\n", id, h.TicksCount, h.Target)

	return h, nil
}

func GetHabit(ctx context.Context, cli api.HabitsClient, id string) (habit.Habit, error) {
	resp, err := cli.GetHabitStatus(ctx, &api.GetHabitStatusRequest{HabitId: id})
	if err != nil {
		return habit.Habit{}, err
	}

	h := habit.Habit{
		ID:         resp.Habit.Id,
		Name:       resp.Habit.Name,
		Desc:       resp.Habit.Description,
		Target:     int(resp.Habit.WeeklyFrequency),
		TicksCount: int(resp.TicksCount),
	}

	log.Printf("habit %s retrieved\n", id)

	return h, nil
}

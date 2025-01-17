// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package mocks

//go:generate minimock -i learngo-pockets/habits/internal/habit.tickAdder -o tick_adder_mock.go -n TickAdderMock -p mocks

import (
	"context"
	mm_habit "github.com/doniacld/charmify/habit-tracker-service/internal/habit"
	"sync"
	mm_atomic "sync/atomic"
	"time"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// TickAdderMock implements habit.tickAdder
type TickAdderMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcAddTick          func(ctx context.Context, id mm_habit.ID, t time.Time) (err error)
	inspectFuncAddTick   func(ctx context.Context, id mm_habit.ID, t time.Time)
	afterAddTickCounter  uint64
	beforeAddTickCounter uint64
	AddTickMock          mTickAdderMockAddTick
}

// NewTickAdderMock returns a mock for habit.tickAdder
func NewTickAdderMock(t minimock.Tester) *TickAdderMock {
	m := &TickAdderMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AddTickMock = mTickAdderMockAddTick{mock: m}
	m.AddTickMock.callArgs = []*TickAdderMockAddTickParams{}

	return m
}

type mTickAdderMockAddTick struct {
	mock               *TickAdderMock
	defaultExpectation *TickAdderMockAddTickExpectation
	expectations       []*TickAdderMockAddTickExpectation

	callArgs []*TickAdderMockAddTickParams
	mutex    sync.RWMutex
}

// TickAdderMockAddTickExpectation specifies expectation struct of the tickAdder.AddTick
type TickAdderMockAddTickExpectation struct {
	mock    *TickAdderMock
	params  *TickAdderMockAddTickParams
	results *TickAdderMockAddTickResults
	Counter uint64
}

// TickAdderMockAddTickParams contains parameters of the tickAdder.AddTick
type TickAdderMockAddTickParams struct {
	ctx context.Context
	id  mm_habit.ID
	t   time.Time
}

// TickAdderMockAddTickResults contains results of the tickAdder.AddTick
type TickAdderMockAddTickResults struct {
	err error
}

// Expect sets up expected params for tickAdder.AddTick
func (mmAddTick *mTickAdderMockAddTick) Expect(ctx context.Context, id mm_habit.ID, t time.Time) *mTickAdderMockAddTick {
	if mmAddTick.mock.funcAddTick != nil {
		mmAddTick.mock.t.Fatalf("TickAdderMock.AddTick mock is already set by Set")
	}

	if mmAddTick.defaultExpectation == nil {
		mmAddTick.defaultExpectation = &TickAdderMockAddTickExpectation{}
	}

	mmAddTick.defaultExpectation.params = &TickAdderMockAddTickParams{ctx, id, t}
	for _, e := range mmAddTick.expectations {
		if minimock.Equal(e.params, mmAddTick.defaultExpectation.params) {
			mmAddTick.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAddTick.defaultExpectation.params)
		}
	}

	return mmAddTick
}

// Inspect accepts an inspector function that has same arguments as the tickAdder.AddTick
func (mmAddTick *mTickAdderMockAddTick) Inspect(f func(ctx context.Context, id mm_habit.ID, t time.Time)) *mTickAdderMockAddTick {
	if mmAddTick.mock.inspectFuncAddTick != nil {
		mmAddTick.mock.t.Fatalf("Inspect function is already set for TickAdderMock.AddTick")
	}

	mmAddTick.mock.inspectFuncAddTick = f

	return mmAddTick
}

// Return sets up results that will be returned by tickAdder.AddTick
func (mmAddTick *mTickAdderMockAddTick) Return(err error) *TickAdderMock {
	if mmAddTick.mock.funcAddTick != nil {
		mmAddTick.mock.t.Fatalf("TickAdderMock.AddTick mock is already set by Set")
	}

	if mmAddTick.defaultExpectation == nil {
		mmAddTick.defaultExpectation = &TickAdderMockAddTickExpectation{mock: mmAddTick.mock}
	}
	mmAddTick.defaultExpectation.results = &TickAdderMockAddTickResults{err}
	return mmAddTick.mock
}

// Set uses given function f to mock the tickAdder.AddTick method
func (mmAddTick *mTickAdderMockAddTick) Set(f func(ctx context.Context, id mm_habit.ID, t time.Time) (err error)) *TickAdderMock {
	if mmAddTick.defaultExpectation != nil {
		mmAddTick.mock.t.Fatalf("Default expectation is already set for the tickAdder.AddTick method")
	}

	if len(mmAddTick.expectations) > 0 {
		mmAddTick.mock.t.Fatalf("Some expectations are already set for the tickAdder.AddTick method")
	}

	mmAddTick.mock.funcAddTick = f
	return mmAddTick.mock
}

// When sets expectation for the tickAdder.AddTick which will trigger the result defined by the following
// Then helper
func (mmAddTick *mTickAdderMockAddTick) When(ctx context.Context, id mm_habit.ID, t time.Time) *TickAdderMockAddTickExpectation {
	if mmAddTick.mock.funcAddTick != nil {
		mmAddTick.mock.t.Fatalf("TickAdderMock.AddTick mock is already set by Set")
	}

	expectation := &TickAdderMockAddTickExpectation{
		mock:   mmAddTick.mock,
		params: &TickAdderMockAddTickParams{ctx, id, t},
	}
	mmAddTick.expectations = append(mmAddTick.expectations, expectation)
	return expectation
}

// Then sets up tickAdder.AddTick return parameters for the expectation previously defined by the When method
func (e *TickAdderMockAddTickExpectation) Then(err error) *TickAdderMock {
	e.results = &TickAdderMockAddTickResults{err}
	return e.mock
}

// AddTick implements habit.tickAdder
func (mmAddTick *TickAdderMock) AddTick(ctx context.Context, id mm_habit.ID, t time.Time) (err error) {
	mm_atomic.AddUint64(&mmAddTick.beforeAddTickCounter, 1)
	defer mm_atomic.AddUint64(&mmAddTick.afterAddTickCounter, 1)

	if mmAddTick.inspectFuncAddTick != nil {
		mmAddTick.inspectFuncAddTick(ctx, id, t)
	}

	mm_params := TickAdderMockAddTickParams{ctx, id, t}

	// Record call args
	mmAddTick.AddTickMock.mutex.Lock()
	mmAddTick.AddTickMock.callArgs = append(mmAddTick.AddTickMock.callArgs, &mm_params)
	mmAddTick.AddTickMock.mutex.Unlock()

	for _, e := range mmAddTick.AddTickMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmAddTick.AddTickMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAddTick.AddTickMock.defaultExpectation.Counter, 1)
		mm_want := mmAddTick.AddTickMock.defaultExpectation.params
		mm_got := TickAdderMockAddTickParams{ctx, id, t}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAddTick.t.Errorf("TickAdderMock.AddTick got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmAddTick.AddTickMock.defaultExpectation.results
		if mm_results == nil {
			mmAddTick.t.Fatal("No results are set for the TickAdderMock.AddTick")
		}
		return (*mm_results).err
	}
	if mmAddTick.funcAddTick != nil {
		return mmAddTick.funcAddTick(ctx, id, t)
	}
	mmAddTick.t.Fatalf("Unexpected call to TickAdderMock.AddTick. %v %v %v", ctx, id, t)
	return
}

// AddTickAfterCounter returns a count of finished TickAdderMock.AddTick invocations
func (mmAddTick *TickAdderMock) AddTickAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddTick.afterAddTickCounter)
}

// AddTickBeforeCounter returns a count of TickAdderMock.AddTick invocations
func (mmAddTick *TickAdderMock) AddTickBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddTick.beforeAddTickCounter)
}

// Calls returns a list of arguments used in each call to TickAdderMock.AddTick.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAddTick *mTickAdderMockAddTick) Calls() []*TickAdderMockAddTickParams {
	mmAddTick.mutex.RLock()

	argCopy := make([]*TickAdderMockAddTickParams, len(mmAddTick.callArgs))
	copy(argCopy, mmAddTick.callArgs)

	mmAddTick.mutex.RUnlock()

	return argCopy
}

// MinimockAddTickDone returns true if the count of the AddTick invocations corresponds
// the number of defined expectations
func (m *TickAdderMock) MinimockAddTickDone() bool {
	for _, e := range m.AddTickMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddTickMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddTickCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddTick != nil && mm_atomic.LoadUint64(&m.afterAddTickCounter) < 1 {
		return false
	}
	return true
}

// MinimockAddTickInspect logs each unmet expectation
func (m *TickAdderMock) MinimockAddTickInspect() {
	for _, e := range m.AddTickMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TickAdderMock.AddTick with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddTickMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddTickCounter) < 1 {
		if m.AddTickMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to TickAdderMock.AddTick")
		} else {
			m.t.Errorf("Expected call to TickAdderMock.AddTick with params: %#v", *m.AddTickMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddTick != nil && mm_atomic.LoadUint64(&m.afterAddTickCounter) < 1 {
		m.t.Error("Expected call to TickAdderMock.AddTick")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *TickAdderMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockAddTickInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *TickAdderMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *TickAdderMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAddTickDone()
}

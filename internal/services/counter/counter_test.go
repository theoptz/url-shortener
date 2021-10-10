package counter

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/theoptz/url-shortener/internal/interfaces/irepositories"
	"github.com/theoptz/url-shortener/internal/repositories"
)

func TestNewCounter(t *testing.T) {
	testCases := []struct {
		startIndex int64
		rangeItem  repositories.RangeItem
		rangeRepo  irepositories.RangeRepository

		expectedIndex int64
	}{
		{
			startIndex:    0,
			rangeItem:     repositories.RangeItem{Start: 100, End: 200},
			expectedIndex: 100,
		},
		{
			rangeRepo:     &irepositories.MockRangeRepository{},
			startIndex:    1000,
			rangeItem:     repositories.RangeItem{Start: 100, End: 200},
			expectedIndex: 1000,
		},
	}

	for _, tc := range testCases {
		counter := NewCounter(tc.rangeRepo, tc.startIndex, tc.rangeItem)

		assert.Equal(t, tc.rangeRepo, counter.rangeRepo)
		assert.Equal(t, tc.expectedIndex, counter.index, "index should be %d", tc.expectedIndex)
	}
}

func TestCounter_Inc(t *testing.T) {
	testCases := []struct {
		startIndex int64
		rangeItem  repositories.RangeItem

		returnedRange *repositories.RangeItem
		returnedErr   error

		expectedValue int64
		expectedErr   error
	}{
		{
			startIndex: 0,
			rangeItem:  repositories.RangeItem{Start: 0, End: 99},

			expectedValue: 1,
		},
		{
			startIndex: 98,
			rangeItem:  repositories.RangeItem{Start: 0, End: 99},

			expectedValue: 99,
		},
		{
			startIndex: max,
			rangeItem:  repositories.RangeItem{Start: 0, End: max + 1},

			expectedValue: 0,
			expectedErr:   errors.New("limit exceeded"),
		},
		{
			startIndex: 99,
			rangeItem:  repositories.RangeItem{Start: 0, End: 99},

			returnedErr: errors.New("unknown error"),

			expectedValue: 0,
			expectedErr:   errors.New("unknown error"),
		},
		{
			startIndex: 99,
			rangeItem:  repositories.RangeItem{Start: 0, End: 99},

			returnedRange: &repositories.RangeItem{Start: 100, End: 199},

			expectedValue: 100,
		},
	}

	for _, tc := range testCases {
		repo := &irepositories.MockRangeRepository{}
		repo.On("GetNext", mock.MatchedBy(func(_ context.Context) bool { return true })).
			Return(tc.returnedRange, tc.returnedErr)

		counter := NewCounter(repo, tc.startIndex, tc.rangeItem)

		val, err := counter.Inc()
		assert.Equal(t, tc.expectedValue, val)
		assert.Equal(t, tc.expectedErr, err)
	}
}

func TestCounter_Inc_Sequentially(t *testing.T) {
	counter := NewCounter(&irepositories.MockRangeRepository{}, 0, repositories.RangeItem{Start: 0, End: 100})

	testCases := []struct {
		expected int64
	}{
		{
			expected: 1,
		},
		{
			expected: 2,
		},
		{
			expected: 3,
		},
	}

	for _, tc := range testCases {
		res, _ := counter.Inc()
		assert.Equal(t, tc.expected, res, "should return %d", tc.expected)
	}
}

func TestCounter_Inc_Concurrently(t *testing.T) {
	counter := NewCounter(&irepositories.MockRangeRepository{}, 0, repositories.RangeItem{Start: 0, End: 100})
	wg := &sync.WaitGroup{}

	count := 10
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func() {
			_, _ = counter.Inc()
			wg.Done()
		}()
	}

	wg.Wait()

	assert.Equal(t, int64(count), counter.index, "index should be %d", int64(count))
}

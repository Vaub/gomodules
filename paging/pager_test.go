package paging

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type PagerMock struct {
	mock.Mock
}

func (m *PagerMock) Len() int {
	args := m.Called()
	return args.Int(0)
}

func (m *PagerMock) PerPage() int {
	args := m.Called()
	return args.Int(0)
}

func TestListFromPage(t *testing.T) {
	aPager := new(PagerMock)
	aPager.On("Len").Return(8)
	aPager.On("PerPage").Return(3)

	pages, _ := ListPage(aPager, 2)
	assert.Equal(t, 3, pages.From, "Did not receive the right From item")
	assert.Equal(t, 6, pages.To, "Did not receive the To item")
}

func TestGivenAPagerWithNoItemWhenListFromPageShouldThrowAnError(t *testing.T) {
	aPager := new(PagerMock)
	aPager.On("Len").Return(0)
	aPager.On("PerPage").Return(3)

	_, err := ListPage(aPager, 2)
	assert.Error(t, err, "Should be in error")
}

func TestListPageRange(t *testing.T) {
	aPager := new(PagerMock)
	aPager.On("Len").Return(8)
	aPager.On("PerPage").Return(3)

	pages, _ := ListPageRange(aPager, 1, 2)
	assert.Equal(t, 0, pages.From, "Did not receive the right From")
	assert.Equal(t, 6, pages.To, "Did not receive the right To")
}

func TestWhenListStartingFromAPageThatDoesNotExistsShouldThrowAnError(t *testing.T) {
	aPager := new(PagerMock)
	aPager.On("Len").Return(8)
	aPager.On("PerPage").Return(3)

	_, err := ListPageRange(aPager, 4, 5)
	assert.Error(t, err, "Expected an error!")
}

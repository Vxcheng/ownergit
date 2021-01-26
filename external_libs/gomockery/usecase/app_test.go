package usecase

import (
	"ownergit/external_libs/gomockery/domain/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMockery(t *testing.T) {
	c := &mocks.C{}
	a := &mocks.A{}
	b := &mocks.B{}

	c.On("CreateB", mock.Anything).Return(b)
	collector := &cImpl{
		C: c,
	}

	got := collector.Collect(a)
	assert.Equal(t, b, got)
}

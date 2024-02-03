package logger_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subrahamanyam341/andes-core-21/core/check"
	logger "github.com/subrahamanyam341/andes-logger-21"
	"github.com/subrahamanyam341/andes-logger-21/mock"
)

func TestNewLogLineWrapperFormatter_NilMarshalizerShouldErr(t *testing.T) {
	t.Parallel()

	llwf, err := logger.NewLogLineWrapperFormatter(nil)

	assert.True(t, check.IfNil(llwf))
	assert.Equal(t, logger.ErrNilMarshalizer, err)
}

func TestNewLogLineWrapperFormatter_ShouldWork(t *testing.T) {
	t.Parallel()

	llwf, err := logger.NewLogLineWrapperFormatter(&mock.MarshalizerStub{})

	assert.False(t, check.IfNil(llwf))
	assert.Nil(t, err)
}

//------- Output

func TestLogLineWrapperFormatter_OutputNilLogLineWrapperShouldRetNil(t *testing.T) {
	t.Parallel()

	llwf, _ := logger.NewLogLineWrapperFormatter(&mock.MarshalizerStub{})

	buff := llwf.Output(nil)

	assert.Nil(t, buff)
}

func TestLogLineWrapperFormatter_OutputMarshalizerErrorsShouldRetNil(t *testing.T) {
	t.Parallel()

	llwf, _ := logger.NewLogLineWrapperFormatter(&mock.MarshalizerStub{
		MarshalCalled: func(obj interface{}) (bytes []byte, e error) {
			return nil, errors.New("")
		},
	})

	buff := llwf.Output(&logger.LogLineWrapper{})

	assert.Nil(t, buff)
}

func TestLogLineWrapperFormatter_OutputShouldWork(t *testing.T) {
	t.Parallel()

	marshalizedData := []byte("test data")
	llwf, _ := logger.NewLogLineWrapperFormatter(&mock.MarshalizerStub{
		MarshalCalled: func(obj interface{}) (bytes []byte, e error) {
			return marshalizedData, nil
		},
	})

	buff := llwf.Output(&logger.LogLineWrapper{})

	assert.Equal(t, marshalizedData, buff)
}

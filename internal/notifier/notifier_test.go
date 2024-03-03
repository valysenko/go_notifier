package notifier

import (
	"go_notifier/internal/common"
	"go_notifier/internal/notifier/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotifierProvider(t *testing.T) {
	notifier1 := mocks.NewNotifier(t)
	notifier2 := mocks.NewNotifier(t)
	provider := NewNotifierProvider(notifier1, notifier2)

	t.Run("notifier provider returns first notifier", func(t *testing.T) {
		notifier1.On("Supports", common.FIREBASE_APP).Once().Return(true)
		notifier2.AssertNotCalled(t, "Supports", common.FIREBASE_APP)
		n := provider.Provide(common.FIREBASE_APP)
		assert.Equal(t, notifier1, n)
	})

	t.Run("notifier provider returns second notifier", func(t *testing.T) {
		notifier1.On("Supports", common.FIREBASE_APP).Once().Return(false)
		notifier2.On("Supports", common.FIREBASE_APP).Once().Return(true)
		n := provider.Provide(common.FIREBASE_APP)
		assert.Equal(t, notifier2, n)
	})

	t.Run("notifier provider returns no notifier", func(t *testing.T) {
		notifier1.On("Supports", common.FIREBASE_APP).Once().Return(false)
		notifier2.On("Supports", common.FIREBASE_APP).Once().Return(false)
		n := provider.Provide(common.FIREBASE_APP)
		assert.Equal(t, nil, n)
	})

}

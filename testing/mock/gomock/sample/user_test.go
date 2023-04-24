package user_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	user "github.com/dushaoshuai/go-usage-examples/testing/mock/gomock/sample"
)

func TestGrabPointer(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockIndex := NewMockIndex(ctrl)
	mockIndex.EXPECT().Ptr(gomock.Any()).SetArg(0, 7) // set first argument to 7

	i := user.GrabPointer(mockIndex)
	if i != 7 {
		t.Errorf("want 7, got %d", i)
	}
}

func TestExpectTrueNil(t *testing.T) {
	// Make sure that passing "nil" to Expect (thus as a nil interface value),
	// will correctly match a nil concrete type.
	ctrl := gomock.NewController(t)

	mockIndex := NewMockIndex(ctrl)
	mockIndex.EXPECT().Ptr(nil).AnyTimes() // this nil is a nil interface
	mockIndex.Ptr(nil)
}

func TestDoAndReturnSignature(t *testing.T) {
	t.Run("wrong number of return args", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockIndex := NewMockIndex(ctrl)

		mockIndex.EXPECT().Slice(gomock.Any(), gomock.Any()).DoAndReturn(
			func(_ []int, _ []byte) {},
		)

		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic didn't happen")
			}
		}()

		mockIndex.Slice([]int{0}, []byte("meow"))
	})

	t.Run("wrong type of return arg", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockIndex := NewMockIndex(ctrl)

		mockIndex.EXPECT().Slice(gomock.Any(), gomock.Any()).DoAndReturn(
			func(_ []int, _ []byte) bool {
				return true
			})

		mockIndex.Slice([]int{0}, []byte("meow"))
	})

}

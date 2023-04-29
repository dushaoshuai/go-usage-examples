package user_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	user "github.com/dushaoshuai/go-usage-examples/testing/mock/gomock/sample"
	"github.com/dushaoshuai/go-usage-examples/testing/mock/gomock/sample/imp1"
)

func TestRemember(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockIndex := NewMockIndex(ctrl)
	mockIndex.EXPECT().Put("a", 1)
	mockIndex.EXPECT().Put("b", gomock.Eq(2))

	mockIndex.EXPECT().NillableRet()
	boolc := make(chan bool)
	mockIndex.EXPECT().ConcreteRet().Return(boolc)
	mockIndex.EXPECT().ConcreteRet().Return(nil)

	mockIndex.EXPECT().Ellip("%d", 0, 1, 1, 2, 3)
	tri := []any{1, 3, 6, 10, 15}
	mockIndex.EXPECT().Ellip("%d", tri...)
	mockIndex.EXPECT().EllipOnly(gomock.Eq("arg"))

	user.Remember(mockIndex, []string{"a", "b"}, []any{1, 2})
	if c := mockIndex.ConcreteRet(); c != boolc {
		t.Errorf("ConcreteRet: got %v, want %v", c, boolc)
	}
	if c := mockIndex.ConcreteRet(); c != nil {
		t.Errorf("ConcreteRet: got %v, want nil", c)
	}

	calledString := ""
	mockIndex.EXPECT().Put(gomock.Any(), gomock.Any()).Do(func(key string, _ any) {
		calledString = key
	})
	mockIndex.EXPECT().NillableRet()
	user.Remember(mockIndex, []string{"blah"}, []any{7})
	if calledString != "blah" {
		t.Fatalf("%q != %q", calledString, "blah")
	}

	mockIndex.EXPECT().Put("nil-key", gomock.Any()).Do(func(key string, value any) {
		if value != nil {
			t.Errorf("Put did not pass through nil; got %v", value)
		}
	})
	mockIndex.EXPECT().NillableRet()
	user.Remember(mockIndex, []string{"nil-key"}, []any{nil})
}

func TestGrabPointer(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockIndex := NewMockIndex(ctrl)
	mockIndex.EXPECT().Ptr(gomock.Any()).SetArg(0, 7) // set first argument to 7

	i := user.GrabPointer(mockIndex)
	if i != 7 {
		t.Errorf("want 7, got %d", i)
	}
}

func TestVariadicFunction(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockIndex := NewMockIndex(ctrl)
	mockIndex.EXPECT().Ellip("%d", 5, 6, 7, 8).Do(func(format string, nums ...int) {
		sum := 0
		for _, value := range nums {
			sum += value
		}
		if sum != 26 {
			t.Errorf("want %d, got %d", 5+6+7+8, sum)
		}
	})
	mockIndex.EXPECT().Ellip("%d", gomock.Any()).Do(func(format string, nums ...int) {
		sum := 0
		for _, v := range nums {
			sum += v
		}
		if sum != 10 {
			t.Errorf("want 10, got %d", sum)
		}
	})
	mockIndex.EXPECT().Ellip("%d", gomock.Any()).Do(func(format string, nums ...int) {
		sum := 0
		for _, v := range nums {
			sum += v
		}
		if sum != 0 {
			t.Errorf("want 0, got %d", sum)
		}
	})
	mockIndex.EXPECT().Ellip("%d", gomock.Any()).Do(func(format string, nums ...int) {
		sum := 0
		for _, v := range nums {
			sum += v
		}
		if sum != 0 {
			t.Errorf("want 0, got %d", sum)
		}
	})
	mockIndex.EXPECT().Ellip("%d", gomock.Any()).Do(func(format string, nums ...int) {
		sum := 0
		for _, v := range nums {
			sum += v
		}
		if sum != 0 {
			t.Errorf("want 0, got %d", sum)
		}
	})

	mockIndex.Ellip("%d", 1, 2, 3, 4)
	mockIndex.Ellip("%d", 5, 6, 7, 8)
	mockIndex.Ellip("%d", 0)
	mockIndex.Ellip("%d")
	mockIndex.Ellip("%d")
}

func TestEmbeddedInterface(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockEmbed := NewMockEmbed(ctrl)
	mockEmbed.EXPECT().RegularMethod()
	mockEmbed.EXPECT().EmbeddedMethod()
	mockEmbed.EXPECT().ForeignEmbeddedMethod()

	mockEmbed.RegularMethod()
	mockEmbed.EmbeddedMethod()
	var emb imp1.ForeignEmbedded = mockEmbed // also does interface check
	emb.ForeignEmbeddedMethod()
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

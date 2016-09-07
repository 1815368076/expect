package expect_test

import (
	"testing"

	"github.com/a8m/expect"
)

// TODO(Ariel): Create mock that implement TB interface
// and stub `Error` and `Fatal`

func TestStartWith(t *testing.T) {
	expect := expect.New(t)
	expect("foo").To.StartWith("f")
	expect("foo").Not.To.StartWith("bar")
}

func TestEndWith(t *testing.T) {
	expect := expect.New(t)
	expect("bar").To.EndWith("ar")
	expect("bar").Not.To.EndWith("az")
}

func TestContains(t *testing.T) {
	expect := expect.New(t)
	expect("foobar").To.Contains("ba")
	expect("foobar").Not.To.Contains("ga")
}

func TestMatch(t *testing.T) {
	expect := expect.New(t)
	expect("Foo").To.Match("(?i)foo")
}

func TestEqual(t *testing.T) {
	expect := expect.New(t)
	expect("a").To.Equal("a")
	expect(1).To.Equal(1)
	expect(false).Not.To.Equal("true")
	expect(map[int]int{}).To.Equal(map[int]int{})
	expect(struct{ X, Y int }{1, 2}).Not.To.Equal(&struct{ X, Y int }{1, 2})
}

func TestPanic(t *testing.T) {
	expect := expect.New(t)
	expect(func() {}).Not.To.Panic()
	expect(func() {
		panic("foo")
	}).To.Panic()
	expect(func() {
		panic("bar")
	}).To.Panic("bar")
}

func TestToChaining(t *testing.T) {
	expect := expect.New(t)
	expect("foobarbaz").To.StartWith("foo").And.EndWith("baz").And.Contains("bar")
	expect("foo").Not.To.StartWith("bar").And.EndWith("baz").And.Contains("bob")
	expect("foo").To.Match("f").And.Match("(?i)F")
}

func TestToFailNow(t *testing.T) {
	mockT := newMockT()
	expect := expect.New(mockT)
	expect("foo").To.Equal("foo").Else.FailNow()
	select {
	case <-mockT.FailNowCalled:
		t.Errorf("Expected FailNow() on passing test not to be called")
	default:
	}
	expect("foo").To.Equal("bar").Else.FailNow()
	select {
	case <-mockT.FailNowCalled:
	default:
		t.Errorf("Expected FailNow() on failing test to be called")
	}
}

func TestNotToFailNow(t *testing.T) {
	mockT := newMockT()
	expect := expect.New(mockT)
	expect("foo").Not.To.Equal("bar").Else.FailNow()
	select {
	case <-mockT.FailNowCalled:
		t.Errorf("Expected FailNow() on passing test not to be called")
	default:
	}
	expect("foo").Not.To.Equal("foo").Else.FailNow()
	select {
	case <-mockT.FailNowCalled:
	default:
		t.Errorf("Expected FailNow() on failing test to be called")
	}
}

func TestToAndHaveFailNow(t *testing.T) {
	mockT := newMockT()
	expect := expect.New(mockT)
	expect("foo").To.Equal("bar").And.Have.Len(3).Else.FailNow()
	select {
	case <-mockT.FailNowCalled:
	default:
		t.Errorf("Expected FailNow() on failing test to be called")
	}
}

func TestToAndBeFailNow(t *testing.T) {
	mockT := newMockT()
	expect := expect.New(mockT)
	expect("foo").To.Equal("bar").And.Be.String().Else.FailNow()
	select {
	case <-mockT.FailNowCalled:
	default:
		t.Errorf("Expected FailNow() on failing test to be called")
	}
}

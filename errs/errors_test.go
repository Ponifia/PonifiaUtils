package errs

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func Test_newShouldCreateErrorFromString(t *testing.T) {
	NoLines = true
	BasePath, _ = os.Getwd()
	var err error = New("error occured")
	if nil == err {
		t.Fatal("should return error")
	}
	if err.Error() != "/errors_test.go (error occured)" {
		t.Fatalf("wrong error %s", err.Error())
	}
}

func Test_newShouldCreateErrorFromFormattedString(t *testing.T) {
	NoLines = true
	BasePath, _ = os.Getwd()
	var err error = New("error occured %s, %d", "test", 4)
	if nil == err {
		t.Fatal("should return error")
	}
	if err.Error() != "/errors_test.go (error occured test, 4)" {
		t.Fatalf("wrong error %s", err.Error())
	}
}

func Test_newShouldCreateErrorFromError(t *testing.T) {
	NoLines = true
	BasePath, _ = os.Getwd()
	e := fmt.Errorf("error occured")
	var err error = New(e)
	if nil == err {
		t.Fatal("should return error")
	}
	if err.Error() != "/errors_test.go (error occured)" {
		t.Fatalf("wrong error %s", err.Error())
	}
}

func Test_newShouldCreateErrorNested(t *testing.T) {
	NoLines = true
	BasePath, _ = os.Getwd()
	e := fmt.Errorf("error occured")
	inner := New(e)
	inner2 := New(inner)
	var err error = New(inner2)
	if nil == err {
		t.Fatal("should return error")
	}
}

func Test_newShouldCreateErrorNestedMsg(t *testing.T) {
	NoLines = true
	BasePath, _ = os.Getwd()
	e := fmt.Errorf("error occured")
	inner := New(e, "inner")
	inner2 := New(inner, "inner %d", 2)
	var err error = New(inner2, "top %s", "cool")
	if nil == err {
		t.Fatal("should return error")
	}
}

var ErrTest = fmt.Errorf("test error")

func Test_errorIs(t *testing.T) {
	var err error = New(ErrTest, "top %s", "cool")
	if !errors.Is(err, ErrTest) {
		t.Fatalf("should be %v", ErrTest)
	}
}

func Test_errorIs2(t *testing.T) {
	var ErroTest = New()
	var err error = New(ErroTest, "top %s", "cool")
	if !errors.Is(err, ErroTest) {
		t.Fatalf("should be %v", ErroTest)
	}
	if !errors.Is(err, err) {
		t.Fatalf("should be %v", err)
	}
}
func Test_errorIs3(t *testing.T) {
	var ErroTest = New(ErrTest)
	var err error = New(ErroTest, "top %s", "cool")
	if !errors.Is(err, ErrTest) {
		t.Fatalf("should be %v", ErrTest)
	}
	if !errors.Is(err, ErroTest) {
		t.Fatalf("should be %v", ErroTest)
	}
	if !errors.Is(err, err) {
		t.Fatalf("should be %v", err)
	}
}

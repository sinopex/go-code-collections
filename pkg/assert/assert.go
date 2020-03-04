package assert

import (
	"fmt"
	"reflect"
	"testing"
)

func Equal(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf(fmt.Sprintf("Not equal: \n"+"expected: %s\n"+"actual  : %s\n", expected, actual))
	}
}
func NotEqual(t *testing.T, expected, actual interface{}) {
	if reflect.DeepEqual(expected, actual) {
		t.Errorf(fmt.Sprintf("equal: \n"+"expected: %s\n"+"actual  : %s\n", expected, actual))
	}
}
func Nil(t *testing.T, actual interface{}) {
	if actual != nil {
		t.Errorf(fmt.Sprintf("not nil: actual  : %+#v\n", actual))
	}
}

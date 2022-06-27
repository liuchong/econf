package econf

import (
	"fmt"
	"os"
	"testing"
)

type myConf1 struct {
	Key1    string
	MyKey22 int64
	Key333  string
}

func TestSetFields(t *testing.T) {
	var myTestConf1 = myConf1{
		Key1:    "",
		MyKey22: 0,
		Key333:  "",
	}

	v1 := "my value 1"
	v2 := int64(123)
	os.Setenv("MY_CONF_1_KEY_1", v1)
	os.Setenv("MY_CONF_1_MY_KEY_22", fmt.Sprintf("%d", v2))

	SetFields(&myTestConf1)

	if myTestConf1.Key1 != v1 {
		t.Errorf("Test econf set fields failed. Expect %s, actual %s", v1, myTestConf1.Key1)
	}
	if myTestConf1.MyKey22 != v2 {
		t.Errorf("Test econf set fields failed. Expect %d, actual %d", v2, myTestConf1.MyKey22)
	}
	if myTestConf1.Key333 != "" {
		t.Errorf("Test econf set fields failed. Expect empty string, actual %s", myTestConf1.Key333)
	}
}

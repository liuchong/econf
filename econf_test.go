package econf

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

type myConf1 struct {
	Key1        string
	MYKey22     int64
	Key333      string
	KeyListNUM  []int32
	KEYListStr1 []string
	KeyListStr2 []string
}

func TestSetFields(t *testing.T) {
	var myTestConf1 = myConf1{
		Key1:        "",
		MYKey22:     0,
		Key333:      "",
		KeyListNUM:  []int32{0},
		KEYListStr1: nil,
		KeyListStr2: nil,
	}

	v1 := "my value 1"
	v2 := int64(123)
	l3 := []int32{138, 186}
	l4 := []string{"hello", "world", "foo bar", ""}
	_ = os.Setenv("MY_CONF_1_KEY_1", v1)
	_ = os.Setenv("MY_CONF_1_MY_KEY_22", fmt.Sprintf("%d", v2))
	_ = os.Setenv("MY_CONF_1_KEY_LIST_NUM", fmt.Sprintf("%d,%d", l3[0], l3[1]))
	_ = os.Setenv("MY_CONF_1_KEY_LIST_STR_1", strings.Join(l4, ","))
	_ = os.Setenv("MY_CONF_1_KEY_LIST_STR_2", strings.Join(l4, "#"))
	_ = os.Setenv("MY_CONF_1_DB_NAME", "my_db_name")

	SetFields(&myTestConf1)

	if myTestConf1.Key1 != v1 {
		t.Errorf("Test econf set fields failed. Expect %s, actual %s", v1, myTestConf1.Key1)
	}
	if myTestConf1.MYKey22 != v2 {
		t.Errorf("Test econf set fields failed. Expect %d, actual %d", v2, myTestConf1.MYKey22)
	}
	if myTestConf1.Key333 != "" {
		t.Errorf("Test econf set fields failed. Expect empty string, actual %s", myTestConf1.Key333)
	}
	if myTestConf1.KeyListNUM[0] != l3[0] || myTestConf1.KeyListNUM[1] != l3[1] {
		t.Errorf("Test econf set fields failed. Expect number list, actual %+v", myTestConf1.KeyListNUM)
	}
	if myTestConf1.KEYListStr1 == nil || len(myTestConf1.KEYListStr1) != 4 ||
		myTestConf1.KEYListStr1[0] != l4[0] || myTestConf1.KEYListStr1[1] != l4[1] ||
		myTestConf1.KEYListStr1[2] != l4[2] || myTestConf1.KEYListStr1[3] != l4[3] {
		t.Errorf("Test econf set fields failed. Expect string list, actual %+v", myTestConf1.KEYListStr1)
	}
	if myTestConf1.KeyListStr2 == nil || len(myTestConf1.KeyListStr2) != 1 || myTestConf1.KeyListStr2[0] != strings.Join(l4, "#") {
		t.Errorf("Test econf set fields failed. Expect string list, actual %+v", myTestConf1.KeyListStr2)
	}

	SetFieldByNameWithSep(&myTestConf1, "KeyListStr2", "#")
	if myTestConf1.KeyListStr2 == nil || len(myTestConf1.KEYListStr1) != 4 ||
		myTestConf1.KeyListStr2[0] != l4[0] || myTestConf1.KeyListStr2[1] != l4[1] ||
		myTestConf1.KeyListStr2[2] != l4[2] || myTestConf1.KeyListStr2[3] != l4[3] {
		t.Errorf("Test econf set fields failed. Expect string list, actual %+v", myTestConf1.KeyListStr2)
	}
}

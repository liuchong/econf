package econf

import (
	"fmt"
	"os"
	"reflect"
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

type TestConfig struct {
	BoolValue    bool
	FloatValue   float64
	Float32Value float32
	StringValue  string
	IntValue     int
	BoolSlice    []bool
	FloatSlice   []float64
	Float32Slice []float32
	StringSlice  []string
	IntSlice     []int
}

type ConfigWithPrivateFields struct {
	PublicStr    string
	privateStr   string
	PublicInt    int
	privateInt   int
	publicSlice  []string
	privateSlice []int
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

func TestTypesParsing(t *testing.T) {
	tests := []struct {
		name     string
		envVars  map[string]string
		expected TestConfig
	}{
		{
			name: "Test basic types",
			envVars: map[string]string{
				"TEST_CONFIG_BOOL_VALUE":     "true",
				"TEST_CONFIG_FLOAT_VALUE":    "123.456",
				"TEST_CONFIG_FLOAT_32_VALUE": "789.123",
				"TEST_CONFIG_STRING_VALUE":   "test",
				"TEST_CONFIG_INT_VALUE":      "42",
			},
			expected: TestConfig{
				BoolValue:    true,
				FloatValue:   123.456,
				Float32Value: 789.123,
				StringValue:  "test",
				IntValue:     42,
			},
		},
		{
			name: "Test slice types",
			envVars: map[string]string{
				"TEST_CONFIG_BOOL_SLICE":     "true,false,true",
				"TEST_CONFIG_FLOAT_SLICE":    "1.1,2.2,3.3",
				"TEST_CONFIG_FLOAT_32_SLICE": "4.4,5.5,6.6",
				"TEST_CONFIG_STRING_SLICE":   "a,b,c",
				"TEST_CONFIG_INT_SLICE":      "1,2,3",
			},
			expected: TestConfig{
				BoolSlice:    []bool{true, false, true},
				FloatSlice:   []float64{1.1, 2.2, 3.3},
				Float32Slice: []float32{4.4, 5.5, 6.6},
				StringSlice:  []string{"a", "b", "c"},
				IntSlice:     []int{1, 2, 3},
			},
		},
		{
			name: "Test invalid bool",
			envVars: map[string]string{
				"TEST_CONFIG_BOOL_VALUE": "not_a_bool",
			},
			expected: TestConfig{},
		},
		{
			name: "Test invalid float",
			envVars: map[string]string{
				"TEST_CONFIG_FLOAT_VALUE": "not_a_float",
			},
			expected: TestConfig{},
		},
		{
			name: "Test invalid int",
			envVars: map[string]string{
				"TEST_CONFIG_INT_VALUE": "not_an_int",
			},
			expected: TestConfig{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clear environment variables first
			os.Clearenv()

			// Set environment variables for the test
			for k, v := range tt.envVars {
				_ = os.Setenv(k, v)
			}

			// Create config
			cfg := &TestConfig{}

			// For invalid value tests, we expect a panic
			if strings.Contains(tt.name, "invalid") {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected panic for invalid value")
					}
				}()
			}

			SetFields(cfg)

			// Skip validation for invalid test cases as they should panic
			if !strings.Contains(tt.name, "invalid") {
				// Check bool value
				if tt.envVars["TEST_CONFIG_BOOL_VALUE"] != "" && cfg.BoolValue != tt.expected.BoolValue {
					t.Errorf("BoolValue = %v, want %v", cfg.BoolValue, tt.expected.BoolValue)
				}

				// Check float values
				if tt.envVars["TEST_CONFIG_FLOAT_VALUE"] != "" && cfg.FloatValue != tt.expected.FloatValue {
					t.Errorf("FloatValue = %v, want %v", cfg.FloatValue, tt.expected.FloatValue)
				}
				if tt.envVars["TEST_CONFIG_FLOAT_32_VALUE"] != "" && cfg.Float32Value != tt.expected.Float32Value {
					t.Errorf("Float32Value = %v, want %v", cfg.Float32Value, tt.expected.Float32Value)
				}

				// Check string value
				if tt.envVars["TEST_CONFIG_STRING_VALUE"] != "" && cfg.StringValue != tt.expected.StringValue {
					t.Errorf("StringValue = %v, want %v", cfg.StringValue, tt.expected.StringValue)
				}

				// Check int value
				if tt.envVars["TEST_CONFIG_INT_VALUE"] != "" && cfg.IntValue != tt.expected.IntValue {
					t.Errorf("IntValue = %v, want %v", cfg.IntValue, tt.expected.IntValue)
				}

				// Check slices
				if tt.envVars["TEST_CONFIG_BOOL_SLICE"] != "" {
					if !reflect.DeepEqual(cfg.BoolSlice, tt.expected.BoolSlice) {
						t.Errorf("BoolSlice = %v, want %v", cfg.BoolSlice, tt.expected.BoolSlice)
					}
				}
				if tt.envVars["TEST_CONFIG_FLOAT_SLICE"] != "" {
					if !reflect.DeepEqual(cfg.FloatSlice, tt.expected.FloatSlice) {
						t.Errorf("FloatSlice = %v, want %v", cfg.FloatSlice, tt.expected.FloatSlice)
					}
				}
				if tt.envVars["TEST_CONFIG_FLOAT_32_SLICE"] != "" {
					if !reflect.DeepEqual(cfg.Float32Slice, tt.expected.Float32Slice) {
						t.Errorf("Float32Slice = %v, want %v", cfg.Float32Slice, tt.expected.Float32Slice)
					}
				}
				if tt.envVars["TEST_CONFIG_STRING_SLICE"] != "" {
					if !reflect.DeepEqual(cfg.StringSlice, tt.expected.StringSlice) {
						t.Errorf("StringSlice = %v, want %v", cfg.StringSlice, tt.expected.StringSlice)
					}
				}
				if tt.envVars["TEST_CONFIG_INT_SLICE"] != "" {
					if !reflect.DeepEqual(cfg.IntSlice, tt.expected.IntSlice) {
						t.Errorf("IntSlice = %v, want %v", cfg.IntSlice, tt.expected.IntSlice)
					}
				}
			}
		})
	}
}

func TestPrivateFields(t *testing.T) {
	tests := []struct {
		name     string
		envVars  map[string]string
		expected ConfigWithPrivateFields
	}{
		{
			name: "Test private and public fields",
			envVars: map[string]string{
				"CONFIG_WITH_PRIVATE_FIELDS_PUBLIC_STR":    "public string",
				"CONFIG_WITH_PRIVATE_FIELDS_PRIVATE_STR":   "private string",
				"CONFIG_WITH_PRIVATE_FIELDS_PUBLIC_INT":    "42",
				"CONFIG_WITH_PRIVATE_FIELDS_PRIVATE_INT":   "24",
				"CONFIG_WITH_PRIVATE_FIELDS_PUBLIC_SLICE":  "a,b,c",
				"CONFIG_WITH_PRIVATE_FIELDS_PRIVATE_SLICE": "1,2,3",
			},
			expected: ConfigWithPrivateFields{
				PublicStr:    "public string",
				privateStr:   "private string",
				PublicInt:    42,
				privateInt:   24,
				publicSlice:  []string{"a", "b", "c"},
				privateSlice: []int{1, 2, 3},
			},
		},
		{
			name: "Test partial fields",
			envVars: map[string]string{
				"CONFIG_WITH_PRIVATE_FIELDS_PRIVATE_STR": "only private",
				"CONFIG_WITH_PRIVATE_FIELDS_PUBLIC_INT":  "100",
			},
			expected: ConfigWithPrivateFields{
				privateStr: "only private",
				PublicInt:  100,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clear environment variables first
			os.Clearenv()

			// Set environment variables for the test
			for k, v := range tt.envVars {
				os.Setenv(k, v)
			}

			// Create config
			cfg := &ConfigWithPrivateFields{}
			SetFields(cfg)

			// Check public string
			if tt.envVars["CONFIG_WITH_PRIVATE_FIELDS_PUBLIC_STR"] != "" &&
				cfg.PublicStr != tt.expected.PublicStr {
				t.Errorf("PublicStr = %v, want %v", cfg.PublicStr, tt.expected.PublicStr)
			}

			// Check private string
			if tt.envVars["CONFIG_WITH_PRIVATE_FIELDS_PRIVATE_STR"] != "" &&
				cfg.privateStr != tt.expected.privateStr {
				t.Errorf("privateStr = %v, want %v", cfg.privateStr, tt.expected.privateStr)
			}

			// Check public int
			if tt.envVars["CONFIG_WITH_PRIVATE_FIELDS_PUBLIC_INT"] != "" &&
				cfg.PublicInt != tt.expected.PublicInt {
				t.Errorf("PublicInt = %v, want %v", cfg.PublicInt, tt.expected.PublicInt)
			}

			// Check private int
			if tt.envVars["CONFIG_WITH_PRIVATE_FIELDS_PRIVATE_INT"] != "" &&
				cfg.privateInt != tt.expected.privateInt {
				t.Errorf("privateInt = %v, want %v", cfg.privateInt, tt.expected.privateInt)
			}

			// Check public slice
			if tt.envVars["CONFIG_WITH_PRIVATE_FIELDS_PUBLIC_SLICE"] != "" {
				if !reflect.DeepEqual(cfg.publicSlice, tt.expected.publicSlice) {
					t.Errorf("publicSlice = %v, want %v", cfg.publicSlice, tt.expected.publicSlice)
				}
			}

			// Check private slice
			if tt.envVars["CONFIG_WITH_PRIVATE_FIELDS_PRIVATE_SLICE"] != "" {
				if !reflect.DeepEqual(cfg.privateSlice, tt.expected.privateSlice) {
					t.Errorf("privateSlice = %v, want %v", cfg.privateSlice, tt.expected.privateSlice)
				}
			}
		})
	}
}

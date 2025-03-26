package envs

import (
	"testing"
)

func TestGetStringIsValid(t *testing.T) {
	t.Setenv("EXIST_ENV", "success")
	value, _ := GetString("EXIST_ENV", "failure")
	if value != "success" {
		t.Errorf("FAIL - actual: %s, expected: %s", value, "success")
	}

	value, _ = GetString("NONEXIST_ENV", "failure")
	if value != "failure" {
		t.Errorf("FAIL - actual: %s, expected: %s", value, "failure")
	}
}

func TestMustGetStringIsValid(t *testing.T) {
	t.Setenv("EXIST_ENV", "success")
	value := MustGetString("EXIST_ENV")
	if value != "success" {
		t.Errorf("FAIL - actual: %s, expected: %s", value, "success")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Must panic!!")
		}
	}()
	_ = MustGetString("NONEXIST_ENV")
}

func TestGetIntIsValid(t *testing.T) {
	t.Setenv("EXIST_ENV", "200")
	value, _ := GetInt("EXIST_ENV", 400)
	if value != 200 {
		t.Errorf("FAIL - actual: %d, expected: %d", value, 200)
	}

	value, _ = GetInt("NONEXIST_ENV", 400)
	if value != 400 {
		t.Errorf("FAIL - actual: %d, expected: %d", value, 400)
	}

	t.Setenv("EXIST_ENV", "non_int_value")
	_, err := GetInt("EXIST_ENV", 200)
	if err == nil {
		t.Error("Must raise error")
	}
}

func TestMustGetIntIsValid(t *testing.T) {
	t.Setenv("EXIST_ENV", "200")
	value := MustGetInt("EXIST_ENV")
	if value != 200 {
		t.Errorf("FAIL - actual: %d, expected: %d", value, 200)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Must panic!!")
		}
	}()
	_ = MustGetInt("NONEXIST_ENV")
}

func TestGetBoolIsValid(t *testing.T) {
	t.Setenv("EXIST_ENV", "True")
	value, _ := GetBool("EXIST_ENV", false)
	if value != true {
		t.Errorf("FAIL - actual: %t, expected: %t", value, true)
	}

	value, _ = GetBool("NONEXIST_ENV", false)
	if value != false {
		t.Errorf("FAIL - actual: %t, expected: %t", value, false)
	}

	t.Setenv("EXIST_ENV", "non_bool_value")
	_, err := GetBool("EXIST_ENV", true)
	if err == nil {
		t.Error("Must returned error")
	}
}

func TestMustGetBoolIsValid(t *testing.T) {
	t.Setenv("EXIST_ENV", "True")
	value := MustGetBool("EXIST_ENV")
	if value != true {
		t.Errorf("FAIL - actual: %t, expected: %t", value, true)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Must panic!!")
		}
	}()
	_ = MustGetBool("NONEXIST_ENV")
}

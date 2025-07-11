package object

import "testing"

func TestStringHasheKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is Johnny"}
	diff2 := &String{Value: "My name is Johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("string with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("string with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("string with different content have same hash keys")
	}
}

func TestBooleanHasheKey(t *testing.T) {
	hello1 := &Boolean{Value: true}
	hello2 := &Boolean{Value: true}
	diff1 := &Boolean{Value: false}
	diff2 := &Boolean{Value: false}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("boolean with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("boolean with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("boolean with different content have same hash keys")
	}
}

func TestIntegerHasheKey(t *testing.T) {
	hello1 := &Integer{Value: 1}
	hello2 := &Integer{Value: 1}
	diff1 := &Integer{Value: 2}
	diff2 := &Integer{Value: 2}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("integer with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("integer with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("integer with different content have same hash keys")
	}
}

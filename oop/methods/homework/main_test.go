package main

import (
	"math"
	"testing"
)

func Test_multiply1(t *testing.T) {
	calc := NewCalc()
	calc.SetA(10).SetB(10).Do(multiply)
	result := calc.Result()
	expected := 100.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_multiply2(t *testing.T) {
	calc := NewCalc()
	calc.SetA(0).SetB(10).Do(multiply)
	result := calc.Result()
	expected := 0.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_multiply3(t *testing.T) {
	calc := NewCalc()
	calc.SetA(5).SetB(10).Do(multiply)
	result := calc.Result()
	expected := 50.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_multiply4(t *testing.T) {
	calc := NewCalc()
	calc.SetA(180).SetB(-10).Do(multiply)
	result := calc.Result()
	expected := -1800.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_multiply5(t *testing.T) {
	calc := NewCalc()
	calc.SetA(-10).SetB(5).Do(multiply)
	result := calc.Result()
	expected := -50.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_multiply6(t *testing.T) {
	calc := NewCalc()
	calc.SetA(10).SetB(0).Do(multiply)
	result := calc.Result()
	expected := 0.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_multiply7(t *testing.T) {
	calc := NewCalc()
	calc.SetA(10000).SetB(10).Do(multiply)
	result := calc.Result()
	expected := 100000.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_multiply8(t *testing.T) {
	calc := NewCalc()
	calc.SetA(2.2).SetB(2).Do(multiply)
	result := calc.Result()
	expected := 4.4
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_Divide1(t *testing.T) {
	calc := NewCalc()
	calc.SetA(-10).SetB(-5).Do(divide)
	result := calc.Result()
	expected := 2.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_Divide2(t *testing.T) {
	calc := NewCalc()
	calc.SetA(10).SetB(0).Do(divide)
	result := calc.Result()
	if !math.IsInf(result, 1) {
		t.Errorf("Результат = %v , нужно = +Inf", result)
	}
}

func Test_Divide3(t *testing.T) {
	calc := NewCalc()
	calc.SetA(0).SetB(-5).Do(divide)
	result := calc.Result()
	expected := 0.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_Divide4(t *testing.T) {
	calc := NewCalc()
	calc.SetA(100).SetB(10).Do(divide)
	result := calc.Result()
	expected := 10.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_Divide5(t *testing.T) {
	calc := NewCalc()
	calc.SetA(8).SetB(0.5).Do(divide)
	result := calc.Result()
	expected := 16.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_Divide6(t *testing.T) {
	calc := NewCalc()
	calc.SetA(-100).SetB(10).Do(divide)
	result := calc.Result()
	expected := -10.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_Divide7(t *testing.T) {
	calc := NewCalc()
	calc.SetA(0.2).SetB(0.2).Do(divide)
	result := calc.Result()
	expected := 1.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_Divide8(t *testing.T) {
	calc := NewCalc()
	calc.SetA(50).SetB(-0.5).Do(divide)
	result := calc.Result()
	expected := -100.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_Sum1(t *testing.T) {
	calc := NewCalc()
	calc.SetA(10).SetB(5).Do(sum)
	result := calc.Result()
	expected := 15.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_Sum2(t *testing.T) {
	calc := NewCalc()
	calc.SetA(-10).SetB(5).Do(sum)
	result := calc.Result()
	expected := -5.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_Sum3(t *testing.T) {
	calc := NewCalc()
	calc.SetA(10).SetB(-5).Do(sum)
	result := calc.Result()
	expected := 5.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_Sum4(t *testing.T) {
	calc := NewCalc()
	calc.SetA(-10).SetB(-5).Do(sum)
	result := calc.Result()
	expected := -15.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_Sum5(t *testing.T) {
	calc := NewCalc()
	calc.SetA(10).SetB(0).Do(sum)
	result := calc.Result()
	expected := 10.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_Sum6(t *testing.T) {
	calc := NewCalc()
	calc.SetA(10.5).SetB(5.5).Do(sum)
	result := calc.Result()
	expected := 16.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_Sum7(t *testing.T) {
	calc := NewCalc()
	calc.SetA(1.1).SetB(1.1).Do(sum)
	result := calc.Result()
	expected := 2.2
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_Sum8(t *testing.T) {
	calc := NewCalc()
	calc.SetA(0).SetB(0).Do(sum)
	result := calc.Result()
	expected := 0.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_Average1(t *testing.T) {
	calc := NewCalc()
	calc.SetA(10).SetB(5).Do(average)
	result := calc.Result()
	expected := 7.5
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_Average2(t *testing.T) {
	calc := NewCalc()
	calc.SetA(0).SetB(5).Do(average)
	result := calc.Result()
	expected := 2.5
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_Average3(t *testing.T) {
	calc := NewCalc()
	calc.SetA(10).SetB(0).Do(average)
	result := calc.Result()
	expected := 5.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_Average4(t *testing.T) {
	calc := NewCalc()
	calc.SetA(0).SetB(0).Do(average)
	result := calc.Result()
	expected := 0.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_Average5(t *testing.T) {
	calc := NewCalc()
	calc.SetA(-10).SetB(-5).Do(average)
	result := calc.Result()
	expected := -7.5
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_Average6(t *testing.T) {
	calc := NewCalc()
	calc.SetA(10).SetB(-5).Do(average)
	result := calc.Result()
	expected := 2.5
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_Average7(t *testing.T) {
	calc := NewCalc()
	calc.SetA(10.5).SetB(5.5).Do(average)
	result := calc.Result()
	expected := 8.0
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

func Test_Average8(t *testing.T) {
	calc := NewCalc()
	calc.SetA(0.5).SetB(0.5).Do(average)
	result := calc.Result()
	expected := 0.5
	if result != expected {
		t.Errorf("Результат = %v , нужно = %v", result, expected)
	}
}

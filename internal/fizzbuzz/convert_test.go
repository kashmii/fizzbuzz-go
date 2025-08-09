package fizzbuzz

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		// 3の倍数のテスト
		{
			name:     "3の倍数（3）",
			input:    3,
			expected: "fizz",
		},
		{
			name:     "3の倍数（6）",
			input:    6,
			expected: "fizz",
		},
		{
			name:     "3の倍数（9）",
			input:    9,
			expected: "fizz",
		},
		// 5の倍数のテスト
		{
			name:     "5の倍数（5）",
			input:    5,
			expected: "buzz",
		},
		{
			name:     "5の倍数（10）",
			input:    10,
			expected: "buzz",
		},
		{
			name:     "5の倍数（20）",
			input:    20,
			expected: "buzz",
		},
		// 15の倍数（3と5の両方の倍数）のテスト
		{
			name:     "15の倍数（15）",
			input:    15,
			expected: "fizzbuzz",
		},
		{
			name:     "15の倍数（30）",
			input:    30,
			expected: "fizzbuzz",
		},
		{
			name:     "15の倍数（45）",
			input:    45,
			expected: "fizzbuzz",
		},
		// 3でも5でも割り切れない数のテスト
		{
			name:     "通常の数（1）",
			input:    1,
			expected: "1",
		},
		{
			name:     "通常の数（2）",
			input:    2,
			expected: "2",
		},
		{
			name:     "通常の数（4）",
			input:    4,
			expected: "4",
		},
		{
			name:     "通常の数（7）",
			input:    7,
			expected: "7",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Convert(tt.input)
			if result != tt.expected {
				t.Errorf("Convert(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// テーブル駆動テストのより簡単な例
// func TestConvertSimple(t *testing.T) {
// 	if result := Convert(3); result != "fizz" {
// 		t.Errorf("Convert(3) = %q, want %q", result, "fizz")
// 	}

// 	if result := Convert(5); result != "buzz" {
// 		t.Errorf("Convert(5) = %q, want %q", result, "buzz")
// 	}

// 	if result := Convert(15); result != "fizzbuzz" {
// 		t.Errorf("Convert(15) = %q, want %q", result, "fizzbuzz")
// 	}

// 	if result := Convert(1); result != "1" {
// 		t.Errorf("Convert(1) = %q, want %q", result, "1")
// 	}
// }
package main

import (
	"bytes"
	"os"
	"testing"
)

func TestValidateInput(t *testing.T) {
	tests := []struct {
		name           string
		input          int
		expectedResult bool
		expectedOutput string
	}{
		// 正常値のテスト
		{
			name:           "正常値: 1",
			input:          1,
			expectedResult: true,
			expectedOutput: "",
		},
		{
			name:           "正常値: 50",
			input:          50,
			expectedResult: true,
			expectedOutput: "",
		},
		{
			name:           "正常値: 100",
			input:          100,
			expectedResult: true,
			expectedOutput: "",
		},
		// 境界値のテスト
		{
			name:           "境界値: 0（無効）",
			input:          0,
			expectedResult: false,
			expectedOutput: "Invalid input: 0 must be greater than 0\n",
		},
		{
			name:           "境界値: 101（無効）",
			input:          101,
			expectedResult: false,
			expectedOutput: "Invalid input: 101 must be 100 or less\n",
		},
		// 異常値のテスト
		{
			name:           "負の値: -1",
			input:          -1,
			expectedResult: false,
			expectedOutput: "Invalid input: -1 must be greater than 0\n",
		},
		{
			name:           "大きすぎる値: 200",
			input:          200,
			expectedResult: false,
			expectedOutput: "Invalid input: 200 must be 100 or less\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 標準出力をキャプチャするためのバッファを準備
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// テスト対象の関数を実行
			result := validateInput(tt.input)

			// 標準出力を元に戻し、出力内容を取得
			w.Close()
			os.Stdout = old

			var buf bytes.Buffer
			buf.ReadFrom(r)
			output := buf.String()

			// 結果の検証
			if result != tt.expectedResult {
				t.Errorf("validateInput(%d) = %v, want %v", tt.input, result, tt.expectedResult)
			}

			if output != tt.expectedOutput {
				t.Errorf("validateInput(%d) output = %q, want %q", tt.input, output, tt.expectedOutput)
			}
		})
	}
}

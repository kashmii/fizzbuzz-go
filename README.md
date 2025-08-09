# fizzbuzz-go

Go言語で実装したFizzBuzzプログラムです。

## 概要

このプログラムは、1から指定された数値まで順番に数字を表示し、以下のルールに従って文字列を出力します：

- 3の倍数のときは「fizz」
- 5の倍数のときは「buzz」
- 3と5の両方の倍数（15の倍数）のときは「fizzbuzz」
- それ以外のときは数字をそのまま表示

## 使い方

```bash
go run cmd/fizzbuzz/main.go <数値>
```

### 例

```bash
# 1から15まで出力
go run cmd/fizzbuzz/main.go 15
```

出力例：
```
1
2
fizz
4
buzz
fizz
7
8
fizz
buzz
11
fizz
13
14
fizzbuzz
```

## コマンドライン引数について

- **第1引数のみを使用**: コマンドライン引数は1つ目のみを使用し、2つ目以降の引数は無視されます
- **数値以外の入力**: 数値以外の文字列が入力された場合、エラーメッセージが表示されます

## 入力値の制限

このプログラムには以下の入力値制限があります：

- **0以下の数値**: 0以下の数値（負の数を含む）を入力するとエラーが表示されます
- **101以上の数値**: 101以上の数値は大きすぎる数字として扱われ、エラーが表示されます
- **有効な範囲**: 1から100までの整数のみが有効な入力です

### エラー例

```bash
# 0以下の数値
$ go run cmd/fizzbuzz/main.go 0
Invalid input: 0 must be greater than 0

# 101以上の数値
$ go run cmd/fizzbuzz/main.go 150
Invalid input: 150 must be 100 or less

# 数値以外の入力
$ go run cmd/fizzbuzz/main.go abc
Invalid input: abc is not a number

# 引数なし
$ go run cmd/fizzbuzz/main.go
Please provide a number as an argument.
```

## プロジェクト構成

```
fizzbuzz-go/
├── cmd/fizzbuzz/          # メインアプリケーション
│   ├── main.go           # エントリーポイント
│   └── main_test.go      # メイン関数のテスト
├── internal/fizzbuzz/     # 内部パッケージ
│   ├── convert.go        # FizzBuzz変換ロジック
│   └── convert_test.go   # 変換ロジックのテスト
├── go.mod                # Go モジュール定義
└── README.md             # このファイル
```

## テスト実行

```bash
# 全てのテストを実行
go test ./...

# 特定のパッケージのテストを実行
go test ./internal/fizzbuzz
go test ./cmd/fizzbuzz
```

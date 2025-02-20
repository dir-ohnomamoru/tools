# tools
utility tools for work

# CSV Splitter

## 概要
このプログラムは指定されたCSVファイルを指定した行数ごとに分割し、新しいCSVファイルとして出力します。

## 使用方法

### 実行コマンド
```sh
go run split_csv.go <csv_file_path> <split_size>
```

### 引数
- `<csv_file_path>`: 分割対象のCSVファイルのパスを指定します。
- `<split_size>`: 分割する行数を整数で指定します。

### 使用例
```sh
go run split_csv.go sample.csv 100
```
このコマンドを実行すると、`sample.csv` を100行ごとに分割し、`output_1.csv`, `output_2.csv`, ... の形式で出力します。

## 注意事項
- CSVファイルにはヘッダが含まれている必要があります。
- 各分割ファイルにはヘッダが付与されます。
- 分割後のファイル名は `output_1.csv`, `output_2.csv`, ... となります。

## ライセンス
このプロジェクトはMITライセンスの下で提供されます。


# gotestcoveragecheck

Go のテストカバレッジ基準値をCIに設定するためのサンプルレポジトリです。

## 概要

Go の標準ツールでテストカバレッジを計測し、マージ可能な基準を設定しています。


```shell
go test -v -coverprofile=coverage.out ./...
go tool cover -func=coverage.out > report.out
# 除外したいファイルがある場合
# cat coverage.out |grep -E -v '(mock_*.go|*.pb.go)' > filtered.out
# go tool cover -func=filtered.out > report.out

# coverage rate limit (ここでは 60% としている)
declare -r limit=60
if (( $(cat report.out | awk 'END{gsub("%", "", $3); print ($3 < '$limit')}') )); then
    echo 'Coverage is below the lower limit! Failed!'
    exit 1;
fi
```

## 補足

### 一部のファイルを対象外にしたい場合

生成ファイルなど、カバレッジの対象から除きたい場合などもあるでしょう。
そういうときは、 `coverageprofile` の出力を `grep` などで絞ったものを `go tool cover` に与えれば良いだけです。
逆に言うと、ファイル中の特定の箇所を対象外にするなど、これ以上細かい制御をする場合は、標準だけでやるのは難しいです。

### 各種 CI の設定例

CI への設定例も用意しました。このプロジェクトでは、例のために両方設定していますが、実際に利用する場合はどちらか片方だけで良いです。

[Github Actions の例](https://github.com/nrnrk/gotestcoveragecheck/blob/main/.github/workflows/go.yaml#L24-L37)
[CircleCI の例](https://github.com/nrnrk/gotestcoveragecheck/blob/main/.circleci/config.yml#L10-L24)

実際に[テストのない PR](https://github.com/nrnrk/gotestcoveragecheck/pull/1) を作成しました。CI のチェックが落ちている様子を確認できます。

### 課題

全体のカバレッジの値を取得するために `awk` を使って値を抜き出しているところはあまりイケていません。
標準だけでももしかしたら、もう少しシンプルにできる方法はあるかもしれません。

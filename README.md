# practice-scraping-with-go

Goでスクレイピングの練習をする

※過度なスクレイピングはやめましょう

## リポジトリをクローン

```
$ git clone https://github.com/blue-goheimochi/practice-scraping-with-go.git
```

## goqueryのインストール

```
$ go get github.com/PuerkitoBio/goquery
```

## 静的サイトのスクレイピング

[日本経済新聞](https://www.nikkei.com/)さんの `タイトル` と `アクセスランキング` をスクレイピングしてみる

### スクレイピング実行

```
$ go run static.go
```

### 実行結果

```
--- <title>タグ ----------
経済、株価、ビジネス、政治のニュース:日経電子版

--- アクセスランキング ----------
1. XXXX銀、YYYY融資「ZZZZの可能性」
2. XXXX「YYYY」販売休止　ZZZZ不足［有料会員限定］
3. XXXX銀、融資YYYY億円　ZZZZ割［有料会員限定］
4. XXXX銀行、YYYYはZZZZか［有料会員限定］
5. XXXX、YYYYへ時間との戦い　YYYYの判断ZZZZ日に
6. XXXX「YYYYの危機」　ZZZZで苦境
7. XXXXのYYYY　ZZZZ新たな一歩［有料会員限定］
8. XXXXのYYYYへ　ZZZZ期、YYYYを解消
9. XXXXがYYYYを勝手にZZZZ　いよいよYYYYに［有料会員限定］
10. XXXXのYYYY対策、ZZZZ過半が「YYYY」［有料会員限定］
```

## 動的サイトのスクレイピング

[Medium](https://medium.com/)さんの `タイトル` と `Popular on Medium` をスクレイピングしてみる

### Splashの起動

```
$ docker pull scrapinghub/splash
$ docker run -d -p 5023:5023 -p 8050:8050 -p 8051:8051 scrapinghub/splash
```

### スクレイピング実行

```
$ go run dynamic.go
```

### 実行結果

```
--- <title>タグ ----------
Medium – Read, write and share stories that matter

--- アクセスランキング ----------
01. How XXXX Launched YYYY — and ZZZZ of YYYY Just Like It
02. Why XXXX, and YYYY Are ZZZZ
03. Who XXXX to Be YYYY?
04. XXXX YYYY's ZZZZ
```
# web-scraping-example Go (golang)

[![Build Status](https://drone.io/github.com/kyokomi/web-scraping-example/status.png)](https://drone.io/github.com/kyokomi/web-scraping-example/latest)

## Install

```sh
$ go get github.com/kyokomi/web-scraping-example
```

## Usage

```sh
$ mkdir -p ~/.web-scraping-example
$ cd ~/.web-scraping-example
$ curl -O https://raw.githubusercontent.com/kyokomi/web-scraping-example/master/example/config.json
```

### config.json Sample

```
{
    "keyword": "ラブライブ",
    "outputDir": "images",
    "pageSettings": [
        {
            "baseUrl": "http://blog.livedoor.jp/nizigami/search",
            "query": "?q=",
            "findKey": ".article .article-body .article-continue a",
            "imgFindKey": ".t_b img"
        }
    ]
}
```

## Output

image in the `outputDir` of `config.json`

![](https://raw.githubusercontent.com/kyokomi/web-scraping-example/master/example/sample.png)

## LICENSE

[MIT](https://github.com/kyokomi/web-scraping-example/blob/master/LICENSE.md)

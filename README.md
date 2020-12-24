# tl2t

Internal project to convert an entries in an RSS feed into loglines in a per-category log file.

# Help

```shell
Usage of tl2t:
  -categoriesregexp string
        Default categories regexp (default "Foo|Bar")
  -feedurl string
        Default feed URL (default "https://localhost/invalid")
```

# Usage

```shell
$ ./tl2t -feedurl https://kevq.uk/feed.xml -categoriesregexp "Newsletter"
$ cat logs/newsletter
Craving Coffee 06 - Do What You Love
Craving Coffee 05 - Should I Kill This Newsletter?
Craving Coffee 04 - Simplicity Vs Complexity
```

# rsslagg

A command-line RSS link aggregator that displays the latest post links
from RSS feeds you follow.


## Installation

Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed. 

You can then install `rsslagg` with:

```bash
go install github.com/mustafa-ozturk/rsslagg
```

## Config

Create a `.rsslaggconfig.json` file in your home directory with the following structure:

```json
{
    "max_posts_displayed": 3,
    "rss_feed_links": [
        "https://www.wagslane.dev/index.xml",
        "https://blog.boot.dev/index.xml"
    ]
}

```

## Usage

Running `rsslagg` will print you the latest `max_posts_displayed` posts:

```bash
rsslagg

- 2024-12-14 | Boot.dev Blog | Functions in Python:
        https://blog.boot.dev/tutorials/python/functions/

- 2024-12-23 | Boot.dev Blog | Loops in Python:
        https://blog.boot.dev/tutorials/python/loops/

- 2024-12-25 | Boot.dev Blog | Lists in Python:
        https://blog.boot.dev/tutorials/python/lists/
```



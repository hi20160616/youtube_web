# Youtube Fetcher

Use Google Youtube Data Api v3.

Use https://infinite-scroll.com/ for infinite scroll load videos.

> Infinite scrolling the right way:   
> https://medium.com/walmartglobaltech/infinite-scrolling-the-right-way-11b098a08815  
> https://jsfiddle.net/valkyris/43fmku20/693/  

# Quick Start

1. get enit from https://github.com/hi20160616/enit  
2. `enit set yt_api <youtube api key>`
3. `./youtube_web`

Look port:
on Mac:  
```
lsof -i:8080
netstat plten | grep 8080
```
on Linux:
```
ps -A | grep 8080
netstat pantu | grep 8080
```

All your channelIds should list in `./internal/pkg/db/json/cids.txt`

# TODO

- [X] Start jobs with heartbeat
- [X] Infinite scroll loading complete on cid page.
- [X] Optimize view of index
- [X] Search implement
- [ ] Use context in deep control.
- [ ] Scrollspy index


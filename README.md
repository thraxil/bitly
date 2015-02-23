# Simple Go Library for bit.ly API

## Get it

    $ go get github.com/thraxil/bitly

## Use it

    import (
        "log"
        "github.com/thraxil/bitly"
    )
    ...
    c := bitly.NewConnection(accessToken='...')
    link, err := c.Shorten('http://www.google.com/')
    if err != nil {
        log.Fatal(err)
    }
    log.Println(link)

So far, this is an extremely limited implementation:

* the only method available is `Shorten`. Pull Requests welcome with
  support for the rest of the API.
* it only supports the generic oauth `access_token` (generate one
  [here](https://bitly.com/a/oauth_apps)). No support for the older
  API key approach (that's deprecated anyway).
* it only uses bitly's HTTPS endpoint

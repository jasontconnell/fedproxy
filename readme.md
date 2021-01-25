*** FED Proxy

A proxy for developing FED resources locally where the website is not executable locally.

Sample config:

```
{
    "proxyHost": "dev.example.dev",
    "proxyScheme": "https",
    "localHost": "example.local.dev",
    "localScheme": "https",
    "localKeyFile": "local.example.dev.key",
    "localCrtFile": "local.example.dev.crt",
    "localPort": 5454,
    "requestHeaders": {"Authorization": "Basic xxxxx"},
    "intercepts": [
        { "extension": "css", "mimeType": "text/css" },
        { "extension": "js", "mimeType": "application/javascript" }
    ],
    "localStartPath": "C:\\path\\to\\intercepted\\resources"
}
```

Any requests for the extensions in the `intercepts` section of the configuration will be intercepted and will instead be served from the `localStartPath`

Cache-Control no-cache is added to local response headers. All headers are carried through for remote requests and responses.
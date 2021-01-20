*** FED Proxy

A proxy for developing FED resources locally where the website is not executable locally.

Sample config:

```
{
    "proxyHost": "dev.example.com",
    "proxyScheme": "https",
    "localPort": 5454,
    "requestHeaders": {"Authorization": "Basic xxxxx"},
    "intercepts": [
        { "extension": "css", "mimeType": "text/css" },
        { "extension": "js", "mimeType": "application/javascript" }
    ],
    "localStartPath": "C:\\sites\\exampledev\\"
}```

Any requests for the extensions in the `intercepts` section of the configuration will be intercepted and will instead be served from the `localStartPath`

No caching headers are added to local requests. All headers are carried through for remote requests and responses.
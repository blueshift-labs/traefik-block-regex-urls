# Traefik Block Regex Urls

A [Traefik](https://github.com/traefik/traefik) plugin to block access to certain paths using a list of regex values and return a defined status code.

## Configuration

Sample configuration in Traefik.

### Configuration as local plugin

static.yml (development)

```yaml
experimental:
  localPlugins:
    block-regex-urls:
      moduleName: 'github.com/blueshift-labs/traefik-block-regex-urls'
```

static.yml (production)

```yaml
experimental:
  plugins:
    block-regex-urls:
      moduleName: 'github.com/blueshift-labs/traefik-block-regex-urls'
      version: 'v0.0.2'
```

dynamic-configuration.yml

```yaml
http:
  middlewares:
    block-scan-paths:
      plugin:
        block-regex-urls:
          exact_match:
            - "some_exact_string_with_regex_chars_?/._to_block"
          regex:
            - "^something.mydomain.tld\\/scan\\?uid=12345(.*)&gid=6789(.*)"
            - "^something.mydomain.tld\\/scan\\?uid=345$"
          statusCode: 404
```


# Traefik Block Regex Urls

A [Traefik](https://github.com/containous/traefik) plugin to block access to certain paths using a list of regex values and return a defined status code.

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
  plugins:
    block-regex-urls:
      moduleName: 'github.com/blueshift-labs/traefik-block-regex-urls'
      version: 'v0.0.1'
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

docker-compose.yml

```yaml
services:
  traefik:
    image: traefik
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - /docker/config/traefik/traefik.yml:/etc/traefik/traefik.yml
      - /docker/config/traefik/dynamic-configuration.yml:/etc/traefik/dynamic-configuration.yml
      - /plugin/traefik-block-regex-urls:/plugins-local/src/github.com/blueshift-labs/traefik-block-regex-urls/
    ports:
      - "80:80"
  hello:
    image: containous/whoami
    labels:
      - traefik.enable=true
      - traefik.http.routers.hello.entrypoints=http
      - traefik.http.routers.hello.rule=Host(`hello.localhost`)
      - traefik.http.services.hello.loadbalancer.server.port=80
      - traefik.http.routers.hello.middlewares=my-plugin@file
```

## Sample configuration

- `regex`:  List of regex values to use for url blocking.
- `exact_match`:  List of exact matching strings to use for url blocking.
- `statusCode`: Return value of the status code.

```yaml
my-block-regex-urls:
  plugin:
    block-regex-urls:
      exact_match:
        - "some_string_to_block"
      regex:
        - "^something.mydomain.tld\\/scan\\?uid=12345(.*)&gid=6789(.*)"
        - "^something.mydomain.tld\\/scan\\?uid=345$"
      statusCode: 418
```

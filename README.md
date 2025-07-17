# Traefik Block Regex Urls

A [Traefik](https://github.com/containous/traefik) plugin to block access to certain paths using a list of regex values and return a defined status code.

## Configuration

Sample configuration in Traefik.

### Configuration as local plugin

traefik.yml

```yaml

experimental:
  localPlugins:
    block-regex-urls:
      moduleName: github.com/blueshift-labs/traefik-block-regex-urls
```

dynamic-configuration.yml

```yaml
http:
  middlewares:
    block-scan-paths:
      plugin:
        block-regex-urls:
          allowLocalRequests: true
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

- `allowLocalRequests`: If set to true, will not block request from [Private IP Ranges](https://en.wikipedia.org/wiki/Private_network)
- `regex`:  List of regex values to use for url blocking.
- `strings`:  List of string values to use for url blocking.
- `statusCode`: Return value of the status code.

```yaml
my-block-regex-urls:
  plugin:
    block-regex-urls:
      allowLocalRequests: true
      exact_match:
        - "some_string_to_block"
      regex:
        - "^something.mydomain.tld\\/scan\\?uid=12345(.*)&gid=6789(.*)"
        - "^something.mydomain.tld\\/scan\\?uid=345$"
      statusCode: 418
```

## Contributors

| [<img alt="ShantanuGadgil" src="https://avatars.githubusercontent.com/u/2508915?v=4" width="117"/>](https://github.com/blueshift-labs) |
| :---------------------------------------------------------------------------------------------------------------------------------------: |
| [Shantanu Gadgil](https://github.com/blueshift-labs) |

## License

traefik-block-regex-urls is distributed under the MIT license. [See LICENSE](LICENSE) for details.

```
MIT License

Copyright (c) 2025-today Shantanu Gadgil

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

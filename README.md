# Tmpl

[![Build Status](https://github.com/tomMoulard/traefik-plugin-tmpl/actions/workflows/main.yml/badge.svg)](https://github.com/tomMoulard/traefik-plugin-tmpl/actions/workflows/main.yml)

This is a plugin for [Traefik](https://traefik.io) to FIXME.

## Usage

To use this template, you should:

- Replace all `tmpl` by the name of your plugin:
```
find -type f -exec sed -i "s/tmpl/[my-plugin-name]/" {} \;
```
- Replace all `Tmpl` by the 'pretty' name of your plugin.
- Fix all `FIXME`.

To publish your plugin, you should (see [ref](https://github.com/traefik/plugindemo#prerequisites)):

- Create a GitHub release.
- Make sure that the repository is public and have the `traefik-plugin` topic.

### Configuration

Here is an example of a file provider dynamic configuration (given here in
YAML), where the interesting part is the `http.middlewares` section:

```yaml
# Dynamic configuration

http:
  routers:
    my-tmpl-router:
      rule: host(`tmpl.localhost`)
      service: service-tmpl@file
      middlewares:
        - traefik-plugin-tmpl

  middlewares:
    traefik-plugin-tmpl:
      plugin:
        traefik-plugin-tmpl:
          foo: "bar"

  services:
   service-tmpl:
      loadBalancer:
        servers:
          - url: http://127.0.0.1:8000
```

### Local test

There is a `docker-compose.yml` file to test the plugin locally:

```bash
docker-compose up -d
```

Then, you can go to [http://tmpl.localhost](http://tmpl.localhost) to see the
result.

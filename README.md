# mika

Mika is a life log app.

## Usage

In order to run this project, you need to install Docker.

### Setup

Create a `.env` file in the root directory of your project.

```env
JWT_PRIVATE_KEY_BASE64=
JWT_PUBLIC_KEY_BASE64=

FRONTEND_URL=http://localhost

DB_USER=
DB_PASSWORD=
DB_NAME=

REDIS_PASSWORD=

DOMAIN=
```

### Development

Recommended to use [devcontainer](https://code.visualstudio.com/docs/devcontainers/containers).

If you don't use devcontainer, run the following commands.

```Shell
$ docker compose build
$ docker compose up -d
```

### Production

Only `compose.yaml` is needed. Run the following commands.

```Shell
$ docker compose -f compose.yaml build
$ docker compose -f compose.yaml up -d
```

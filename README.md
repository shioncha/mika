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
```

### Development

```Shell
docker compose -f compose.dev.yaml up --build
```

### Production

```Shell
docker compose up --build
```

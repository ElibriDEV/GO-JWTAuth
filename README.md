# GO: JWT-Auth

## Run application

### Create .env:

```
# JWT Config
SIGN_KEY="YOUR KEY"
# Days
REFRESH_TTL=180
# Minutes
ACCESS_TTL=30

# Crypt
HASH="YOUR HASH"

# Application
APPLICATION_PORT=8000

# MongoDB
MONGO_INITDB_ROOT_USERNAME="admin"
MONGO_INITDB_ROOT_PASSWORD="admin"
MONGO_PORT=27017
```

### Start application:

#### Simple Run:

```bash
$ docker-compose up
```

#### Separate run:

Add to env:
```
MONGO_URL="YOUR URL WITH CREDENTIALS"
```

Run MongoDB:

```bash
$ docker-compose -f docker-compose-local.yml up
```

Run application:

```bash
$ go build main.go
```

## Using case

### Swagger
```
http://localhost:your_port/docs/index.html
```

#### Sign-In
```
http://localhost:your_port/api/auth/sign-in/?guid=YOUR_GUID
```

You will get 2 tokens. Set them in Access and Refresh headers:
```
Access: "Bearer YOUR_ACCESS_TOKEN"
Refresh: "Bearer YOUR_REFRESH_TOKEN"
```

#### Refresh
```
http://localhost:your_port/api/auth/refresh/
```

Change your headers.

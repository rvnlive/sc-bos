Area Controller Command
=======================

This command runs an Area Controller instance.

## Data Directory

- `private-key.pem` - private key used for Smart Core, PKCS#8 wrapped in PEM
- `self-signed.crt` - self-signed X.509 certificate for `private-key.pem` - used before controller has enrolled
- `roots.pem` - certificates representing the roots of trust, used before enrollment
- `tenants.json` - a json list of tenants and their hashed secrets. Used when Tenant OAuth is enabled but no manager
  node is available to ask.
- `users.json` - a json list of users and their hashed secrets. Used to authenticate requests locally using a
  username/password
- `enrollment/`
  - `enrollment.json` - data file generated upon enrollment
  - `ca.crt` - Root CA for the Smart Core installation
  - `cert.crt` - X.509 certificate for `../private-key.pem` signed by the Root CA
- `cache/`
  - `publications/` - cache of management server publications, including configuration

## Local Authentication

The area controller may optionally be configured to accept local authentication - i.e. a username and password. This
method of authentication is less secure but useful when the area controller is running without any supporting
infrastructure, like a dedicated OAuth server.

Local authentication uses [OAuth2 password grant](https://www.oauth.com/oauth2-servers/access-tokens/password-grant/) to
exchange a username and password for a trusted token.

To setup local authentication you need to do two things:

1. Let the area controller know about the local accounts you wish to allow
2. Turn on the local password-based OAuth server

For step 1 we need to create a `users.json` file in the area controllers data directory (defaults to
`.data/area-controller-01/users.json`). The structure of this file, like `tenants.json`, is a JSON list of accounts and
their secrets:

```json
[
  {
    "id": "email@example.com",
    "title": "My Name",
    "secrets": [
      {"hash": "$2a$10$/uBhiEncrKMgJ8q5AjyRFuqe1dzTNTsOjX1noIzu/lI5JQ78EUvLO"}
    ],
    "zones": ["Floor1", "Floor2"]
  }
]
```

The secret hash can be generated using the locally provided `pash` tool:

```shell
go run github.com/vanti-dev/bsp-ew/cmd/pash           
Password: <enter your password>
$2a$10$/uBhiEncrKMgJ8q5AjyRFuqe1dzTNTsOjX1noIzu/lI5JQ78EUvLO
```

Finally we can turn on the OAuth password flow for the area controller via the `--local-auth` command line flag

```shell
go run github.com/vanti-dev/bsp-ew/cmd/area-controller --local-auth
```

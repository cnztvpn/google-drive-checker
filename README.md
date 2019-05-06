# google-drive-checker

checker of files in google drive for myself

## check

```
$ export GD_CHECKER_CRED_JSON='{"installed"~~~~}'  # Google API OAuth2 Credential JSON
$ export GD_CHECKER_PARENT_ID='' # Google Drive Directory ID

$ go test main_test.go
```

## How to get Google Drive Refresh token


- Auth Code: can get from browser, short expiry
- Refresh Token: can get from API using Auth Code, never expire until to revoke

you need to prepare
- Client ID from service account credential JSON
- Client Secret from service account credential JSON


### get Auth Code (Open your browser)
need to overwrite your project id

```
https://accounts.google.com/o/oauth2/auth?access_type=offline&client_id=<<your client id>>&redirect_uri=urn%3Aietf%3Awg%3Aoauth%3A2.0%3Aoob&response_type=code&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fdrive&state=state-token
```

-> copy auth code

### get refresh token

```
$ export AUTH_CODE=
$ export CLIENT_ID=
$ export CLIENT_SECRET=

$ curl --data "code=${AUTH_CODE}" --data "client_id=${CLIENT_ID}" --data "client_secret=${CLIENT_SECRET}" --data "redirect_uri=urn:ietf:wg:oauth:2.0:oob" --data "grant_type=authorization_code" --data "access_type=offline" https://www.googleapis.com/oauth2/v4/token

response example:
{
  "access_token": "",
  "expires_in": 3600,
  "refresh_token": "",
  "scope": "https://www.googleapis.com/auth/drive",
  "token_type": "Bearer"
}

```

-> Right!
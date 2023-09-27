## Go API authentication


#### set environment 
```
API_DB_HOST=192.168.1.1                  
API_DB_USER=postgresadmin
API_DB_PASSWORD=postgresadmin
API_DB_NAME=db
API_DB_PORT=5432
API_SECRET_KEY=example_key_1234
API_BIND_PORT=8000
API_BIND_HOST=0.0.0.0
API_PATH_PRIVATEKEY=private.pem
API_ADMIN_USERNAME='admin'
API_ADMIN_PASSWORD='admin'
```

## License

The MIT License (MIT) - see LICENSE.md for more details


## Usage 
 create super admin
 ```bash
 make createsuperuser
 ```


### TODO:
- [ ] implement integration test
- [ ] handle error connetion db in response

---
## API Document Roure

### User
#### Signin
* permission: `null`

* endpoint: `/users/signin/`

* HTTP method: `POST`

* Header: `Content-Type: application/json`

* request:
```json
{
    "username":"admin",
    "password":"admin"
}
```
* ok response:
HTTP status code: `200`
```json
{
  "error": false,
  "error_code": 0,
  "error_message": "",
  "data": {
    "id": "4b5f2ffa-c64f-4369-adc6-b2df80f21e94",
    "username": "admin",
    "email": "admin@api.co",
    "is_admin": true,
    "created_at": "2023-09-27T06:57:40.841052Z",
    "updated_at": "2023-09-27T06:57:40.841052Z",
    "api_key": "VXz5iqLhbsk+lZXKz/iQk0t02av2tpwbIn1ZTCJrEWqEv3f/Uj9bOIX094/HwA+++uLlH+wwJfFbQpWZtHrl4A==",
    "token": "eyJhbGciOiJSU0ExXzUiLCJlbmMiOiJBMTI4Q0JDLUhTMjU2In0.Tb7cNQ2sbuPk-DoRY6qTzWkw_ryiEMLuXN6lY-v3qkZOwSeVf3vLPgha9aFymrRtx-an1Tcd4a1LtjgmPMYoVwlOAlc8pm5C2i5thz9FVkXPEm9c4SWYI4La9HecIvJi8cvi-ke-dWn6T2ZL_MdmkNj_SPTNV6faGQCq3YEiJAgzcnBfKJ2beERjYKK79tbaxpYPxcXQb-qRuZij4P1BRKwiUSlBPnXEbx9UDbW5ATdfkgPrOUdj9H_fCWh6mnNzjjv6MGBCYSM1mdB6UJ7IGg3UcejEsixgsBcifgpMNhdlrC8kBxc-QGUeyUqtgDArFH_vg-RYQDgSAYT_AmXzwQ.WZDYjJOwXCqcjN-Tgtd0gQ.80xnAsdUKd2USxOP4DfBeAxPTvCPy0eqnRLx4er3Hxf5JoNjhWev4fhz16En_4QeuYPKHXV2j9VxDGDKT5zz6X8PKwDRPZdhK8Uyb3rmCIEbGETXSAjdsUNZp7hHh1DAzThv3u3wk45J_Tg_WpjIy2UZ8Xx8D9lFGp75D6tThqstpFyVNNi106NWPr8MGteQcwFq1mMBFiSp669MvGh8pIzK7s5j0vpbB-G252VoTE4PmN3760vvCqV4zBMSFnimMnr7p9-EMPnbidfBA-MGdRkTCwXbNnroMlHldtuxCy8p3eUqfkAnRVKKg-vleHPfgie4KfN3d177NxsfpjrRMdMcdPUgXe4M1PNBMMPvxR8h2Jbk7g4E4VpSBOPSG5atBIhDxQyyBqujmtIbKKzZfGLKIWe-tZbyha-9NgeVVxkvvKRw0lY5GqYtE0sBbbQNvwWaw0dXbp-2qz5F29_BXK8SpUtoGjw5HCcSzwjliTGYbCU7TGQ0ERpNhc4TkG12Nn0-Updsnz7oFTwUfHGDcRTVNhR--l5me3iF_euLKNotEfrmKug-CWmugrfmgMGMrW5-nBZ1Ned1WXEJgJmxWAsJvxODzBTRy9ugALBqeL4.boLqBbg-7P1Z_eDPcoijcw"
  }
}
```
* non existence user

HTTP status code: `404`
```json
{
  "error": false,
  "error_code": 3,
  "error_message": "record not found",
  "data": null
}
```




* wrong password

HTTP status code: `401`
```json
{
  "error": false,
  "error_code": 6,
  "error_message": "invalide credentials,wrong password",
  "data": null
}
```

### Protected Requests
all protected requests must have `Authorization` header with `token` raelm
```
Authorization: token eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyIjp7InVzZXJfaWQiOiJlMDM1Zjc1MC01ZGZlLTRiODctYThhZC1mZjkwN2FhNzUyM2EifX0.Nl8ioe4fqrosgJ4M7ifXQPCSNPyUV8L-cquPL4zPmjA
```
OR
all protected requests must have `http` header with `Api-Key` field.
```
API_KEY: VXz5iqLhbsk+lZXKz/iQk0t02av2tpwbIn1ZTCJrEWqEv3f/Uj9bOIX094/HwA+++uLlH+wwJfFbQpWZtHrl4A==
```
#### Unauthorize Responses
* Invalid bearer
HTTP status code: `403`
```json
{
  "error": true,
  "error_code": 4,
  "error_message": "bearer is not valid",
  "data": null
}
```

* Not valid token

HTTP status code: `403`
```json
{
"error": true,
"error_code": 4,
"error_message": "invalid token",
"data": null
}
```
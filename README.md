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
- [ ] test case user api
- [ ] dockerize api
- [ ] implement integration test
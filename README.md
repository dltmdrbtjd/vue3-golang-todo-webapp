# Vue3-golang-todo-webapp

## stack
- vue3
- pinia
- typescript
- golang
- gin
- google oauth2
- mongodb atlas

## IDE setting
- VScode
- Volar
- Go

## How to start
### What you need before running
#### Front-End
```js
1. cd todo-client
2. create .env file
3. env file setup
  VIET_APP_BASE_URL={ this is server url }
  VITE_GOOGLE_OAUTH2_CLIENT_ID={ your google oauth2 client id }
  VITE_GOOGLE_OAUTH2_CLIENT_PW={ your google oauth2 client pw }
4. npm ci
```
#### Back-End
```js
1. cd todo-server
2. create .env file
3. env file setup
  DB_USERNAME={ your mongodb user name }
  DB_PASSWORD={ your mongodb user password }
  DB_PORT={ your mongodb atlas port }
  DB_NAME={ your db name }
  // example 
  // uri := fmt.Sprintf("mongodb+srv://%s:%s@%s", config.DBUserName, config.DBPassword, config.DBPort) use username,password,url
  // database = client.Database(config.DBName) use db name

  PORT={this port front-end env base url port }
  ENV=production
  GOOGLE_CLIENT_ID={ your google client id // it is same to front-end .env google client id }
  GOOGLE_SECRET_KEY={ your google secret key // it is same to front-end .env google client pw }
  GOOGLE_OAUTH_REDIRECT_URL={ your google client setting redirect url }
  // bottom url is google page your setup google client id,key,redirect url
  // https://console.developers.google.com/home
  
4. make build // if your haven't "make" package, you need brew install make
5. ./cs // just start complie file
```

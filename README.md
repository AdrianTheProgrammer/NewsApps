# NewsApps
**API Documentation**
```
https://app.swaggerhub.com/apis/HAFIZDARMAWAN1996/NewsApp/1.0.0
```

**Import Packages**
```
go get github.com/joho/godotenv
go get gorm.io/driver/postgres
go get gorm.io/gorm
go get github.com/golang-jwt/jwt/v5
go get github.com/labstack/echo-jwt/v4
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware
```

**Unit Testing**
```
mockery --all
go test ./... --coverprofile cover.out
go tool cover -func cover.out
```


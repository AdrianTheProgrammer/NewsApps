# NewsApps

NewsApps adalah aplikasi sosial media sederhana yang bisa digunakan untuk memposting cerita-cerita yang ingin kamu sebarluaskan secara publik.

Fitur :
- Melihat semua hal yang telah diposting berbagai macam pengguna di halaman utama.
- Pengguna yang telah mendaftar/login bisa memposting cerita mereka sendiri serta menyisipkan sebuah gambar.
- Pengguna yang telah mendaftar/login bisa memberikan komentar kepada postingan orang lain.
- Pengguna bisa mengedit dan menghapus postingan maupun komentar yang telah mereka post.
- Pengguna juga bisa mengatur kembali informasi akun yang telah didaftarkan.


### Informasi Teknis

Go version: `Go 1.22.3`
[API Documentation](https://app.swaggerhub.com/apis/HAFIZDARMAWAN1996/NewsApp/1.0.0)

**Import Packages**
```
go get github.com/joho/godotenv
go get gorm.io/driver/postgres
go get gorm.io/gorm
go get github.com/golang-jwt/jwt/v5
go get github.com/labstack/echo-jwt/v4
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware
go get github.com/cloudinary/cloudinary-go
```

**Unit Testing**
```
mockery --all
go test ./... --coverprofile cover.out
go tool cover -func cover.out
```


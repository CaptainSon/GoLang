

gRPC:  là giao thức giao tiếp giữa request - respone thông thường tuy nhiên nó được dùng cho việc giao tiếp giữa các server với nhau nhiều hơn là client-server.

folde

internal:

​	dto: định nghĩa struct

​	services: xử lý logic, nhận dto và trả dto.

​	repository: tương tác với database

​	transformers: 

proto command:

 `protoc --proto_path=proto --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=proto/pb --go_out=proto/pb proto/*.proto`





### jwt ( json web token)

​	là đoạn mã. khi người dùng đăng nhập vào sẽ sinh ra đoạn mã. mà sau đó mỗi lần thực hiện tác vụ thì sẽ gưi kèm đoạn này để xác thực đây chính là user này mà không phải ai khác.

​	**Authentication**: Đây là trường hợp phổ biến nhất thường sử dụng JWT. Khi người dùng đã đăng nhập vào hệ thống thì những request tiếp theo từ phía người dùng sẽ chứa thêm mã JWT. Điều này cho phép người dùng được cấp quyền truy cập vào các url, service, và resource mà mã Token đó cho phép. Phương pháp này không bị ảnh hưởng bởi Cross-Origin Resource Sharing (CORS) do nó không sử dụng cookie.

​	https://viblo.asia/p/authenticate-jwt-voi-golang-p2-QpmleyzDlrd

### Validator

​	\- Govalidator cung cấp rất nhiều function hỗ trợ chúng ta validate dữ liệu theo các dạng thông dụng như: URL, Email, Alpha, Numeric Alphanumeric, Regex

​	\- Ngoài ra Govalidator còn hỗ trợ chúng ta **validate struct** (kiểm tra tính hợp lệ của các field trong struct) bằng cách sử dụng tag **valid**. Ví dụ: 	

```go
type User struct {
	Id       int    `valid:"required"`
	Name     string `valid:"required"`
	Email    string `valid:"email"`
	Password string `valid:"required"`
}
```

Sau khi khai báo Struct xong, để kiểm tra toàn bộ struct ta gọi hàm **ValidateStruct()**: 

```go
user := User{
    Id:       1,
	Name:     "",
	Phone:    "1080",
	Email:    "fake_email",
	Password: "secret",
}

if ok, err := govalidator.ValidateStruct(user); err != nil {
	fmt.Print(err)   // In ra các thông báo lỗi
} else {
	fmt.Print("Struct hợp lệ")
}
```

### Cache

​	là nơi lưu trữ tạp thời cho dữ liệu. phục vụ cho việc truy xuất dữ liệu một cách nhanh chóng.



### Redis



### Defer, Panic and Recover

+ **Link :** [tham khảo](https://go.dev/blog/defer-panic-and-recover)
+ **Defer :**  các câu lệnh có defer đằng trước đẩy chúng vào một stack. và thực hiện khi các câu lệnh trong hàm thực hiện xong.
+ **Panic :** phá vỡ một luồng bình thường và tạo ra một panicking goroutine.
+ **Recover :**  lấy lại quyền điều khiển panicking goroutine

``

```
package main

import "fmt"

func main() {
    f()
    fmt.Println("Returned normally from f.")
}

func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}
```

Kết quả:

```
Calling g.
Printing in g 0
Printing in g 1
Printing in g 2
Printing in g 3
Panicking!
Defer in g 3
Defer in g 2
Defer in g 1
Defer in g 0
Recovered in f 4
Returned normally from f.
```

ghi chú: các lời gọi hàm sau của g và dòng `fmt.Println("Returned normally from g.")` bị bỏ qua vì lời gọi panic. sau đỏ nhảy thẳng lên dòng

`if r := recover(); r != nil`

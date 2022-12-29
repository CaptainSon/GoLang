gRPC:  là giao thức giao tiếp giữa request - respone thông thường tuy nhiên nó được dùng cho việc giao tiếp giữa các server với nhau nhiều hơn là client-server.

folde

internal:

​	dto: định nghĩa struct

​	services: xử lý logic, nhận dto và trả dto.

​	repository: tương tác với database

​	transformers: 



jwt: là đoạn mã. khi người dùng đăng nhập vào sẽ sinh ra đoạn mã. mà sau đó mỗi lần thực hiện tác vụ thì sẽ gưi kèm đoạn này để xác thực.



proto command:

 `protoc --proto_path=proto --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=proto/pb --go_out=proto/pb proto/*.proto`


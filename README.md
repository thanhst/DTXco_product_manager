# **DTX-COMPANY TEST PROJECT**

## **Giới thiệu về dự án quản lý sản phẩm:**
* Có sử dụng API
* Sử dụng MongoDB để lưu trữ sản phẩm
* Sử dụng websocket của gorilla mux
* Thiết lập giao diện để sử dụng ( chưa hoàn thành)

## **Frontend**
## **Backend**
### Tổng quan:
- Hệ thống sử dụng hệ mô hình MVC + repository pattern
### *Setting*:
* Yêu  cầu cơ bản:
    1. Cài đặt mongoDB trên cổng 27017:27017 ( em sử dụng trên docker)
    2. Mongo và các thư viện khác (Đã được thêm vào folder vendor)
    3. Cổng chạy server : 8386
* Cấu trúc thư mục:
    1. public : chứa file main.go là file main để chạy chương trình ( có thể chạy bằng fresh).
    2. config : chứa file database.go là file kết nối tới mongoDB // initDB.
    3. controller: chứa các file điều khiển // websocket_controller , user_controller, product_controller.
    4. html : chứa file giao diện ( chưa code giao diện )
    5. middleware: chứa file check token jws ( check đăng nhập ) //jws_middleware.
    6. model : chứa các model khai báo cho các đối tượng //product , user.
    7. repository : kho chứa các đối tượng //product_repo , user_repo.
    8. router : chứa file cấu hình đường dẫn url // routes.
    9. service : chứa các file dịch vụ cung cấp dịch giao tiếp giữa các repository với controller // auth_service , product_service , user_service.
    10. vendor : file chứa các thư viện đã cài ( jws , websocket gorilla-mux , fresh , ...).
    11. websocket : file chứa xử lý websocket // websocket.go
    12. go.mod
    13. go.sum
### *Testing*:
* API:
    1. URL API là : http://localhost:8386/api/
    2. URL Websocket : ws://localhost:8386/ws ( không yêu cầu auth)
    3. URL user :
        * Register: http://localhost:8386/register
            1. Header: Không có yêu cầu
            2. Body : định dạng raw json
            {
                "username":"",
                "password":""
            }
        * Login: http://localhost:8386/login
            1. Header: Không có yêu cầu
            2. Body : định dạng raw json
            {
                "username":"",
                "password":""
            }
    4. URL product:
        * Get all product: http://localhost:8386/api/get
        * Get product by ID : http://localhost:8386/api/get/{id} với id là idObject của mongoDB
        * Create product : http://localhost:8386/api/create
            1. Header : yêu cầu có authorization token jws được trả về từ login
            2. Body : định dạng raw json
            {
                "name":"",
                "description":"",
                "price":
            }
        * Update product : http://localhost:8386/api/update
            1. Header : yêu cầu có authorization token jws được trả về từ login
            2. Body : định dạng raw json 
            {
                "id":"id object của mongoDB"
                "name":"",
                "description":"",
                "price":
            }
        * Delete product : http://localhost:8386/api/delete
            1. Header : yêu cầu có authorization token jws được trả về từ login
            2. Body : định dạng raw json
            {
                "id":"id object của mongoDB"
            }
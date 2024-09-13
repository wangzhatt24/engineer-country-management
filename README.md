# engineer-country-management
Engineer-Country-Management App: A lightweight application built in Go that helps manage engineers and their country assignments.

Một số điểm có thể cải thiện:

Xử lý lỗi: Code hiện tại có xu hướng in lỗi ra console thay vì trả về cho client. Nên xem xét việc log lỗi và trả về lỗi phù hợp cho client.
Giải quyết: log lỗi có thể dùng gói log hoặc thư viện, nên trả về lỗi


Cấu trúc code: Có thể tách logic database và cache thành các package riêng để dễ quản lý hơn.
Giải quyết


Sử dụng context: Nên truyền context xuống các hàm database và cache để có thể hủy các thao tác khi cần.

Xử lý đồng thời: Cần đảm bảo xử lý đúng khi có nhiều request cùng lúc, đặc biệt là khi cập nhật cache.

Quản lý kết nối: Nên sử dụng connection pool cho MySQL và Redis.

Cấu hình: Nên đưa các thông số cấu hình như địa chỉ database, Redis vào file cấu hình hoặc biến môi trường.

Validation: Chưa thấy có validate input, nên thêm vào để đảm bảo tính đúng đắn của dữ liệu.

Unit test: Chưa thấy có unit test, nên thêm vào để đảm bảo chất lượng code.

Về việc đảm bảo dữ liệu giữa redis và mysql, liệu có thể sử dụng cái này không?

Lock Redis, lock quá trình ghi bằng SETNX

Mở Transaction Mysql

Update Mysql, nếu lỗi rollback
	
Update Redis, nếu lỗi rollback Mysql

Commit Mysql

Unlock Redis

=> Không được vì khả năng đảm bảo transaction trên nhiều service đối với redis bị hạn chế, duy mysql hỗ trợ


Thử nghiệm 

Thêm chức năng đếm số lượng đếm get country bẳng redis

Khi có người gọi get country by id -> thất bại -> nil, lỗi

Nếu thành công trả về dữ liệu + update redis đếm bằng INCR( Tính Atomic) -> 

pub/sub: khi có người get country by id thành công -> đẩy vào pub -> sub đọc và log ra

queue-cron-job: cứ sau 10s sao lưu đếm trên redis vào database

Làm sao triển khai elastic search

Làm cho chức năng tìm kiếm country trước, tìm kiếm bằng tên nhé, nhưng tìm thì tìm sao?
Cho nó tìm kiếm chính xác hay fuzzy

Ok làm cả 2, 1 cho chính xác, 2 fuzzy

Trước khi làm 2 chức năng này, dữ liệu trên elastic search ở đâu ra?

- Viết bulk code golang gọi tất cả dữ liệu trong mysql -> đổ vào -> quá nhiều -> mất nhiều thời gian

- Triển khai bắt sự kiện insert, update, delete dữ liệu trong service của mình:
  + Khi insert thành công vào mysql -> Bỏ dữ liệu vào queue insert es -> Sau đó trả về dữ liệu cho người dùng(Không chặn), còn việc insert vào es để consumer lo
  + Khi update thành công vào mysql -> Bỏ vào queue update es -> Sau đó trả về dữ liệu cho người dùng(không chặn), còn việc update vào es để consumer lo
  + Khi delete thành công trong mysql cũng xử lí tương tự
  + Khi gọi search es không trả về được dữ liệu


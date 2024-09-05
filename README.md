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
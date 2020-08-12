## Websocket
* Websocket là giao thức hỗ trợ giao tiếp hai chiều giữa client và server để tạo một kết nối trao đổi dữ liệu. 
* Giao thức này không sử dụng HTTP mà thực hiện nó qua TCP. 
* Mặc dù được thiết kế để chuyên sử dụng cho các ứng dụng web, lập trình viên vẫn có thể đưa chúng vào bất kì loại ứng dụng nào.
## Ưu điểm
* WebSocket cung cấp giao thức giao tiếp hai chiều mạnh mẽ. No có độ trễ thấp và dễ xử lý lỗi.
* Websocket thường được sử dụng cho những trường hợp yêu cầu real time như chat, hiển thị biểu đồ hay thông tin chứng khoán.
## Cấu trúc WebSocket
![alt](https://stackjava.com/wp-content/uploads/2018/04/Screenshot_1-3.png)

* Giao thức chuẩn thông thường của WebSocket là ws:// , giao thức secure là wss:// . 
* Chuẩn giao tiếp là String và hỗ trợ buffered arrays và blobs.
## Các thuộc tính của WebSocket
| THUỘC TÍNH       | MÔ TẢ                                                                    |
| ---------------- |:------------------------------------------------------------------------:|
| readyState       | Diễn tả trạng thái kết nối. Nó có các giá trị sau:						  |
|				   |	* Giá trị 0: kết nối vẫn chưa được thiết lập (WebSocket.CONNECTING)   |
|				   |    * Giá trị 1: kết nối đã thiết lập và có thể giao tiếp (WebSocket.OPEN)|
|				   |	* Giá trị 2: kết nối đang qua handshake đóng (WebSocket.CLOSING)      | 
|				   |	* Giá trị 3: kết nối đã được đóng (WebSocket.CLOSED)				  |
|bufferedAmount	   | Biểu diễn số byte của UTF-8 mà đã được xếp hàng 						  |
|				   |	bởi sử dụng phương thức send()										  |
## Các sự kiện WebSocket
| SỰ KIỆN| EVENT HANDLER | MÔ TẢ |
| ------ |:-------------:|:-----:|
|open|onopen|Khi một WebSocket chuyển sang trạng thái mở, “onopen” sẽ được gọi.|
|message|onmessage|Khi WebSocket nhận dữ liệu từ Server.|
|error|onerror|Có bất kỳ lỗi nào trong giao tiếp.|
|close|onclose|Kết nối được đóng. Những sự kiện được truyền cho “onclose” có ba tham số là “code”, “reason”, và “wasClean”.|
## Các phương thức của WebSocket
| PHƯƠNG THỨC | MÔ TẢ |
| ----------- |:-----:|
| send()	  | send(data) gửi dữ liệu tới server. Message data là string, ArrayBuffer, blob.|
| close()	  |Đóng kết nối đang tồn tại.


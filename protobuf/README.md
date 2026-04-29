# protobuf

## 1. Format chung của 1 dòng benchmark

Ví dụ:

```
BenchmarkJSON-16  562732  1791 ns/op  576 B/op  14 allocs/op
```

Nó có 5 phần chính:

```
BenchmarkName-GOMAXPROCS  iterations  time/op  bytes/op  allocs/op
```

---

## 2. Giải thích từng cột

### (1) `BenchmarkJSON-16`

* `BenchmarkJSON` → tên benchmark function của bạn
* `-16` → số **GOMAXPROCS** (số thread OS Go dùng để chạy benchmark)

👉 16 nghĩa là Go runtime đang dùng 16 logical CPUs để chạy test song song (không phải 16 goroutine của bạn).

---

### (2) `562732` (iterations)

Đây là số lần loop `b.N` được chạy.

Ví dụ:

```go
for i := 0; i < b.N; i++
```

* JSON: 562,732 lần chạy
* Protobuf: 2,404,488 lần chạy

👉 Protobuf nhanh hơn nên Go tăng số lần chạy nhiều hơn để đo ổn định.

---

### (3) `1791 ns/op` (quan trọng nhất)

👉 Thời gian trung bình cho 1 operation (1 vòng loop)

* JSON: ~1791 nanoseconds/op (~1.79 µs)
* Protobuf: ~514.7 ns/op (~0.51 µs)

👉 Kết luận:

* Protobuf ~3.5x nhanh hơn JSON trong case của bạn

---

### (4) `576 B/op`

👉 Số bytes heap cấp phát trung bình mỗi operation

* JSON: 576 bytes/op
* Protobuf: 200 bytes/op

👉 Ý nghĩa:

* JSON tạo nhiều object/intermediate string/struct hơn
* Protobuf tối ưu hơn về memory layout (binary encoding)

---

### (5) `14 allocs/op`

👉 số lần cấp phát memory (heap allocations) mỗi operation

* JSON: 14 allocations
* Protobuf: 8 allocations

👉 càng ít allocs:

* càng ít GC pressure
* càng tốt cho hệ thống high-throughput (banking, API, SDK)

---

## 3. Tổng kết so sánh của bạn

| Metric      | JSON       | Protobuf  | Kết luận              |
| ----------- | ---------- | --------- | --------------------- |
| Speed       | 1791 ns/op | 514 ns/op | Protobuf ~3.5x faster |
| Memory      | 576 B/op   | 200 B/op  | Protobuf ~2.8x less   |
| Allocations | 14         | 8         | Protobuf ít GC hơn    |

---

## 4. Vì sao Protobuf nhanh hơn?

Trong case của bạn:

#### JSON:

* parse text (string-based)
* encode/decode runtime reflection (`encoding/json`)
* nhiều intermediate allocations

#### Protobuf:

* binary format (fixed schema)
* generated code (không reflection nặng)
* ít parsing overhead
* memory layout compact hơn

---

## 5. Lưu ý quan trọng (hay bị hiểu sai)

### ❌ Không phải lúc nào Protobuf cũng thắng JSON

Protobuf thắng khi:

* high throughput API
* microservices internal
* mobile SDK / network serialization

JSON có lợi khi:

* debugging / human-readable
* public API
* flexibility schema

---

### ⚠️ Benchmark của bạn đang đo gì?

Bạn đang benchmark:

```go
Marshal + Unmarshal
```

=> tức là **end-to-end serialization cost**, không phải chỉ encode hay decode riêng.

---

## 6. Nếu muốn benchmark “chuẩn hơn”

Bạn có thể tách:

```go
b.Run("marshal", func(b *testing.B) { ... })
b.Run("unmarshal", func(b *testing.B) { ... })
```

Hoặc reset timer đúng chỗ:

```go
data, _ := json.Marshal(u)
b.ResetTimer()
for ...
```

---

## 7. Insight cho hệ thống banking (liên quan job của bạn)

Trong hệ thống fintech / SDK:

* Protobuf giúp giảm:

  * latency p99
  * GC pressure
  * bandwidth (quan trọng mobile SDK)

* JSON vẫn cần:

  * external API
  * debugging logs

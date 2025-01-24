# ใช้ base image ของ Golang ที่มีขนาดเล็ก
FROM golang:1.22-alpine

# กำหนด working directory ใน container
WORKDIR /usr/src/app

# คัดลอกโค้ดโปรเจกต์ทั้งหมด
COPY . .

# ตรวจสอบและ initialize Go module (ถ้ายังไม่มี)
RUN go mod init user-service || true

# ติดตั้ง dependencies
RUN go mod tidy
RUN go mod download

# สร้างไฟล์ binary ของแอปพลิเคชัน
RUN go build -o /usr/local/bin/user-service cmd/main.go

# เปิด port 8080 สำหรับแอปพลิเคชัน
EXPOSE 8080

# รันแอปพลิเคชัน
CMD ["/usr/local/bin/user-service"]

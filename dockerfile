# Gunakan base image Go
FROM golang:1.20

# Set direktori kerja dalam container
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Unduh dependensi Go
RUN go mod tidy

# Build aplikasi
RUN go build -o main .

# Jalankan aplikasi
CMD ["/app/main"]

# SIMPLE MICROSERVICE TODO LIST
Project Todo list backend sederhana untuk tugas RedHat OpenShift Training.
Simple Microservice ini menggunakan bahasa pemrograman golang, komunikasi antar service menggunakan
grpc.

## Tools yang digunakan
- go-micro
- golang migrate
- sqlc

## Service
### Auth
Service Auth menggunakan database postgresql berisi fungsi untuk melakukan autentikasi, meliputi:
- Login
- Verifikasi Token

### Todo
Service Todo menggukanan database mongodb merupakan inti dari fungsi aplikasi todo list, meliputi:
- Create
- Delete
- Done

### User
Service User menggunakan database postgresql yang mengatur data user, meliputi:
- Register
- Get User

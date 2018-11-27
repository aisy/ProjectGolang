# SATKKOM INDO Banking Web Service with golang
Project ini adalah back-end dari aplikasi React native mobile banking yang di buat oleh Programmer-programmer Satkomindo Satkom Indo. 

## Daftar Konten
- [SATKKOM INDO Banking Web Service with golang](#satkkom-indo-banking-web-service-with-golang)
   - [Daftar Konten](#daftar-konten)
   - [Kebutuhan](#kebutuhan)
   - [Instalasi Project](#instalasi-project)
      - [Set GOPATH](#set-gopath)
      - [Masuk ke Project GOPATH](#masuk-ke-project-gopath)
      - [Clone project](#clone-project)
      - [Masuk ke project](#masuk-ke-project)
      - [Memasang package untuk keperluan Project](#memasang-package-untuk-keperluan-project)
   - [Struktur Folder](#struktur-folder)
   - [Test apps](#test-apps)
   - [Build apps](#build-apps)
   - [Run Apss](#run-apss)
   - [Check Error](#check-error)


## Kebutuhan
ada beberapa kebutuhan yang di gunakan :

- [Golang](https://golang.org/)
- [Visual Studio Code](https://code.visualstudio.com/)
- [Extention Visual Studio untuk golang](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)
- [Postman](https://www.getpostman.com/)
- Web Browser


## Instalasi Project
Untuk langkah awakl silahkan download golang di halaman [resminya](https://golang.org/dl/), dan tersedia untuk beberapa OS.


### Set GOPATH
GOPATH ini adalah tempat kerja untuk mengerjakan project dengan GO. dan wajib menggunakan ini agar bisa dikelola oleh GO. Untuk OS Windows sendiri ke search->Tulis Environment->klik list pertama->pilih Environment Variable

Jika tidak ada Variable GOPATH maka buat terlebih dahulu dengan cara Klik New dan isi :
   - Variable name : GOPATH
   - Variable value : %USERPROFILE%\go
  
simpan dan jalankan aplikasi cmd untuk cek apakah GO sudah terinstall atau tidak dengan perintah :
``` cli
go
```


### Masuk ke Project GOPATH
Seperti yang sudah dibahas sebelumnya, project go harus ada di GOPATH, untuk bisa akses ke GOPATH masukkan perintah ini pada cmd :
```
cd %GOPATH%
```
di dalam GOPATH tersebut ada 3 folder yaitu :
- bin
- pkg
- src
  
Untuk folder project kita letakkan di dalam src :
```
cd src
```
Di dalam folder src ini ada dua folder yaitu :
- folder github.com/, ini adalah folder dimana package third party dimasukkan disni
- golang.org/, ini package bawaan golang


### Clone project
setelah masuk folder src, clone project ini :
```
https://gitlab.com/SATKOMINDO-Developers/backendgolang.git
```
Tunggu sampai proses clone selesai

### Masuk ke project
Setelah clone selesai, masuk ke folder project bernama `backendgolang` dan disinilah kita bisa berkreasi

### Memasang package untuk keperluan Project
-  [**gorilla/mux**](#gorilla/mux) 🦍
    package ini digunakan sebagai router http 
    untuk referensi ada [disini](https://github.com/gorilla/mux).
    untuk instalasi :
    ```
    go get github.com/gorilla/mux
    ```
- [**go-sql-driver**](#go-sql-driver)
    package ini digunakan untuk menghubungkan database mysql dengan go, untuk referensi bisa dilihat [disini](https://github.com/go-sql-driver/mysql).
    untuk instalasi :
    ```
    go get github.com/go-sql-driver/mysql
    ```
- Library atau modul lainnya bisa menyusul tergantung kebutuhan kedepan dan penggunaan lain yang sudah di sepakati bersama.

Happy Coding ❤❤🙌


## Struktur Folder
Untuk Sementara golang tidak memiliki struktur folder 😫😫

## Test apps
Untuk Testing cukup jalankan perintah ini :
```
go test
```

## Build apps
Untuk Build jalankan perintah ini
```
go build
```
maka akan menghasilkan file berformat .exe sesuai dengan nama project `backendgolang`

## Run Apss
Untuk menjalankan apps, ketik nama saja nama file dari hasil build tersebut :
```
backendgolang.exe
```
maka aplikasi akan berjalan di, dan untuk testingnya bisa menggunakan POSTMAN atau Browser secara langsung

## Check Error
Untuk check error, bisa saja menggunakan Test app, namun apabila ada error saat running bisa di cek di terminal.


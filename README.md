# Belajar Golang

## Requirement
Ketika memulai project, tambahkan beberapa package ini:

```bash
    go get gorm.io/gorm
    go get gorm.io/driver/postgres
    go get github.com/joho/godotenv
```

## Penjelasan Folder
- Project Root
  - configs /
  - internal /
    - controllers /
    - models /
  - readme.md
  - .gitignore
  - go.mod
  - go.sum
  - main.go


### `configs`

Memuat code yang bertujuan untuk persiapan dan konfigurasi segala sesuatu yang dibutuhkan project untuk beroperasi

### `internal`

Berisi rangkaian utama code yang digunakan dalam project. Biasanya masih terdapat beberapa folder lagi untuk mengakomodir keperluan struktur project yang diadopsi

### `controllers`

Berisi code yang mencakup fungsi fitur yang diakomodasi oleh program. Dalam folder ini, setiap fitur biasanya dibagi dalam file-file sesuai dengan konteks masing-masing. Contoh `user.go` akan memuat segala fungsi program yang berkaitan dengan data user misalnya login dan register.

### `models`

Berisi code yang bertugas untuk berhubungan dengan database. Code yang diletakkan dalam folder ini sangat spesifik digunakan untuk bersinggungan dengan database, segala sesuatau yang berhubungan dengan ORM dan database dilakukan pada folder ini. 

Keperluan code untuk beberapa konteks tertentu biasanya di pisahkan dalam bentuk file. Misalnya `user.go` pada folder `models` berisi code ORM yang mengarah ke database user.
# gin-simplebank

การตั้งค่าโปรเจกต์ Go สำหรับ Simple Bank
TODO: ขั้นตอนที่ 1: ติดตั้ง Migrate
# Migrate เป็นเครื่องมือที่ใช้ในการจัดการการย้ายข้อมูล (migrations) สำหรับฐานข้อมูลของคุณ

# ตั้งค่าตัวแปรเวอร์ชัน
using: version=v4.15.0

# ดาวน์โหลด Migrate
using: curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.linux-amd64.tar.gz -o migrate.tar.gz

# แตกไฟล์
using: tar xvzf migrate.tar.gz

# ย้ายไฟล์ Migrate ไปที่ /usr/local/bin
using: sudo mv migrate /usr/local/bin/migrate

# ตรวจสอบเวอร์ชัน
using: migrate -version

TODO: ขั้นตอนที่ 2: สร้างไฟล์ Migrate
# สร้างการย้ายข้อมูล (migration)
using: migrate create -ext sql -dir db/migration -seq init_schema

TODO: ขั้นตอนที่ 3: คำสั่งพื้นฐานของ Docker
# เข้าสู่คอนเทนเนอร์ของ PostgreSQL
using: docker exec -it go-postgres /bin/sh

# เข้าสู่ PostgreSQL
using: psql simple_bank

# ลบฐานข้อมูล simple_bank
using: dropdb simple_bank

# สร้างฐานข้อมูล simple_bank ใหม่
using: docker exec -it go-postgres createdb --username=root --owner=root simple_bank

# เข้าสู่ PostgreSQL ด้วยฐานข้อมูล simple_bank
using: docker exec -it go-postgres psql -U root simple_bank


TODO: ขั้นตอนที่ 4: ตั้งค่า SQLC

SQLC เป็นเครื่องมือที่ใช้ในการสร้างโค้ดจากการ query SQL

# เริ่มการตั้งค่า SQLC
using: sqlc init

TODO: ขั้นตอนที่ 5: สร้างไฟล์ go.mod

go.mod เป็นไฟล์ที่สำคัญในระบบการจัดการโมดูลของ Go (Go Modules) ซึ่งถูกนำมาใช้ในการจัดการการพึ่งพา (dependencies) ของโปรเจกต์ Go ไฟล์นี้จะเก็บข้อมูลเกี่ยวกับโมดูลของโปรเจกต์ รวมถึงรายการของแพคเกจภายนอกที่โปรเจกต์ของคุณต้องการใช้งานและเวอร์ชันที่เฉพาะเจาะจงของแต่ละแพคเกจเหล่านั้น

# สร้างไฟล์ go.mod
using: go mod init github.com/techschool/simple_bank

TODO: ขั้นตอนที่ 6: ทำความสะอาด Dependencies
# เพิ่ม dependencies ที่ขาดหาย
# ลบ dependencies ที่ไม่ใช้งาน
# อัปเดตไฟล์ go.sum
using: go mod tidy

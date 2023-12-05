# Aplikasi Enigma Laundry

Ini adalah aplikasi manajemen laundry sederhana yang dikembangkan menggunakan Go (Golang) dengan PostgreSQL sebagai basis data.

## Prasyarat

Sebelum menjalankan aplikasi, pastikan Anda telah menginstal hal-hal berikut:

- Go (Golang): [Panduan Instalasi](https://golang.org/doc/install)
- PostgreSQL: [Unduh](https://www.postgresql.org/download/)

## Menjalankan aplikasi 
1. Klon repository:
   
   git clone https://github.com/nurulaulia25/laundry-app.git
   cd laundry-app
   
## Jalankan aplikasi:
  go run main.go
  Perintah ini akan menjalankan aplikasi dan terhubung ke basis data PostgreSQL.
  
## Fungsi Aplikasi 

# Lihat Data Pelanggan 
Mengambil data pelanggan dari database dan menampilkan informasi pelanggan termasuk ID, nama, nomor telepon, tanggal masuk, tanggal keluar, dan tagihan.
customers := viewCustomers()
for _, customer := range customers {
    entryDateStr := customer.EntryDate.Format("2006-01-02")
    outDateStr := customer.OutDate.Format("2006-01-02")
    fmt.Printf("%d %s %s %s %s %.2f\n", customer.Id, customer.Name, customer.Phone, entryDateStr, outDateStr, customer.Bill)
}

# Memperbaharui Data Pelanggan
Memperbarui data pelanggan dengan ID tertentu atau menambahkan pelanggan baru jika ID tidak ada.
customer := entity.Customers{
    Id:        67890,
    Name:      "Cyntia",
    Phone:     "0897375863",
    EntryDate: time.Date(2022, 8, 28, 0, 0, 0, 0, time.Local),
    OutDate:   time.Date(2022, 8, 30, 0, 0, 0, 0, time.Local),
    Bill:      21000,
}
addUpdateCustomers(customer)

# Menghapus Pelanggan
Menghapus data pelanggan berdasarkan ID tertentu.
err := deleteCustomers("67890")
if err != nil {
    fmt.Println("Error:", err)
}

# Melihat Data Layanan
Menampilkan data layanan laundry termasuk ID, nama layanan, dan harga.
services := viewServices()
for _, service := range services {
    fmt.Printf("%d %s %.2f\n", service.Id, service.Name, service.Price)
}
# Memperbaharui Data Layanan
Memperbarui data layanan laundry dengan ID tertentu atau menambahkan layanan baru jika ID tidak ada.
service := entity.Services{
    Id:    4,
    Name:  "Cuci saja",
    Price: 5000,
}
addUpdateServices(service)

# Menghapus Layanan 
Menghapus layanan laundry berdasarkan ID tertentu.
err := deleteService("4")
if err != nil {
    fmt.Println("Error:", err)
}

# Melihat Data Transaksi 
Menampilkan data transaksi laundry termasuk ID pelanggan, ID layanan, jumlah, unit, tanggal transaksi, harga, dan total harga.
transactions := viewTransactions()
for _, transaction := range transactions {
    entryDate := transaction.DateEntry.Format("2006-01-02")
    fmt.Printf("%d %d %d %s %s %.2f %.2f\n", transaction.CustomerId, transaction.ServiceId, transaction.Quantity, transaction.Unit, entryDate, transaction.Price, transaction.TotalPrice)
}

# Melihat total Tagihan Pelanggan Bedasarkan Id
enghitung dan menampilkan total tagihan untuk seorang pelanggan berdasarkan ID pelanggan.
customerID := 12345
totalAmount, err := sumTotalPriceByCustomerId(customerID)
if err != nil {
    fmt.Println("Error fetching total amount:", err)
} else {
    fmt.Printf("Total amount for customer %d: %.2f\n", customerID, totalAmount)
}



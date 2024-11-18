package dbsharding

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"math/rand"

	"gorm.io/gorm"
)

func Start() {
	_, err := ConnectAndMigratePostgres()
	if err != nil {
		log.Fatal(err)
	}

	// prepareDataMaster(db)

	// StartWorker(db)

}

func StartWorker(db *gorm.DB) {
	const numWorkers = 20   // Jumlah worker
	const numJobs = 2000000 // Jumlah pekerjaan

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// WaitGroup untuk memastikan semua worker selesai
	var wg sync.WaitGroup

	// Membuat worker
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(db, w, jobs, results, &wg)
	}

	// Mengirimkan pekerjaan ke channel jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Menutup channel jobs, tidak ada pekerjaan lagi

	// Tunggu semua worker selesai
	wg.Wait()
	close(results) // Tutup channel results
}

func worker(db *gorm.DB, id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		msgErr := "Transaction success"
		err := insertDataOrder(db)
		if err != nil {
			msgErr = fmt.Sprintf("Transaction failed: %v", err)
		}
		fmt.Printf("worker %d - job: %d - %s\n", id, job, msgErr)

		results <- job
	}
}

func generateRandomTime() time.Time {
	start := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC)

	// Menghitung durasi antara start dan end
	duration := end.Sub(start)

	// Menghasilkan angka acak dalam rentang durasi
	randomDuration := time.Duration(rand.Int63n(int64(duration)))

	// Mengembalikan waktu acak dengan menambahkan randomDuration ke start
	return start.Add(randomDuration)
}

func insertDataOrder(db *gorm.DB) error {

	errTx := db.Transaction(func(tx *gorm.DB) error {
		// insert salesorder
		now := generateRandomTime()

		// get customer
		// var cust Customer
		custID := rand.Int63n(207) + 1
		// resultDB := db.First(&cust, "id = ?", custID)
		// if resultDB.Error != nil && resultDB.Error != gorm.ErrRecordNotFound {
		// 	log.Fatalf("Error retrieving customer: %v\n", resultDB.Error)
		// }

		var listSalesOrderItem []SalesOrderItem
		var totalAmount float64
		countProduct := rand.Intn(10) + 1
		for i := 0; i < countProduct; i++ {
			productID := rand.Int63n(190) + 1
			var product Product
			err := db.First(&product, "id = ?", productID).Error
			if err != nil && err != gorm.ErrRecordNotFound {
				return err
			}

			soItem := SalesOrderItem{
				ProductID:    productID,
				Quantity:     int64(rand.Intn(10) + 1),
				SellingPrice: product.Price,
			}
			listSalesOrderItem = append(listSalesOrderItem, soItem)
			totalAmount += product.Price * float64(soItem.Quantity)
		}
		// Format waktu hingga nanodetik (YYYYMMDDHHMMSSNNNNNNNNN)
		timeString := now.Format("20060102-150405.000000000")
		salesOrderNumber := strings.Replace(timeString, ".", "-", -1)
		so := SalesOrder{
			SalesOrderNumber: salesOrderNumber,
			SalesOrderDate:   now,
			CustomerID:       custID,
			TotalAmount:      totalAmount,
			Status:           "Pending",
			CreatedAt:        now,
			UpdatedAt:        now,
			CreatedBy:        "admin",
			UpdatedBy:        "admin",
		}

		err := tx.Create(&so).Error
		if err != nil {
			return err
		}

		for _, soItem := range listSalesOrderItem {
			soItem.SalesOrderID = so.ID
			err = tx.Create(&soItem).Error
			if err != nil {
				return err
			}
		}
		// Jika semua operasi berhasil, transaksi akan di-*commit* otomatis
		return nil
	})

	return errTx
}

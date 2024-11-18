package randomfolder

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func ConnectDB() error {
	// Membuka koneksi ke PostgreSQL
	connStr := "user=postgres password=mysecretpassword dbname=mydatabase sslmode=disable"
	var err error
	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	return nil
}
func insertRandomWord(job int) {
	// URL dari API
	apiURL := "https://random-word-api.herokuapp.com/word?number=1000"

	// Melakukan request ke API
	resp, err := http.Get(apiURL)
	if err != nil {
		log.Fatalf("Error ketika mengambil data dari API: %v", err)
	}

	// Membaca body dari response API
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error ketika membaca response body: %v", err)
	}
	resp.Body.Close()

	// Men-parse JSON dari API ke dalam slice struct APIResponse
	var listWord []string
	err = json.Unmarshal(body, &listWord)
	if err != nil {
		log.Fatalf("Error ketika melakukan unmarshal JSON: %v", err)
	}

	// Memasukkan data ke dalam tabel PostgreSQL
	for _, word := range listWord {
		insertQuery := `INSERT INTO random_table (name) VALUES ($1)`
		_, err := Db.Exec(insertQuery, word)
		if err != nil {
			log.Fatalf("Error ketika menyimpan data ke database: %v", err)
		}
	}

	fmt.Printf("%d - Data berhasil disimpan ke database!\n", job)
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		insertRandomWord(job)
		results <- job // Contoh hasil pekerjaan (misalnya job dikali 2)
	}
}

func StartWorker() {
	const numWorkers = 10 // Jumlah worker
	const numJobs = 1000  // Jumlah pekerjaan

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// WaitGroup untuk memastikan semua worker selesai
	var wg sync.WaitGroup

	// Membuat worker
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Mengirimkan pekerjaan ke channel jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Menutup channel jobs, tidak ada pekerjaan lagi

	// Tunggu semua worker selesai
	wg.Wait()
	close(results) // Tutup channel results

	// Mengambil hasil dari channel results
	// for result := range results {
	// 	fmt.Printf("Hasil: %d\n", result)
	// }
}

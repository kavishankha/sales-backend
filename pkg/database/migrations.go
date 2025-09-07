package database

//
//import "log"
//
//func RunMigrations() {
//	tables := []string{
//		`CREATE TABLE IF NOT EXISTS users (...);`,
//		`CREATE TABLE IF NOT EXISTS products (...);`,
//		`CREATE TABLE IF NOT EXISTS transactions (...);`,
//		`CREATE TABLE IF NOT EXISTS transaction_details (...);`,
//	}
//
//	for _, table := range tables {
//		if _, err := DB.Exec(table); err != nil {
//			log.Fatalf("Migration failed: %v", err)
//		}
//	}
//
//	log.Println("Migrations complete!")
//}

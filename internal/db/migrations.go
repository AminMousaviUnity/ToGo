package db


import (
	"log"
	"os"
	"os/exec"
)

func RunMigrations() {
	dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        log.Fatal("DATABASE_URL environment variable is not set")
    }

	cmd := exec.Command("migrate", "-database", dsn, "-path", "migrations", "up")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}

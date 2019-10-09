package migration

import (
	"fmt"
	"log"

	"gopi.com/gopi/storage"
	gormigrate "gopkg.in/gormigrate.v1"
)

// Migration class
type Migration struct{}

// getMigration create migration instance
func (m *Migration) getMigration() *gormigrate.Gormigrate {
	g := gormigrate.New(storage.GopiDB, gormigrate.DefaultOptions, []*gormigrate.Migration{
		Migration20181113233300,
		Migration20181113233301,
	})

	return g
}

// RunMigrations run migrations
func (m *Migration) RunMigrations() {
	fmt.Println("Preparing to migrate")

	g := m.getMigration()
	if err := g.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	fmt.Println("Migration completed successfully")
}

// RollbackLast rollback last migration
func (m *Migration) RollbackLast() {
	fmt.Println("Preparing to rollback")

	g := m.getMigration()
	if err := g.RollbackLast(); err != nil {
		log.Fatalf("Could not rollback: %v", err)
	}

	fmt.Println("Rollback completed successfully")
}

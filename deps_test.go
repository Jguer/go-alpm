package alpm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	alpm "github.com/Jguer/go-alpm/v2"
)

func TestDBList_FindSatisfier(t *testing.T) {
	h, err := alpm.Initialize("/", "/var/lib/pacman")
	defer h.Release()
	if err != nil {
		t.Fatalf("Failed to initialize alpm: %v", err)
	}

	dblist, err := h.SyncDBs()
	if err != nil {
		t.Fatalf("Failed to get sync DBs: %v", err)
	}

	// Dynamically pick a package name from the sync DBs
	var foundPkgName string
	for _, db := range dblist.Slice() {
		pkgs := db.PkgCache().Slice()
		if len(pkgs) > 0 {
			foundPkgName = pkgs[0].Name()
			break
		}
	}
	if foundPkgName == "" {
		t.Skip("No packages found in sync DBs")
	}

	// Test a dependency that should exist
	pkg, err := dblist.FindSatisfier(foundPkgName)
	assert.NoError(t, err)
	assert.NotNil(t, pkg)
	if pkg != nil {
		assert.Equal(t, foundPkgName, pkg.Name())
	}

	// Test a dependency that should not exist
	pkg, err = dblist.FindSatisfier("thispackagedoesnotexist>=1.0")
	assert.Error(t, err)
	assert.Nil(t, pkg)
}

func TestPackageList_FindSatisfier(t *testing.T) {
	h, err := alpm.Initialize("/", "/var/lib/pacman")
	defer h.Release()
	if err != nil {
		t.Fatalf("Failed to initialize alpm: %v", err)
	}

	db, err := h.LocalDB()
	if err != nil {
		t.Fatalf("Failed to get local DB: %v", err)
	}

	pkglist := db.PkgCache()

	// Test a dependency that should exist
	pkg, err := pkglist.FindSatisfier("glibc>=2.12")
	assert.NoError(t, err)
	assert.NotNil(t, pkg)
	if pkg != nil {
		assert.Equal(t, "glibc", pkg.Name())
	}

	// Test a dependency that should not exist
	pkg, err = pkglist.FindSatisfier("thispackagedoesnotexist>=1.0")
	assert.Error(t, err)
	assert.Nil(t, pkg)
}

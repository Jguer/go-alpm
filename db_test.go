package alpm

import (
	"os"
	"testing"
)

func getTestAlpmHandle(t *testing.T) *Handle {
	// Try to use environment variables for test root/dbpath, or fallback to system defaults
	root := os.Getenv("ALPM_TEST_ROOT")
	if root == "" {
		root = "/"
	}
	dbpath := os.Getenv("ALPM_TEST_DBPATH")
	if dbpath == "" {
		dbpath = "/var/lib/pacman/"
	}

	h, err := Initialize(root, dbpath)
	if err != nil {
		t.Skipf("Could not initialize alpm: %v", err)
	}
	return h
}

func TestHandle_RegisterSyncDB_and_SyncDBByName(t *testing.T) {
	h := getTestAlpmHandle(t)
	defer h.Release()

	dbName := "core" // This should exist on most Arch systems
	db, err := h.RegisterSyncDB(dbName, 0)
	if err != nil {
		t.Fatalf("RegisterSyncDB failed: %v", err)
	}
	if db == nil {
		t.Fatalf("RegisterSyncDB returned nil DB")
	}

	found, err := h.SyncDBByName(dbName)
	if err != nil {
		t.Fatalf("SyncDBByName failed: %v", err)
	}
	if found == nil {
		t.Fatalf("SyncDBByName returned nil DB")
	}
	if found.Name() != dbName {
		t.Errorf("Expected DB name %q, got %q", dbName, found.Name())
	}
}

func TestHandle_UnregisterAllSyncDBs(t *testing.T) {
	h := getTestAlpmHandle(t)
	defer h.Release()

	dbName := "core"
	_, err := h.RegisterSyncDB(dbName, 0)
	if err != nil {
		t.Fatalf("RegisterSyncDB failed: %v", err)
	}

	err = h.UnregisterAllSyncDBs()
	if err != nil {
		t.Fatalf("UnregisterAllSyncDBs failed: %v", err)
	}
}

// --- DB methods ---

func TestDB_Search(t *testing.T) {
	h := getTestAlpmHandle(t)
	defer h.Release()

	db, err := h.RegisterSyncDB("core", 0)
	if err != nil {
		t.Skipf("RegisterSyncDB failed: %v", err)
	}

	// Dynamically pick a package name from the sync DB
	pkgs := db.PkgCache().Slice()
	if len(pkgs) == 0 {
		t.Skip("No packages found in the sync DB; skipping test")
	}
	target := pkgs[0].Name()

	searchResults := db.Search([]string{target})
	if searchResults == nil {
		t.Fatalf("Search returned nil for target %q", target)
	}
	slice := searchResults.Slice()
	if len(slice) == 0 {
		t.Errorf("Expected at least one package for target %q, got 0", target)
	}
	found := false
	for _, pkg := range slice {
		if pkg != nil && pkg.Name() == target {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Did not find package with name %q in search results", target)
	}

	// Search for a non-existent package
	nonexistent := "thispackagedoesnotexist12345"
	none := db.Search([]string{nonexistent})
	if none == nil {
		t.Fatalf("Search returned nil for non-existent target")
	}
	if len(none.Slice()) != 0 {
		t.Errorf("Expected 0 results for non-existent package, got %d", len(none.Slice()))
	}
}

func TestDB_AddServer(t *testing.T) {
	h := getTestAlpmHandle(t)
	defer h.Release()

	db, err := h.RegisterSyncDB("core", 0)
	if err != nil {
		t.Skipf("RegisterSyncDB failed: %v", err)
	}

	server := "http://example.com/$repo/os/$arch"
	db.AddServer(server)

	servers := db.Servers()
	found := false
	for _, s := range servers {
		if s == server {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Server %q not found in DB server list: %v", server, servers)
	}
}

// --- DBList group method ---

func TestDBList_Slice_Empty(t *testing.T) {
	var dblist DBList // nil list
	slice := dblist.Slice()
	if len(slice) != 0 {
		t.Errorf("Expected empty slice, got %d elements", len(slice))
	}
}

func TestDBList_ForEach_and_Slice_NonEmpty(t *testing.T) {
	h := getTestAlpmHandle(t)
	defer h.Release()

	dblist, err := h.SyncDBs()
	if err != nil {
		t.Fatalf("SyncDBs failed: %v", err)
	}

	var count int
	names := make(map[string]bool)
	err = dblist.ForEach(func(db IDB) error {
		count++
		name := db.Name()
		names[name] = true
		return nil
	})
	if err != nil {
		t.Fatalf("ForEach failed: %v", err)
	}

	slice := dblist.Slice()
	if len(slice) != count {
		t.Errorf("Slice length %d does not match ForEach count %d", len(slice), count)
	}
	for _, db := range slice {
		if !names[db.Name()] {
			t.Errorf("DB %q in slice not found in ForEach names", db.Name())
		}
	}
}

func TestDBList_Append(t *testing.T) {
	h := getTestAlpmHandle(t)
	defer h.Release()

	db, err := h.RegisterSyncDB("core", 0)
	if err != nil {
		t.Skipf("RegisterSyncDB failed: %v", err)
	}

	dblist := h.NewDBList()
	if dblist == nil {
		t.Fatalf("NewDBList returned nil")
	}

	// Should be empty initially
	slice := dblist.Slice()
	if len(slice) != 0 {
		t.Errorf("Expected empty slice, got %d elements", len(slice))
	}

	dblist.Append(db)
	slice = dblist.Slice()
	if len(slice) != 1 {
		t.Errorf("Expected slice of length 1 after append, got %d", len(slice))
	}
	if slice[0].Name() != "core" {
		t.Errorf("Expected DB name 'core', got %q", slice[0].Name())
	}
}

func TestDBList_FindGroupPkgs(t *testing.T) {
	h := getTestAlpmHandle(t)
	defer h.Release()

	dblist, err := h.SyncDBs()
	if err != nil {
		t.Fatalf("SyncDBs failed: %v", err)
	}

	group := "base" // Common group in Arch
	pkgs := dblist.FindGroupPkgs(group)
	if pkgs == nil {
		t.Fatalf("FindGroupPkgs returned nil for group %q", group)
	}
	// We can't guarantee the group exists, but we can check that the call succeeded
}

func TestHandle_SyncDBListByDBName(t *testing.T) {
	h := getTestAlpmHandle(t)
	defer h.Release()

	dbName := "core"
	_, err := h.RegisterSyncDB(dbName, 0)
	if err != nil {
		t.Skipf("RegisterSyncDB failed: %v", err)
	}

	dblist, err := h.SyncDBListByDBName(dbName)
	if err != nil {
		t.Fatalf("SyncDBListByDBName failed: %v", err)
	}
	if dblist == nil {
		t.Fatalf("SyncDBListByDBName returned nil list")
	}
	slice := dblist.Slice()
	if len(slice) != 1 {
		t.Errorf("Expected list of length 1, got %d", len(slice))
	}
	if slice[0].Name() != dbName {
		t.Errorf("Expected DB name %q, got %q", dbName, slice[0].Name())
	}

	// Negative case
	_, err = h.SyncDBListByDBName("doesnotexist123")
	if err == nil {
		t.Errorf("Expected error for non-existent DB name, got nil")
	}
}

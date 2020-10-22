package alpm

import (
	"time"
)

// IPackage is an interface type for alpm.Package
type IPackage interface {
	FileName() string
	Base() string
	Base64Signature() string
	Validation() Validation
	// Architecture returns the package target Architecture.
	Architecture() string
	// Backup returns a list of package backups.
	Backup() BackupList
	// BuildDate returns the BuildDate of the package.
	BuildDate() time.Time
	// Conflicts returns the conflicts of the package as a DependList.
	Conflicts() DependList
	// DB returns the package's origin database.
	DB() IDB
	// Depends returns the package's dependency list.
	Depends() DependList
	// Depends returns the package's optional dependency list.
	OptionalDepends() DependList
	// Depends returns the package's check dependency list.
	CheckDepends() DependList
	// Depends returns the package's make dependency list.
	MakeDepends() DependList
	// Description returns the package's description.
	Description() string
	// Files returns the file list of the package.
	Files() []File
	// ContainsFile checks if the path is in the package filelist
	ContainsFile(path string) (File, error)
	// Groups returns the groups the package belongs to.
	Groups() StringList
	// ISize returns the package installed size.
	ISize() int64
	// InstallDate returns the package install date.
	InstallDate() time.Time
	// Licenses returns the package license list.
	Licenses() StringList
	// SHA256Sum returns package SHA256Sum.
	SHA256Sum() string
	// MD5Sum returns package MD5Sum.
	MD5Sum() string
	// Name returns package name.
	Name() string
	// Packager returns package packager name.
	Packager() string
	// Provides returns DependList of packages provides by package.
	Provides() DependList
	// Reason returns package install reason.
	Reason() PkgReason
	// Origin returns package origin.
	Origin() PkgFrom
	// Replaces returns a DependList with the packages this package replaces.
	Replaces() DependList
	// Size returns the packed package size.
	Size() int64
	// URL returns the upstream URL of the package.
	URL() string
	// Version returns the package version.
	Version() string
	// ComputeRequiredBy returns the names of reverse dependencies of a package
	ComputeRequiredBy() []string
	// ComputeOptionalFor returns the names of packages that optionally
	// require the given package
	ComputeOptionalFor() []string
	ShouldIgnore() bool

	// SyncNewVersion checks if there is a new version of the
	// package in a given DBlist.
	SyncNewVersion(l IDBList) IPackage

	Type() string
}

// IPackageList exports the alpm.PackageList symbols
type IPackageList interface {
	// ForEach executes an action on each package of the PackageList.
	ForEach(func(IPackage) error) error
	// Slice converts the PackageList to a Package Slice.
	Slice() []IPackage
	// SortBySize returns a PackageList sorted by size.
	SortBySize() IPackageList
	// FindSatisfier finds a package that satisfies depstring from PkgList
	FindSatisfier(string) (IPackage, error)
}

// IDB is an interface type for alpm.DB
type IDB interface {
	Unregister() error
	// Name returns name of the db
	Name() string
	// Servers returns host server URL.
	Servers() []string
	// SetServers sets server list to use.
	SetServers(servers []string)
	// AddServers adds a string to the server list.
	AddServer(server string)
	// SetUsage sets the Usage of the database
	SetUsage(usage Usage)
	// Name searches a package in db.
	Pkg(name string) IPackage
	// PkgCache returns the list of packages of the database
	PkgCache() IPackageList
	Search([]string) IPackageList
}

// IDBList interfaces alpm.DBList
type IDBList interface {
	// ForEach executes an action on each DB.
	ForEach(func(IDB) error) error
	// Slice converts DB list to DB slice.
	Slice() []IDB
	// PkgCachebyGroup returns a PackageList of packages belonging to a group
	FindGroupPkgs(string) IPackageList
	// FindSatisfier searches a DBList for a package that satisfies depstring
	// Example "glibc>=2.12"
	FindSatisfier(string) (IPackage, error)
}

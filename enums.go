// enums.go - libaplm enumerations.
//
// Copyright (c) 2013 The go-alpm Authors
//
// MIT Licensed. See LICENSE for details.

package alpm

// Install reason of a package.
type PkgReason uint

const (
	PkgReasonExplicit PkgReason = 0
	PkgReasonDepend   PkgReason = 1
)

func (r PkgReason) String() string {
	switch r {
	case PkgReasonExplicit:
		return "Explicitly installed"
	case PkgReasonDepend:
		return "Installed as a dependency of another package"
	}
	return ""
}

// Source of a package structure.
type PkgFrom uint

const (
	FromFile PkgFrom = iota + 1
	FromLocalDB
	FromSyncDB
)

// Dependency constraint types.
type DepMod uint

const (
	DepModAny DepMod = iota + 1 // Any version.
	DepModEq                    // Specific version.
	DepModGE                    // Test for >= some version.
	DepModLE                    // Test for <= some version.
	DepModGT                    // Test for > some version.
	DepModLT                    // Test for < some version.
)

func (mod DepMod) String() string {
	switch mod {
	case DepModEq:
		return "="
	case DepModGE:
		return ">="
	case DepModLE:
		return "<="
	case DepModGT:
		return ">"
	case DepModLT:
		return "<"
	}
	return ""
}

// Signature checking level.
type SigLevel int

const (
	SigPackage SigLevel = 1 << iota
	SigPackageOptional
	SigPackageMarginalOk
	SigPackageUnknownOk
)
const (
	SigDatabase SigLevel = 1 << (10 + iota)
	SigDatabaseOptional
	SigDatabaseMarginalOk
	SigDatabaseUnknownOk
)
const SigUseDefault SigLevel = 1 << 30

// Signature status
type SigStatus int

const (
	SigStatusValid SigStatus = iota
	SigStatusKeyExpired
	SigStatusSigExpired
	SigStatusKeyUnknown
	SigStatusKeyDisabled
)

// Logging levels.
const (
	LogError uint16 = 1 << iota
	LogWarning
	LogDebug
	LogFunction
)

type QuestionType uint

const (
	QuestionTypeInstallIgnorepkg QuestionType = 1 << iota
	QuestionTypeReplacePkg
	QuestionTypeConflictPkg
	QuestionTypeCorruptedPkg
	QuestionTypeRemovePkgs
	QuestionTypeSelectProvider
	QuestionTypeImportKey
)

type Validation int

const (
	ValidationNone Validation = 1 << iota
	ValidationMD5Sum
	ValidationSHA256Sum
	ValidationSignature
	ValidationUnkown Validation = 0
)

type Usage int

const (
	UsageSync Usage = 1 << iota
	UsageSearch
	UsageInstall
	UsageUpgrade
	UsageAll = (1 << 4) - 1
)

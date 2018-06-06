// Package srcinfo is a parser for srcinfo files. Typically generated by
// makepkg, part of the pacman package manager.
//
// Split packages and architecture dependent fields are fully supported.
//
// This Package aimes to parse srcinfos but not interpret them in any way.
// All values are fundamentally strings, other tools should be used for
// things such as dependency parsing, validity checking etc.
package srcinfo

import (
	"fmt"
)

// ArchString describes string values that may be architecture dependent.
// For Example depends_x86_64.
// If Arch is an empty string then the field is not architecture dependent.
type ArchString struct {
	Arch  string // Architecture name
	Value string // Value
}

// Package describes the fields of a pkgbuild that may be overwritten by
// in build_<pkgname> function.
type Package struct {
	Pkgdesc    string
	Arch       []string
	URL        string
	License    []string
	Groups     []string
	Depends    []ArchString
	OptDepends []ArchString
	Provides   []ArchString
	Conflicts  []ArchString
	Replaces   []ArchString
	Backup     []string
	Options    []string
	Install    string
	Changelog  string
}

// PackageBase describes the fields of a pkgbuild that may not be overwritten
// in package_<pkgname> function.
type PackageBase struct {
	Pkgbase      string
	Pkgnames     []string
	Pkgver       string
	Pkgrel       string
	Epoch        string
	Source       []ArchString
	ValidPGPKeys []string
	NoExtract    []string
	MD5Sums      []ArchString
	SHA1Sums     []ArchString
	SHA224Sums   []ArchString
	SHA256Sums   []ArchString
	SHA384Sums   []ArchString
	SHA512Sums   []ArchString
	MakeDepends  []ArchString
	CheckDepends []ArchString
}

// Srcinfo represents a full srcinfo. All global fields are defined here while
// fields overwritten in the package_<pkgname> function are defined in the
// Packages field.
//
// Note: The Packages field only contains the values that each package
// overrides, global fields will be missing. A Package containing both global
// and overwritten fields can be generated using the SplitPackage function.
type Srcinfo struct {
	PackageBase
	Package
	Packages []Package
}

// SplitPackage generates a Package that contains all fields that the specified
// pkgname has. But will fall back on global fields if they are not defined in
// the Package.
//
// Note slice values will be passed by reference, it is not recommended you
// modify this struct after it is returned.
func (si *Srcinfo) SplitPackage(pkgname string) (*Package, error) {
	for n, name := range si.Pkgnames {
		if name == pkgname {
			return mergeSplitPackage(&si.Package, &si.Packages[n]), nil
		}
	}

	return nil, fmt.Errorf("Package \"%s\" is not part of this package base", pkgname)
}

func mergeSplitPackage(base, split *Package) *Package {
	pkg := &Package{}
	*pkg = *base

	if split.Pkgdesc != "" {
		pkg.Pkgdesc = split.Pkgdesc
	}

	if len(split.Arch) != 0 {
		pkg.Arch = split.Arch
	}

	if split.URL != "" {
		pkg.URL = split.URL
	}

	if len(split.License) != 0 {
		pkg.License = split.License
	}

	if len(split.Groups) != 0 {
		pkg.Groups = split.Groups
	}

	if len(split.License) != 0 {
		pkg.License = split.License
	}

	if len(split.Depends) != 0 {
		pkg.Depends = split.Depends
	}

	if len(split.OptDepends) != 0 {
		pkg.OptDepends = split.OptDepends
	}

	if len(split.Provides) != 0 {
		pkg.Provides = split.Provides
	}

	if len(split.Conflicts) != 0 {
		pkg.Conflicts = split.Conflicts
	}

	if len(split.Replaces) != 0 {
		pkg.Replaces = split.Replaces
	}

	if len(split.Backup) != 0 {
		pkg.Backup = split.Backup
	}

	if len(split.Options) != 0 {
		pkg.Options = split.Options
	}

	if split.Changelog != "" {
		pkg.Changelog = split.Changelog
	}

	if split.Install != "" {
		pkg.Install = split.Install
	}

	return pkg
}

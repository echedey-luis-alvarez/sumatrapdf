package main

import (
	"fmt"
	"path/filepath"
)

func buildPreRelease() {
	detectSigntoolPath() // early exit if missing

	ver := getVerForBuildType(buildTypePreRel)
	s := fmt.Sprintf("buidling pre-release version %s", ver)
	defer makePrintDuration(s)()

	verifyGitCleanMust()
	verifyOnMasterBranchMust()
	verifyTranslationsMust()

	clean()
	setBuildConfigPreRelease()
	defer revertBuildConfig()

	build(rel32Dir, "Release", "Win32")
	nameInZip := fmt.Sprintf("SumatraPDF-prerel-%s-32.exe", ver)
	createExeZipWithGoWithNameMust(rel32Dir, nameInZip)

	build(rel64Dir, "Release", "x64")
	nameInZip = fmt.Sprintf("SumatraPDF-prerel-%s-64.exe", ver)
	createExeZipWithGoWithNameMust(rel64Dir, nameInZip)

	createManifestMust()

	dstDir := filepath.Join("out", "final-prerel")
	prefix := fmt.Sprintf("SumatraPDF-prerel-%s", ver)
	copyBuiltFiles(dstDir, rel32Dir, prefix)
	copyBuiltFiles(dstDir, rel64Dir, prefix+"-64")
	copyBuiltManifest(dstDir, prefix)
}

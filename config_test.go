package alpm

import (
	"testing"

	"github.com/Morganamilo/go-pacmanconf"
	"github.com/stretchr/testify/assert"
)

type AlpmExecutor struct {
	handle       *Handle
	localDB      IDB
	syncDB       IDBList
	syncDBsCache []IDB
	conf         *pacmanconf.Config
}

func (ae *AlpmExecutor) RefreshHandle() error {
	if ae.handle != nil {
		if errRelease := ae.handle.Release(); errRelease != nil {
			return errRelease
		}
	}

	alpmHandle, err := Initialize(ae.conf.RootDir, ae.conf.DBPath)
	if err != nil {
		return err
	}

	if errConf := configureAlpm(ae.conf, alpmHandle); errConf != nil {
		return errConf
	}

	ae.handle = alpmHandle
	ae.syncDBsCache = nil

	ae.syncDB, err = alpmHandle.SyncDBs()
	if err != nil {
		return err
	}

	ae.localDB, err = alpmHandle.LocalDB()

	return err
}

func toUsage(usages []string) Usage {
	if len(usages) == 0 {
		return UsageAll
	}

	var ret Usage

	for _, usage := range usages {
		switch usage {
		case "Sync":
			ret |= UsageSync
		case "Search":
			ret |= UsageSearch
		case "Install":
			ret |= UsageInstall
		case "Upgrade":
			ret |= UsageUpgrade
		case "All":
			ret |= UsageAll
		}
	}

	return ret
}

func configureAlpm(pacmanConf *pacmanconf.Config, alpmHandle *Handle) error {
	for _, repo := range pacmanConf.Repos {
		// TODO: set SigLevel
		alpmDB, err := alpmHandle.RegisterSyncDB(repo.Name, 0)
		if err != nil {
			return err
		}

		alpmDB.SetServers(repo.Servers)
		alpmDB.SetUsage(toUsage(repo.Usage))
	}

	if err := alpmHandle.SetCacheDirs(pacmanConf.CacheDir); err != nil {
		return err
	}

	// add hook directories 1-by-1 to avoid overwriting the system directory
	for _, dir := range pacmanConf.HookDir {
		if err := alpmHandle.AddHookDir(dir); err != nil {
			return err
		}
	}

	if err := alpmHandle.SetGPGDir(pacmanConf.GPGDir); err != nil {
		return err
	}

	if err := alpmHandle.SetLogFile(pacmanConf.LogFile); err != nil {
		return err
	}

	if err := alpmHandle.SetIgnorePkgs(pacmanConf.IgnorePkg); err != nil {
		return err
	}

	if err := alpmHandle.SetIgnoreGroups(pacmanConf.IgnoreGroup); err != nil {
		return err
	}

	if err := alpmHandle.SetArchitectures(pacmanConf.Architecture); err != nil {
		return err
	}

	if err := alpmHandle.SetNoUpgrades(pacmanConf.NoUpgrade); err != nil {
		return err
	}

	if err := alpmHandle.SetNoExtracts(pacmanConf.NoExtract); err != nil {
		return err
	}

	if err := alpmHandle.SetUseSyslog(pacmanConf.UseSyslog); err != nil {
		return err
	}

	return alpmHandle.SetCheckSpace(pacmanConf.CheckSpace)
}

func NewExecutor(pacmanConf *pacmanconf.Config) (*AlpmExecutor, error) {
	ae := &AlpmExecutor{conf: pacmanConf}

	err := ae.RefreshHandle()
	if err != nil {
		return nil, err
	}

	ae.localDB, err = ae.handle.LocalDB()
	if err != nil {
		return nil, err
	}

	ae.syncDB, err = ae.handle.SyncDBs()
	if err != nil {
		return nil, err
	}

	return ae, nil
}

func TestAlpmExecutor(t *testing.T) {
	t.Parallel()
	pacmanConf := &pacmanconf.Config{
		RootDir:                "/",
		DBPath:                 "/var/lib/pacman/",
		CacheDir:               []string{"/cachedir/", "/another/"},
		HookDir:                []string{"/hookdir/"},
		GPGDir:                 "/gpgdir/",
		LogFile:                "/logfile",
		HoldPkg:                []string(nil),
		IgnorePkg:              []string{"ignore", "this", "package"},
		IgnoreGroup:            []string{"ignore", "this", "group"},
		Architecture:           []string{"8086"},
		XferCommand:            "",
		NoUpgrade:              []string{"noupgrade"},
		NoExtract:              []string{"noextract"},
		CleanMethod:            []string{"KeepInstalled"},
		SigLevel:               []string{"PackageOptional", "PackageTrustedOnly", "DatabaseOptional", "DatabaseTrustedOnly"},
		LocalFileSigLevel:      []string(nil),
		RemoteFileSigLevel:     []string(nil),
		UseSyslog:              false,
		Color:                  false,
		UseDelta:               0,
		TotalDownload:          true,
		CheckSpace:             true,
		VerbosePkgLists:        true,
		DisableDownloadTimeout: false,
		Repos: []pacmanconf.Repository{
			{Name: "repo1", Servers: []string{"repo1"}, SigLevel: []string(nil), Usage: []string{"All"}},
			{Name: "repo2", Servers: []string{"repo2"}, SigLevel: []string(nil), Usage: []string{"All"}},
		},
	}

	aExec, err := NewExecutor(pacmanConf)
	assert.NoError(t, err)

	assert.NotNil(t, aExec.conf)
	assert.EqualValues(t, pacmanConf, aExec.conf)

	assert.NotNil(t, aExec.localDB)
	assert.NotNil(t, aExec.syncDB)
	h := aExec.handle
	assert.NotNil(t, h)

	root, err := h.Root()
	assert.Nil(t, err)
	assert.Equal(t, "/", root)

	dbPath, err := h.DBPath()
	assert.Nil(t, err)
	assert.Equal(t, "/var/lib/pacman/", dbPath)

	cache, err := h.CacheDirs()
	assert.Nil(t, err)
	assert.Equal(t, []string{"/cachedir/", "/another/"}, cache.Slice())

	log, err := h.LogFile()
	assert.Nil(t, err)
	assert.Equal(t, "/logfile", log)

	gpg, err := h.GPGDir()
	assert.Nil(t, err)
	assert.Equal(t, "/gpgdir/", gpg)

	hook, err := h.HookDirs()
	assert.Nil(t, err)
	assert.Equal(t, []string{"/usr/share/libalpm/hooks/", "/hookdir/"}, hook.Slice())

	arch, err := alpmTestGetArch(h)
	assert.Nil(t, err)
	assert.Equal(t, []string{"8086"}, arch)

	ignorePkg, err := h.IgnorePkgs()
	assert.Nil(t, err)
	assert.Equal(t, []string{"ignore", "this", "package"}, ignorePkg.Slice())

	ignoreGroup, err := h.IgnoreGroups()
	assert.Nil(t, err)
	assert.Equal(t, []string{"ignore", "this", "group"}, ignoreGroup.Slice())

	noUp, err := h.NoUpgrades()
	assert.Nil(t, err)
	assert.Equal(t, []string{"noupgrade"}, noUp.Slice())

	noEx, err := h.NoExtracts()
	assert.Nil(t, err)
	assert.Equal(t, []string{"noextract"}, noEx.Slice())

	check, err := h.CheckSpace()
	assert.Nil(t, err)
	assert.Equal(t, true, check)
}

func alpmTestGetArch(h *Handle) ([]string, error) {
	architectures, err := h.GetArchitectures()

	return architectures.Slice(), err
}

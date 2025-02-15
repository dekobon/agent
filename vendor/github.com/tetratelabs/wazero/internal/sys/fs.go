package sys

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"syscall"
	"time"

	"github.com/tetratelabs/wazero/internal/descriptor"
	"github.com/tetratelabs/wazero/internal/platform"
	"github.com/tetratelabs/wazero/internal/sysfs"
)

const (
	FdStdin int32 = iota
	FdStdout
	FdStderr
	// FdPreopen is the file descriptor of the first pre-opened directory.
	//
	// # Why file descriptor 3?
	//
	// While not specified, the most common WASI implementation, wasi-libc,
	// expects POSIX style file descriptor allocation, where the lowest
	// available number is used to open the next file. Since 1 and 2 are taken
	// by stdout and stderr, the next is 3.
	//   - https://github.com/WebAssembly/WASI/issues/122
	//   - https://pubs.opengroup.org/onlinepubs/9699919799/functions/V2_chap02.html#tag_15_14
	//   - https://github.com/WebAssembly/wasi-libc/blob/wasi-sdk-16/libc-bottom-half/sources/preopens.c#L215
	FdPreopen
)

const modeDevice = uint32(fs.ModeDevice | 0o640)

type stdioFileWriter struct {
	w io.Writer
	s fs.FileInfo
}

// Stat implements fs.File
func (w *stdioFileWriter) Stat() (fs.FileInfo, error) { return w.s, nil }

// Read implements fs.File
func (w *stdioFileWriter) Read([]byte) (n int, err error) {
	return // emulate os.Stdout which returns zero
}

// Write implements io.Writer
func (w *stdioFileWriter) Write(p []byte) (n int, err error) {
	return w.w.Write(p)
}

// Close implements fs.File
func (w *stdioFileWriter) Close() error {
	// Don't actually close the underlying file, as we didn't open it!
	return nil
}

// StdioFilePoller is a strategy for polling a StdioFileReader for a given duration.
// It returns true if the reader has data ready to be read, false and/or an error otherwise.
type StdioFilePoller interface {
	Poll(duration time.Duration) (bool, error)
}

// PollerDefaultStdin is a poller that checks standard input.
var PollerDefaultStdin = &pollerDefaultStdin{}

type pollerDefaultStdin struct{}

// Poll implements StdioFilePoller for pollerDefaultStdin.
func (*pollerDefaultStdin) Poll(duration time.Duration) (bool, error) {
	fdSet := platform.FdSet{}
	fdSet.Set(int(FdStdin))
	count, err := platform.Select(int(FdStdin+1), &fdSet, nil, nil, &duration)
	return count > 0, err
}

// PollerAlwaysReady is a poller that ignores the given timeout, and it returns true and no error.
var PollerAlwaysReady = &pollerAlwaysReady{}

type pollerAlwaysReady struct{}

// Poll implements StdioFilePoller for pollerAlwaysReady.
func (*pollerAlwaysReady) Poll(time.Duration) (bool, error) { return true, nil }

// PollerNeverReady is a poller that waits for the given duration, and it always returns false and no error.
var PollerNeverReady = &pollerNeverReady{}

type pollerNeverReady struct{}

// Poll implements StdioFilePoller for pollerNeverReady.
func (*pollerNeverReady) Poll(d time.Duration) (bool, error) { time.Sleep(d); return false, nil }

// StdioFileReader implements io.Reader for stdio files.
type StdioFileReader struct {
	r    io.Reader
	s    fs.FileInfo
	poll StdioFilePoller
}

// NewStdioFileReader is a constructor for StdioFileReader.
func NewStdioFileReader(reader io.Reader, fileInfo fs.FileInfo, poll StdioFilePoller) *StdioFileReader {
	return &StdioFileReader{
		r:    reader,
		s:    fileInfo,
		poll: poll,
	}
}

// Poll invokes the StdioFilePoller that was given at the NewStdioFileReader constructor.
func (r *StdioFileReader) Poll(duration time.Duration) (bool, error) {
	return r.poll.Poll(duration)
}

// Stat implements fs.File
func (r *StdioFileReader) Stat() (fs.FileInfo, error) { return r.s, nil }

// Read implements fs.File
func (r *StdioFileReader) Read(p []byte) (n int, err error) {
	return r.r.Read(p)
}

// Close implements fs.File
func (r *StdioFileReader) Close() error {
	// Don't actually close the underlying file, as we didn't open it!
	return nil
}

var (
	noopStdinStat  = stdioFileInfo{0, modeDevice}
	noopStdoutStat = stdioFileInfo{1, modeDevice}
	noopStderrStat = stdioFileInfo{2, modeDevice}
)

// stdioFileInfo implements fs.FileInfo where index zero is the FD and one is the mode.
type stdioFileInfo [2]uint32

func (s stdioFileInfo) Name() string {
	switch s[0] {
	case 0:
		return "stdin"
	case 1:
		return "stdout"
	case 2:
		return "stderr"
	default:
		panic(fmt.Errorf("BUG: incorrect FD %d", s[0]))
	}
}

func (stdioFileInfo) Size() int64         { return 0 }
func (s stdioFileInfo) Mode() fs.FileMode { return fs.FileMode(s[1]) }
func (stdioFileInfo) ModTime() time.Time  { return time.Unix(0, 0) }
func (stdioFileInfo) IsDir() bool         { return false }
func (stdioFileInfo) Sys() interface{}    { return nil }

type lazyDir struct {
	fs sysfs.FS
	f  fs.File
}

// Stat implements fs.File
func (r *lazyDir) Stat() (fs.FileInfo, error) {
	if f, err := r.file(); err != 0 {
		return nil, err
	} else {
		return f.Stat()
	}
}

func (r *lazyDir) file() (f fs.File, errno syscall.Errno) {
	if f = r.f; r.f != nil {
		return
	}
	r.f, errno = r.fs.OpenFile(".", os.O_RDONLY, 0)
	f = r.f
	return
}

// Read implements fs.File
func (r *lazyDir) Read(p []byte) (n int, err error) {
	if f, errno := r.file(); errno != 0 {
		return 0, errno
	} else {
		return f.Read(p)
	}
}

// Close implements fs.File
func (r *lazyDir) Close() error {
	f := r.f
	if f == nil {
		return nil // never opened
	}
	return f.Close()
}

// FileEntry maps a path to an open file in a file system.
type FileEntry struct {
	// Name is the name of the directory up to its pre-open, or the pre-open
	// name itself when IsPreopen.
	//
	// Note: This can drift on rename.
	Name string

	// IsPreopen is a directory that is lazily opened.
	IsPreopen bool

	// FS is the filesystem associated with the pre-open.
	FS sysfs.FS

	// cachedStat includes fields that won't change while a file is open.
	cachedStat *cachedStat

	// File is always non-nil.
	File fs.File

	// ReadDir is present when this File is a fs.ReadDirFile and `ReadDir`
	// was called.
	ReadDir *ReadDir

	openPath string
	openFlag int
	openPerm fs.FileMode
}

type cachedStat struct {
	// Ino is the file serial number, or zero if not available.
	Ino uint64

	// Type is the same as what's documented on platform.Dirent.
	Type fs.FileMode
}

// CachedStat returns the cacheable parts of platform.Stat_t or an error if
// they couldn't be retrieved.
func (f *FileEntry) CachedStat() (ino uint64, fileType fs.FileMode, err error) {
	if f.cachedStat == nil {
		if _, err = f.Stat(); err != nil {
			return
		}
	}
	return f.cachedStat.Ino, f.cachedStat.Type, nil
}

// Stat returns the underlying stat of this file.
func (f *FileEntry) Stat() (st platform.Stat_t, err error) {
	var errno syscall.Errno
	if ld, ok := f.File.(*lazyDir); ok {
		var sf fs.File
		if sf, errno = ld.file(); errno == 0 {
			st, errno = platform.StatFile(sf)
		}
	} else {
		st, errno = platform.StatFile(f.File)
	}

	if errno != 0 {
		err = errno
	} else {
		f.cachedStat = &cachedStat{Ino: st.Ino, Type: st.Mode & fs.ModeType}
	}
	return
}

// ReadDir is the status of a prior fs.ReadDirFile call.
type ReadDir struct {
	// CountRead is the total count of files read including Dirents.
	CountRead uint64

	// Dirents is the contents of the last platform.Readdir call. Notably,
	// directory listing are not rewindable, so we keep entries around in case
	// the caller mis-estimated their buffer and needs a few still cached.
	//
	// Note: This is wasi-specific and needs to be refactored.
	// In wasi preview1, dot and dot-dot entries are required to exist, but the
	// reverse is true for preview2. More importantly, preview2 holds separate
	// stateful dir-entry-streams per file.
	Dirents []*platform.Dirent
}

type FSContext struct {
	// rootFS is the root ("/") mount.
	rootFS sysfs.FS

	// openedFiles is a map of file descriptor numbers (>=FdPreopen) to open files
	// (or directories) and defaults to empty.
	// TODO: This is unguarded, so not goroutine-safe!
	openedFiles FileTable
}

// FileTable is a specialization of the descriptor.Table type used to map file
// descriptors to file entries.
type FileTable = descriptor.Table[int32, *FileEntry]

// NewFSContext creates a FSContext with stdio streams and an optional
// pre-opened filesystem.
//
// If `preopened` is not sysfs.UnimplementedFS, it is inserted into
// the file descriptor table as FdPreopen.
func (c *Context) NewFSContext(stdin io.Reader, stdout, stderr io.Writer, rootFS sysfs.FS) (err error) {
	c.fsc.rootFS = rootFS
	inReader, err := stdinReader(stdin)
	if err != nil {
		return err
	}
	c.fsc.openedFiles.Insert(inReader)
	outWriter, err := stdioWriter(stdout, noopStdoutStat)
	if err != nil {
		return err
	}
	c.fsc.openedFiles.Insert(outWriter)
	errWriter, err := stdioWriter(stderr, noopStderrStat)
	if err != nil {
		return err
	}
	c.fsc.openedFiles.Insert(errWriter)

	if _, ok := rootFS.(sysfs.UnimplementedFS); ok {
		return nil
	}

	if comp, ok := rootFS.(*sysfs.CompositeFS); ok {
		preopens := comp.FS()
		for i, p := range comp.GuestPaths() {
			c.fsc.openedFiles.Insert(&FileEntry{
				FS:        preopens[i],
				Name:      p,
				IsPreopen: true,
				File:      &lazyDir{fs: rootFS},
			})
		}
	} else {
		c.fsc.openedFiles.Insert(&FileEntry{
			FS:        rootFS,
			Name:      "/",
			IsPreopen: true,
			File:      &lazyDir{fs: rootFS},
		})
	}

	return nil
}

func stdinReader(r io.Reader) (*FileEntry, error) {
	if r == nil {
		r = eofReader{}
	}
	var freader *StdioFileReader
	if stdioFileReader, ok := r.(*StdioFileReader); ok {
		freader = stdioFileReader
	} else {
		s, err := stdioStat(r, noopStdinStat)
		if err != nil {
			return nil, err
		}
		freader = NewStdioFileReader(r, s, PollerDefaultStdin)
	}
	return &FileEntry{Name: noopStdinStat.Name(), File: freader}, nil
}

func stdioWriter(w io.Writer, defaultStat stdioFileInfo) (*FileEntry, error) {
	if w == nil {
		w = io.Discard
	}
	s, err := stdioStat(w, defaultStat)
	if err != nil {
		return nil, err
	}
	return &FileEntry{Name: s.Name(), File: &stdioFileWriter{w: w, s: s}}, nil
}

func stdioStat(f interface{}, defaultStat stdioFileInfo) (fs.FileInfo, error) {
	if f, ok := f.(*os.File); ok {
		if st, err := f.Stat(); err == nil {
			mode := uint32(st.Mode() & fs.ModeType)
			return stdioFileInfo{defaultStat[0], mode}, nil
		} else {
			return nil, err
		}
	}
	return defaultStat, nil
}

// RootFS returns the underlying filesystem. Any files that should be added to
// the table should be inserted via InsertFile.
func (c *FSContext) RootFS() sysfs.FS {
	return c.rootFS
}

// OpenFile opens the file into the table and returns its file descriptor.
// The result must be closed by CloseFile or Close.
func (c *FSContext) OpenFile(fs sysfs.FS, path string, flag int, perm fs.FileMode) (int32, syscall.Errno) {
	if f, errno := fs.OpenFile(path, flag, perm); errno != 0 {
		return 0, errno
	} else {
		fe := &FileEntry{openPath: path, FS: fs, File: f, openFlag: flag, openPerm: perm}
		if path == "/" || path == "." {
			fe.Name = ""
		} else {
			fe.Name = path
		}
		if newFD, ok := c.openedFiles.Insert(fe); !ok {
			return 0, syscall.EBADF
		} else {
			return newFD, 0
		}
	}
}

// ReOpenDir re-opens the directory while keeping the same file descriptor.
// TODO: this might not be necessary once we have our own File type.
func (c *FSContext) ReOpenDir(fd int32) (*FileEntry, syscall.Errno) {
	f, ok := c.openedFiles.Lookup(fd)
	if !ok {
		return nil, syscall.EBADF
	} else if _, ft, err := f.CachedStat(); err != nil {
		return nil, platform.UnwrapOSError(err)
	} else if ft.Type() != fs.ModeDir {
		return nil, syscall.EISDIR
	}

	if errno := c.reopen(f); errno != 0 {
		return nil, errno
	}

	f.ReadDir.CountRead, f.ReadDir.Dirents = 0, nil
	return f, 0
}

func (c *FSContext) reopen(f *FileEntry) syscall.Errno {
	if err := f.File.Close(); err != nil {
		return platform.UnwrapOSError(err)
	}

	// Re-opens with  the same parameters as before.
	opened, errno := f.FS.OpenFile(f.openPath, f.openFlag, f.openPerm)
	if errno != 0 {
		return errno
	}

	// Reset the state.
	f.File = opened
	return 0
}

// ChangeOpenFlag changes the open flag of the given opened file pointed by `fd`.
// Currently, this only supports the change of syscall.O_APPEND flag.
func (c *FSContext) ChangeOpenFlag(fd int32, flag int) syscall.Errno {
	f, ok := c.LookupFile(fd)
	if !ok {
		return syscall.EBADF
	} else if _, ft, err := f.CachedStat(); err != nil {
		return platform.UnwrapOSError(err)
	} else if ft.Type() == fs.ModeDir {
		return syscall.EISDIR
	}

	if flag&syscall.O_APPEND != 0 {
		f.openFlag |= syscall.O_APPEND
	} else {
		f.openFlag &= ^syscall.O_APPEND
	}

	// Changing the flag while opening is not really supported well in Go. Even when using
	// syscall package, the feasibility of doing so really depends on the platform. For examples:
	//
	// 	* This appendMode (bool) cannot be changed later.
	// 	https://github.com/golang/go/blob/go1.20/src/os/file_unix.go#L60
	// 	* On Windows, re-opening it is the only way to emulate the behavior.
	// 	https://github.com/bytecodealliance/system-interface/blob/62b97f9776b86235f318c3a6e308395a1187439b/src/fs/fd_flags.rs#L196
	//
	// Therefore, here we re-open the file while keeping the file descriptor.
	// TODO: this might be improved once we have our own File type.
	return c.reopen(f)
}

// LookupFile returns a file if it is in the table.
func (c *FSContext) LookupFile(fd int32) (*FileEntry, bool) {
	return c.openedFiles.Lookup(fd)
}

// Renumber assigns the file pointed by the descriptor `from` to `to`.
func (c *FSContext) Renumber(from, to int32) syscall.Errno {
	fromFile, ok := c.openedFiles.Lookup(from)
	if !ok || to < 0 {
		return syscall.EBADF
	} else if fromFile.IsPreopen {
		return syscall.ENOTSUP
	}

	// If toFile is already open, we close it to prevent windows lock issues.
	//
	// The doc is unclear and other implementations do nothing for already-opened To FDs.
	// https://github.com/WebAssembly/WASI/blob/snapshot-01/phases/snapshot/docs.md#-fd_renumberfd-fd-to-fd---errno
	// https://github.com/bytecodealliance/wasmtime/blob/main/crates/wasi-common/src/snapshots/preview_1.rs#L531-L546
	if toFile, ok := c.openedFiles.Lookup(to); ok {
		if toFile.IsPreopen {
			return syscall.ENOTSUP
		}
		_ = toFile.File.Close()
	}

	c.openedFiles.Delete(from)
	if !c.openedFiles.InsertAt(fromFile, to) {
		return syscall.EBADF
	}
	return 0
}

// CloseFile returns any error closing the existing file.
func (c *FSContext) CloseFile(fd int32) syscall.Errno {
	f, ok := c.openedFiles.Lookup(fd)
	if !ok {
		return syscall.EBADF
	}
	c.openedFiles.Delete(fd)
	return platform.UnwrapOSError(f.File.Close())
}

// Close implements io.Closer
func (c *FSContext) Close() (err error) {
	// Close any files opened in this context
	c.openedFiles.Range(func(fd int32, entry *FileEntry) bool {
		if e := entry.File.Close(); e != nil {
			err = e // This means err returned == the last non-nil error.
		}
		return true
	})
	// A closed FSContext cannot be reused so clear the state instead of
	// using Reset.
	c.openedFiles = FileTable{}
	return
}

// WriterForFile returns a writer for the given file descriptor or nil if not
// opened or not writeable (e.g. a directory or a file not opened for writes).
func WriterForFile(fsc *FSContext, fd int32) (writer io.Writer) {
	if f, ok := fsc.LookupFile(fd); !ok {
		return
	} else if w, ok := f.File.(io.Writer); ok {
		writer = w
	}
	return
}

package libuv

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const (
	FS_UNKNOWN   FsType = -1
	FS_CUSTOM    FsType = 0
	FS_OPEN      FsType = 1
	FS_CLOSE     FsType = 2
	FS_READ      FsType = 3
	FS_WRITE     FsType = 4
	FS_SENDFILE  FsType = 5
	FS_STAT      FsType = 6
	FS_LSTAT     FsType = 7
	FS_FSTAT     FsType = 8
	FS_FTRUNCATE FsType = 9
	FS_UTIME     FsType = 10
	FS_FUTIME    FsType = 11
	FS_ACCESS    FsType = 12
	FS_CHMOD     FsType = 13
	FS_FCHMOD    FsType = 14
	FS_FSYNC     FsType = 15
	FS_FDATASYNC FsType = 16
	FS_UNLINK    FsType = 17
	FS_RMDIR     FsType = 18
	FS_MKDIR     FsType = 19
	FS_MKDTEMP   FsType = 20
	FS_RENAME    FsType = 21
	FS_SCANDIR   FsType = 22
	FS_LINK      FsType = 23
	FS_SYMLINK   FsType = 24
	FS_READLINK  FsType = 25
	FS_CHOWN     FsType = 26
	FS_FCHOWN    FsType = 27
	FS_REALPATH  FsType = 28
	FS_COPYFILE  FsType = 29
	FS_LCHOWN    FsType = 30
	FS_OPENDIR   FsType = 31
	FS_READDIR   FsType = 32
	FS_CLOSEDIR  FsType = 33
	FS_STATFS    FsType = 34
	FS_MKSTEMP   FsType = 35
	FS_LUTIME    FsType = 36
)

const (
	DirentUnknown DirentType = iota
	DirentFile
	DirentDir
	DirentLink
	DirentFifo
	DirentSocket
	DirentChar
	DirentBlock
)

type FsType c.Int

type DirentType c.Int

type UvFile c.Int

// ----------------------------------------------

/* Handle types. */

type Fs struct {
	Unused [0]byte
}

type FsEvent struct {
	Unused [0]byte
}

type FsPoll struct {
	Unused [0]byte
}

type Dirent struct {
	Name *c.Char
	Type DirentType
}

type File struct {
	Loop *Loop
	Req  *Fs
}

type Stat struct {
	Unused [0]byte
}

// ----------------------------------------------

/* Function type */

// llgo:type C
type FsCb func(req *Fs)

// llgo:type C
type FsEventCb func(handle *FsEvent, filename *c.Char, events c.Int, status c.Int)

// llgo:type C
type FsPollCb func(handle *FsPoll, status c.Int, events c.Int)

// ----------------------------------------------

/* Fs related function and method */

//go:linkname FsGetType C.uv_fs_get_type
func FsGetType(req *Fs) *FsType

//go:linkname FsGetPath C.uv_fs_get_path
func FsGetPath(req *Fs) *c.Char

//go:linkname FsGetResult C.uv_fs_get_result
func FsGetResult(req *Fs) c.Int

//go:linkname FsGetPtr C.uv_fs_get_ptr
func FsGetPtr(req *Fs) c.Pointer

//go:linkname FsGetSystemError C.uv_fs_get_system_error
func FsGetSystemError(req *Fs) c.Int

//go:linkname FsGetStatBuf C.uv_fs_get_statbuf
func FsGetStatBuf(req *Fs) *Stat

//go:linkname FsReqCleanup C.uv_fs_req_cleanup
func FsReqCleanup(req *Fs)

//go:linkname DefaultLoop C.uv_default_loop
func DefaultLoop() *Loop

//go:linkname FsOpen C.uv_fs_open
func FsOpen(loop *Loop, req *Fs, path *c.Char, flags c.Int, mode c.Int, cb FsCb) c.Int

//go:linkname FsClose C.uv_fs_close
func FsClose(loop *Loop, req *Fs, file UvFile, cb FsCb) c.Int

//go:linkname FsRead C.uv_fs_read
func FsRead(loop *Loop, req *Fs, file UvFile, bufs *Buf, nbufs c.Uint, offset c.LongLong, cb FsCb) c.Int

//go:linkname FsWrite C.uv_fs_write
func FsWrite(loop *Loop, req *Fs, file UvFile, bufs *Buf, nbufs c.Uint, offset c.LongLong, cb FsCb) c.Int

//go:linkname FsUnlink C.uv_fs_unlink
func FsUnlink(loop *Loop, req *Fs, path *c.Char, cb FsCb) c.Int

//go:linkname FsMkdir C.uv_fs_mkdir
func FsMkdir(loop *Loop, req *Fs, path *c.Char, mode c.Int, cb FsCb) c.Int

//go:linkname FsMkdtemp C.uv_fs_mkdtemp
func FsMkdtemp(loop *Loop, req *Fs, tpl *c.Char, cb FsCb) c.Int

//go:linkname FsMkStemp C.uv_fs_mkstemp
func FsMkStemp(loop *Loop, req *Fs, tpl *c.Char, cb FsCb) c.Int

//go:linkname FsRmdir C.uv_fs_rmdir
func FsRmdir(loop *Loop, req *Fs, path *c.Char, cb FsCb) c.Int

//go:linkname FsStat C.uv_fs_stat
func FsStat(loop *Loop, req *Fs, path *c.Char, cb FsCb) c.Int

//go:linkname FsFstat C.uv_fs_fstat
func FsFstat(loop *Loop, req *Fs, file UvFile, cb FsCb) c.Int

//go:linkname FsRename C.uv_fs_rename
func FsRename(loop *Loop, req *Fs, path *c.Char, newPath *c.Char, cb FsCb) c.Int

//go:linkname FsFsync C.uv_fs_fsync
func FsFsync(loop *Loop, req *Fs, file UvFile, cb FsCb) c.Int

//go:linkname FsFdatasync C.uv_fs_fdatasync
func FsFdatasync(loop *Loop, req *Fs, file UvFile, cb FsCb) c.Int

//go:linkname FsFtruncate C.uv_fs_ftruncate
func FsFtruncate(loop *Loop, req *Fs, file UvFile, offset c.LongLong, cb FsCb) c.Int

//go:linkname FsSendfile C.uv_fs_sendfile
func FsSendfile(loop *Loop, req *Fs, outFd c.Int, inFd c.Int, inOffset c.LongLong, length c.Int, cb FsCb) c.Int

//go:linkname FsAccess C.uv_fs_access
func FsAccess(loop *Loop, req *Fs, path *c.Char, flags c.Int, cb FsCb) c.Int

//go:linkname FsChmod C.uv_fs_chmod
func FsChmod(loop *Loop, req *Fs, path *c.Char, mode c.Int, cb FsCb) c.Int

//go:linkname FsFchmod C.uv_fs_fchmod
func FsFchmod(loop *Loop, req *Fs, file UvFile, mode c.Int, cb FsCb) c.Int

//go:linkname FsUtime C.uv_fs_utime
func FsUtime(loop *Loop, req *Fs, path *c.Char, atime c.Int, mtime c.Int, cb FsCb) c.Int

//go:linkname FsFutime C.uv_fs_futime
func FsFutime(loop *Loop, req *Fs, file UvFile, atime c.Int, mtime c.Int, cb FsCb) c.Int

//go:linkname FsLutime C.uv_fs_lutime
func FsLutime(loop *Loop, req *Fs, path *c.Char, atime c.Int, mtime c.Int, cb FsCb) c.Int

//go:linkname FsLink C.uv_fs_link
func FsLink(loop *Loop, req *Fs, path *c.Char, newPath *c.Char, cb FsCb) c.Int

//go:linkname FsSymlink C.uv_fs_symlink
func FsSymlink(loop *Loop, req *Fs, path *c.Char, newPath *c.Char, flags c.Int, cb FsCb) c.Int

//go:linkname FsReadlink C.uv_fs_read
func FsReadlink(loop *Loop, req *Fs, path *c.Char, cb FsCb) c.Int

//go:linkname FsRealpath C.uv_fs_realpath
func FsRealpath(loop *Loop, req *Fs, path *c.Char, cb FsCb) c.Int

//go:linkname FsCopyfile C.uv_fs_copyfile
func FsCopyfile(loop *Loop, req *Fs, path *c.Char, newPath *c.Char, flags c.Int, cb FsCb) c.Int

//go:linkname FsScandir C.uv_fs_scandir
func FsScandir(loop *Loop, req *Fs, path *c.Char, flags c.Int, cb FsCb) c.Int

//go:linkname FsScandirNext C.uv_fs_scandir_next
func FsScandirNext(req *Fs, ent *Dirent) c.Int

//go:linkname FsOpenDir C.uv_fs_opendir
func FsOpenDir(loop *Loop, req *Fs, path *c.Char, cb FsCb) c.Int

//go:linkname FsReaddir C.uv_fs_readdir
func FsReaddir(loop *Loop, req *Fs, dir c.Int, cb FsCb) c.Int

//go:linkname FsCloseDir C.uv_fs_closedir
func FsCloseDir(loop *Loop, req *Fs) c.Int

//go:linkname FsStatfs C.uv_fs_statfs
func FsStatfs(loop *Loop, req *Fs, path *c.Char, cb FsCb) c.Int

//go:linkname FsChown C.uv_fs_chown
func FsChown(loop *Loop, req *Fs, path *c.Char, uid c.Int, gid c.Int, cb FsCb) c.Int

//go:linkname FsFchown C.uv_fs_fchown
func FsFchown(loop *Loop, req *Fs, file UvFile, uid c.Int, gid c.Int, cb FsCb) c.Int

//go:linkname FsLchown C.uv_fs_lchown
func FsLchown(loop *Loop, req *Fs, path *c.Char, uid c.Int, gid c.Int, cb FsCb) c.Int

//go:linkname FsLstat C.uv_fs_lstat
func FsLstat(loop *Loop, req *Fs, path *c.Char, cb FsCb) c.Int

//go:linkname FsEventInit C.uv_fs_event_init
func FsEventInit(loop *Loop, handle *FsEvent) c.Int

//go:linkname FsEventStart C.uv_fs_event_start
func FsEventStart(handle *FsEvent, cb FsEventCb, path *c.Char, flags c.Int) c.Int

//go:linkname FsEventStop C.uv_fs_event_stop
func FsEventStop(handle *FsEvent) c.Int

//go:linkname FsEventClose C.uv_fs_event_close
func FsEventClose(handle *FsEvent) c.Int

//go:linkname FsEventGetpath C.uv_fs_event_getpath
func FsEventGetpath(handle *FsEvent) *c.Char

//go:linkname FsPollInit C.uv_fs_poll_init
func FsPollInit(loop *Loop, handle *FsPoll) c.Int

//go:linkname FsPollStart C.uv_fs_poll_start
func FsPollStart(handle *FsPoll, cb FsPollCb, path *c.Char, interval uint) c.Int

//go:linkname FsPollStop C.uv_fs_poll_stop
func FsPollStop(handle *FsPoll) c.Int

//go:linkname FsPollClose C.uv_fs_poll_close
func FsPollClose(handle *FsPoll) c.Int

//go:linkname FsPollGetPath C.uv_fs_poll_getpath
func FsPollGetPath(handle *FsPoll) *c.Char

//TODO: Implemnt uv_poll_init_socket

// GetType Get the type of the file system request.
func (f *Fs) GetType() *FsType {
	return FsGetType(f)
}

// GetPath Get the path of the file system request.
func (f *Fs) GetPath() string {
	return c.GoString(FsGetPath(f))
}

// GetResult Get the result of the file system request.
func (f *Fs) GetResult() int {
	return int(FsGetResult(f))
}

// GetPtr Get the pointer of the file system request.
func (f *Fs) GetPtr() c.Pointer {
	return FsGetPtr(f)
}

// GetSystemError Get the system error of the file system request.
func (f *Fs) GetSystemError() int {
	return int(FsGetSystemError(f))
}

// GetStatBuf Get the stat buffer of the file system request.
func (f *Fs) GetStatBuf() *Stat {
	return FsGetStatBuf(f)
}

// Cleanup cleans up the file system request.
func (f *Fs) Cleanup() {
	FsReqCleanup(f)
}

func NewFile(loop *Loop, req *Fs) *File {
	return &File{
		Loop: loop,
		Req:  req,
	}
}

// Open opens a file specified by the path with given flags and mode, and returns a file descriptor.
func (f *Fs) Open(loop *Loop, path string, flags int, mode int, cb FsCb) int {
	return int(FsOpen(loop, f, c.AllocaCStr(path), c.Int(flags), c.Int(mode), cb))
}

// Close closes a file descriptor.
func (f *Fs) Close(loop *Loop, file int, cb FsCb) int {
	return int(FsClose(loop, f, UvFile(file), cb))
}

// Read reads data from a file descriptor into a buffer at a specified offset.
func (f *Fs) Read(loop *Loop, file int, bufs *Buf, nbufs c.Uint, offset int64, cb FsCb) int {
	return int(FsRead(loop, f, UvFile(file), bufs, nbufs, c.LongLong(offset), cb))
}

// Write writes data to a file descriptor from a buffer at a specified offset.
func (f *Fs) Write(loop *Loop, file int, bufs *Buf, nbufs c.Uint, offset int64, cb FsCb) int {
	return int(FsWrite(loop, f, UvFile(file), bufs, nbufs, c.LongLong(offset), cb))
}

// Unlink deletes a file specified by the path.
func (f *Fs) Unlink(loop *Loop, path string, cb FsCb) int {
	return int(FsUnlink(loop, f, c.AllocaCStr(path), cb))
}

// Mkdir creates a new directory at the specified path with a specified mode.
func (f *Fs) Mkdir(loop *Loop, path string, mode int, cb FsCb) int {
	return int(FsMkdir(loop, f, c.AllocaCStr(path), c.Int(mode), cb))
}

// Mkdtemp creates a temporary directory with a template path.
func (f *Fs) Mkdtemp(loop *Loop, tpl string, cb FsCb) int {
	return int(FsMkdtemp(loop, f, c.AllocaCStr(tpl), cb))
}

// MkStemp creates a temporary file from a template path.
func (f *Fs) MkStemp(loop *Loop, tpl string, cb FsCb) int {
	return int(FsMkStemp(loop, f, c.AllocaCStr(tpl), cb))
}

// Rmdir removes a directory specified by the path.
func (f *Fs) Rmdir(loop *Loop, path string, cb FsCb) int {
	return int(FsRmdir(loop, f, c.AllocaCStr(path), cb))
}

// Stat retrieves status information about the file specified by the path.
func (f *Fs) Stat(loop *Loop, path string, cb FsCb) int {
	return int(FsStat(loop, f, c.AllocaCStr(path), cb))
}

// Fstat retrieves status information about a file descriptor.
func (f *Fs) Fstat(loop *Loop, file int, cb FsCb) int {
	return int(FsFstat(loop, f, UvFile(file), cb))
}

// Rename renames a file from the old path to the new path.
func (f *Fs) Rename(loop *Loop, path string, newPath string, cb FsCb) int {
	return int(FsRename(loop, f, c.AllocaCStr(path), c.AllocaCStr(newPath), cb))
}

// Fsync synchronizes a file descriptor's state with storage device.
func (f *Fs) Fsync(loop *Loop, file int, cb FsCb) int {
	return int(FsFsync(loop, f, UvFile(file), cb))
}

// Fdatasync synchronizes a file descriptor's data with storage device.
func (f *Fs) Fdatasync(loop *Loop, file int, cb FsCb) int {
	return int(FsFdatasync(loop, f, UvFile(file), cb))
}

// Ftruncate truncates a file to a specified length.
func (f *Fs) Ftruncate(loop *Loop, file int, offset int64, cb FsCb) int {
	return int(FsFtruncate(loop, f, UvFile(file), c.LongLong(offset), cb))
}

// Sendfile sends data from one file descriptor to another.
func (f *Fs) Sendfile(loop *Loop, outFd int, inFd int, inOffset int64, length int, cb FsCb) int {
	return int(FsSendfile(loop, f, c.Int(outFd), c.Int(inFd), c.LongLong(inOffset), c.Int(length), cb))
}

// Access checks the access permissions of a file specified by the path.
func (f *Fs) Access(loop *Loop, path string, flags int, cb FsCb) int {
	return int(FsAccess(loop, f, c.AllocaCStr(path), c.Int(flags), cb))
}

// Chmod changes the permissions of a file specified by the path.
func (f *Fs) Chmod(loop *Loop, path string, mode int, cb FsCb) int {
	return int(FsChmod(loop, f, c.AllocaCStr(path), c.Int(mode), cb))
}

// Fchmod changes the permissions of a file descriptor.
func (f *Fs) Fchmod(loop *Loop, file int, mode int, cb FsCb) int {
	return int(FsFchmod(loop, f, UvFile(file), c.Int(mode), cb))
}

// Utime updates the access and modification times of a file specified by the path.
func (f *Fs) Utime(loop *Loop, path string, atime int, mtime int, cb FsCb) int {
	return int(FsUtime(loop, f, c.AllocaCStr(path), c.Int(atime), c.Int(mtime), cb))
}

// Futime updates the access and modification times of a file descriptor.
func (f *Fs) Futime(loop *Loop, file int, atime int, mtime int, cb FsCb) int {
	return int(FsFutime(loop, f, UvFile(file), c.Int(atime), c.Int(mtime), cb))
}

// Lutime updates the access and modification times of a file specified by the path, even if the path is a symbolic link.
func (f *Fs) Lutime(loop *Loop, path string, atime int, mtime int, cb FsCb) int {
	return int(FsLutime(loop, f, c.AllocaCStr(path), c.Int(atime), c.Int(mtime), cb))
}

// Link creates a new link to an existing file.
func (f *Fs) Link(loop *Loop, path string, newPath string, cb FsCb) int {
	return int(FsLink(loop, f, c.AllocaCStr(path), c.AllocaCStr(newPath), cb))
}

// Symlink creates a symbolic link from the path to the new path.
func (f *Fs) Symlink(loop *Loop, path string, newPath string, flags int, cb FsCb) int {
	return int(FsSymlink(loop, f, c.AllocaCStr(path), c.AllocaCStr(newPath), c.Int(flags), cb))
}

// Readlink reads the target of a symbolic link.
func (f *Fs) Readlink(loop *Loop, path string, cb FsCb) int {
	return int(FsReadlink(loop, f, c.AllocaCStr(path), cb))
}

// Realpath resolves the absolute path of a file.
func (f *Fs) Realpath(loop *Loop, path string, cb FsCb) int {
	return int(FsRealpath(loop, f, c.AllocaCStr(path), cb))
}

// Copyfile copies a file from the source path to the destination path.
func (f *Fs) Copyfile(loop *Loop, path string, newPath string, flags int, cb FsCb) int {
	return int(FsCopyfile(loop, f, c.AllocaCStr(path), c.AllocaCStr(newPath), c.Int(flags), cb))
}

// Scandir scans a directory for entries.
func (f *Fs) Scandir(loop *Loop, path string, flags int, cb FsCb) int {
	return int(FsScandir(loop, f, c.AllocaCStr(path), c.Int(flags), cb))
}

// OpenDir opens a directory specified by the path.
func (f *Fs) OpenDir(loop *Loop, path string, cb FsCb) int {
	return int(FsOpenDir(loop, f, c.AllocaCStr(path), cb))
}

// Readdir reads entries from an open directory.
func (f *Fs) Readdir(loop *Loop, dir int, cb FsCb) int {
	return int(FsReaddir(loop, f, c.Int(dir), cb))
}

// CloseDir closes an open directory.
func (f *Fs) CloseDir(loop *Loop) int {
	return int(FsCloseDir(loop, f))
}

// Statfs retrieves file system status information.
func (f *Fs) Statfs(loop *Loop, path string, cb FsCb) int {
	return int(FsStatfs(loop, f, c.AllocaCStr(path), cb))
}

// Chown Change file ownership
func (f *Fs) Chown(loop *Loop, path string, uid int, gid int, cb FsCb) int {
	return int(FsChown(loop, f, c.AllocaCStr(path), c.Int(uid), c.Int(gid), cb))
}

// Fchown Change file ownership by file descriptor
func (f *Fs) Fchown(loop *Loop, file int, uid int, gid int, cb FsCb) int {
	return int(FsFchown(loop, f, UvFile(file), c.Int(uid), c.Int(gid), cb))
}

// Lchown Change file ownership (symlink)
func (f *Fs) Lchown(loop *Loop, path string, uid int, gid int, cb FsCb) int {
	return int(FsLchown(loop, f, c.AllocaCStr(path), c.Int(uid), c.Int(gid), cb))
}

// Lstat Get file status (symlink)
func (f *Fs) Lstat(loop *Loop, path string, cb FsCb) int {
	return int(FsLstat(loop, f, c.AllocaCStr(path), cb))
}

// Init Initialize a file event handle
func (e *FsEvent) Init(loop *Loop) int {
	return int(FsEventInit(loop, e))
}

// Start listening for file events
func (e *FsEvent) Start(cb FsEventCb, path string, flags int) int {
	return int(FsEventStart(e, cb, c.AllocaCStr(path), c.Int(flags)))
}

// Stop listening for file events
func (e *FsEvent) Stop() int {
	return int(FsEventStop(e))
}

// Close the file event handle
func (e *FsEvent) Close() int {
	return int(FsEventClose(e))
}

// GetPath Get the path of the file event
func (e *FsEvent) GetPath() *c.Char {
	return FsEventGetpath(e)
}

// Init Initialize a file poll handle
func (p *FsPoll) Init(loop *Loop) int {
	return int(FsPollInit(loop, p))
}

// Start polling for file changes
func (p *FsPoll) Start(cb FsPollCb, path string, interval uint) int {
	return int(FsPollStart(p, cb, c.AllocaCStr(path), interval))
}

// Stop polling for file changes
func (p *FsPoll) Stop() int {
	return int(FsPollStop(p))
}

// Close the file poll handle
func (p *FsPoll) Close() int {
	return int(FsPollClose(p))
}

// GetPath Get the path of the file poll
func (p *FsPoll) GetPath() string {
	return c.GoString(FsPollGetPath(p))
}

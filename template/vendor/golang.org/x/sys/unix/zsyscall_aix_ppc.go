// go run mksyscall_aix_ppc.go -aix -tags aix,ppc syscall_aix.go syscall_aix_ppc.go
// Code generated by the command above; see README.md. DO NOT EDIT.

//go:build aix && ppc
// +build aix,ppc

package unix

/*
#include <stdint.h>
#include <stddef.h>
int utimes(uintptr_t, uintptr_t);
int utimensat(int, uintptr_t, uintptr_t, int);
int getcwd(uintptr_t, size_t);
int accept(int, uintptr_t, uintptr_t);
int getdirent(int, uintptr_t, size_t);
int wait4(int, uintptr_t, int, uintptr_t);
int ioctl(int, int, uintptr_t);
int fcntl(uintptr_t, int, uintptr_t);
int fsync_range(int, int, long long, long long);
int acct(uintptr_t);
int chdir(uintptr_t);
int chroot(uintptr_t);
int close(int);
int dup(int);
void exit(int);
int faccessat(int, uintptr_t, unsigned int, int);
int fchdir(int);
int fchmod(int, unsigned int);
int fchmodat(int, uintptr_t, unsigned int, int);
int fchownat(int, uintptr_t, int, int, int);
int fdatasync(int);
int getpgid(int);
int getpgrp();
int getpid();
int getppid();
int getpriority(int, int);
int getrusage(int, uintptr_t);
int getsid(int);
int kill(int, int);
int syslog(int, uintptr_t, size_t);
int mkdir(int, uintptr_t, unsigned int);
int mkdirat(int, uintptr_t, unsigned int);
int mkfifo(uintptr_t, unsigned int);
int mknod(uintptr_t, unsigned int, int);
int mknodat(int, uintptr_t, unsigned int, int);
int nanosleep(uintptr_t, uintptr_t);
int open64(uintptr_t, int, unsigned int);
int openat(int, uintptr_t, int, unsigned int);
int read(int, uintptr_t, size_t);
int readlink(uintptr_t, uintptr_t, size_t);
int renameat(int, uintptr_t, int, uintptr_t);
int setdomainname(uintptr_t, size_t);
int sethostname(uintptr_t, size_t);
int setpgid(int, int);
int setsid();
int settimeofday(uintptr_t);
int setuid(int);
int setgid(int);
int setpriority(int, int, int);
int statx(int, uintptr_t, int, int, uintptr_t);
int sync();
uintptr_t times(uintptr_t);
int umask(int);
int uname(uintptr_t);
int unlink(uintptr_t);
int unlinkat(int, uintptr_t, int);
int ustat(int, uintptr_t);
int write(int, uintptr_t, size_t);
int dup2(int, int);
int posix_fadvise64(int, long long, long long, int);
int fchown(int, int, int);
int fstat(int, uintptr_t);
int fstatat(int, uintptr_t, uintptr_t, int);
int fstatfs(int, uintptr_t);
int ftruncate(int, long long);
int getegid();
int geteuid();
int getgid();
int getuid();
int lchown(uintptr_t, int, int);
int listen(int, int);
int lstat(uintptr_t, uintptr_t);
int pause();
int pread64(int, uintptr_t, size_t, long long);
int pwrite64(int, uintptr_t, size_t, long long);
#define c_select select
int select(int, uintptr_t, uintptr_t, uintptr_t, uintptr_t);
int pselect(int, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t);
int setregid(int, int);
int setreuid(int, int);
int shutdown(int, int);
long long splice(int, uintptr_t, int, uintptr_t, int, int);
int stat(uintptr_t, uintptr_t);
int statfs(uintptr_t, uintptr_t);
int truncate(uintptr_t, long long);
int bind(int, uintptr_t, uintptr_t);
int connect(int, uintptr_t, uintptr_t);
int getgroups(int, uintptr_t);
int setgroups(int, uintptr_t);
int getsockopt(int, int, int, uintptr_t, uintptr_t);
int setsockopt(int, int, int, uintptr_t, uintptr_t);
int socket(int, int, int);
int socketpair(int, int, int, uintptr_t);
int getpeername(int, uintptr_t, uintptr_t);
int getsockname(int, uintptr_t, uintptr_t);
int recvfrom(int, uintptr_t, size_t, int, uintptr_t, uintptr_t);
int sendto(int, uintptr_t, size_t, int, uintptr_t, uintptr_t);
int nrecvmsg(int, uintptr_t, int);
int nsendmsg(int, uintptr_t, int);
int munmap(uintptr_t, uintptr_t);
int madvise(uintptr_t, size_t, int);
int mprotect(uintptr_t, size_t, int);
int mlock(uintptr_t, size_t);
int mlockall(int);
int msync(uintptr_t, size_t, int);
int munlock(uintptr_t, size_t);
int munlockall();
int pipe(uintptr_t);
int poll(uintptr_t, int, int);
int gettimeofday(uintptr_t, uintptr_t);
int time(uintptr_t);
int utime(uintptr_t, uintptr_t);
unsigned long long getsystemcfg(int);
int umount(uintptr_t);
int getrlimit64(int, uintptr_t);
int setrlimit64(int, uintptr_t);
long long lseek64(int, long long, int);
uintptr_t mmap(uintptr_t, uintptr_t, int, int, int, long long);

*/
import "C"

import (
	"unsafe"
)

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func utimes(path string, times *[2]Timeval) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.utimes(C.uintptr_t(_p0), C.uintptr_t(uintptr(unsafe.Pointer(times))))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func utimensat(dirfd int, path string, times *[2]Timespec, flag int) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.utimensat(C.int(dirfd), C.uintptr_t(_p0), C.uintptr_t(uintptr(unsafe.Pointer(times))), C.int(flag))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func getcwd(buf []byte) (err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	var _p1 int
	_p1 = len(buf)
	r0, er := C.getcwd(C.uintptr_t(uintptr(unsafe.Pointer(_p0))), C.size_t(_p1))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func accept(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (fd int, err error) {
	r0, er := C.accept(C.int(s), C.uintptr_t(uintptr(unsafe.Pointer(rsa))), C.uintptr_t(uintptr(unsafe.Pointer(addrlen))))
	fd = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func getdirent(fd int, buf []byte) (n int, err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	var _p1 int
	_p1 = len(buf)
	r0, er := C.getdirent(C.int(fd), C.uintptr_t(uintptr(unsafe.Pointer(_p0))), C.size_t(_p1))
	n = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func wait4(pid Pid_t, status *_C_int, options int, rusage *Rusage) (wpid Pid_t, err error) {
	r0, er := C.wait4(C.int(pid), C.uintptr_t(uintptr(unsafe.Pointer(status))), C.int(options), C.uintptr_t(uintptr(unsafe.Pointer(rusage))))
	wpid = Pid_t(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func ioctl(fd int, req uint, arg uintptr) (err error) {
	r0, er := C.ioctl(C.int(fd), C.int(req), C.uintptr_t(arg))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func FcntlInt(fd uintptr, cmd int, arg int) (r int, err error) {
	r0, er := C.fcntl(C.uintptr_t(fd), C.int(cmd), C.uintptr_t(arg))
	r = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func FcntlFlock(fd uintptr, cmd int, lk *Flock_t) (err error) {
	r0, er := C.fcntl(C.uintptr_t(fd), C.int(cmd), C.uintptr_t(uintptr(unsafe.Pointer(lk))))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func fcntl(fd int, cmd int, arg int) (val int, err error) {
	r0, er := C.fcntl(C.uintptr_t(fd), C.int(cmd), C.uintptr_t(arg))
	val = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func fsyncRange(fd int, how int, start int64, length int64) (err error) {
	r0, er := C.fsync_range(C.int(fd), C.int(how), C.longlong(start), C.longlong(length))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Acct(path string) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.acct(C.uintptr_t(_p0))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Chdir(path string) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.chdir(C.uintptr_t(_p0))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Chroot(path string) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.chroot(C.uintptr_t(_p0))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Close(fd int) (err error) {
	r0, er := C.close(C.int(fd))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Dup(oldfd int) (fd int, err error) {
	r0, er := C.dup(C.int(oldfd))
	fd = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Exit(code int) {
	C.exit(C.int(code))
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Faccessat(dirfd int, path string, mode uint32, flags int) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.faccessat(C.int(dirfd), C.uintptr_t(_p0), C.uint(mode), C.int(flags))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fchdir(fd int) (err error) {
	r0, er := C.fchdir(C.int(fd))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fchmod(fd int, mode uint32) (err error) {
	r0, er := C.fchmod(C.int(fd), C.uint(mode))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fchmodat(dirfd int, path string, mode uint32, flags int) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.fchmodat(C.int(dirfd), C.uintptr_t(_p0), C.uint(mode), C.int(flags))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fchownat(dirfd int, path string, uid int, gid int, flags int) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.fchownat(C.int(dirfd), C.uintptr_t(_p0), C.int(uid), C.int(gid), C.int(flags))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fdatasync(fd int) (err error) {
	r0, er := C.fdatasync(C.int(fd))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getpgid(pid int) (pgid int, err error) {
	r0, er := C.getpgid(C.int(pid))
	pgid = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getpgrp() (pid int) {
	r0, _ := C.getpgrp()
	pid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getpid() (pid int) {
	r0, _ := C.getpid()
	pid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getppid() (ppid int) {
	r0, _ := C.getppid()
	ppid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getpriority(which int, who int) (prio int, err error) {
	r0, er := C.getpriority(C.int(which), C.int(who))
	prio = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getrusage(who int, rusage *Rusage) (err error) {
	r0, er := C.getrusage(C.int(who), C.uintptr_t(uintptr(unsafe.Pointer(rusage))))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getsid(pid int) (sid int, err error) {
	r0, er := C.getsid(C.int(pid))
	sid = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Kill(pid int, sig Signal) (err error) {
	r0, er := C.kill(C.int(pid), C.int(sig))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Klogctl(typ int, buf []byte) (n int, err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	var _p1 int
	_p1 = len(buf)
	r0, er := C.syslog(C.int(typ), C.uintptr_t(uintptr(unsafe.Pointer(_p0))), C.size_t(_p1))
	n = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mkdir(dirfd int, path string, mode uint32) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.mkdir(C.int(dirfd), C.uintptr_t(_p0), C.uint(mode))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mkdirat(dirfd int, path string, mode uint32) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.mkdirat(C.int(dirfd), C.uintptr_t(_p0), C.uint(mode))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mkfifo(path string, mode uint32) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.mkfifo(C.uintptr_t(_p0), C.uint(mode))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mknod(path string, mode uint32, dev int) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.mknod(C.uintptr_t(_p0), C.uint(mode), C.int(dev))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mknodat(dirfd int, path string, mode uint32, dev int) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.mknodat(C.int(dirfd), C.uintptr_t(_p0), C.uint(mode), C.int(dev))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Nanosleep(time *Timespec, leftover *Timespec) (err error) {
	r0, er := C.nanosleep(C.uintptr_t(uintptr(unsafe.Pointer(time))), C.uintptr_t(uintptr(unsafe.Pointer(leftover))))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Open(path string, mode int, perm uint32) (fd int, err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.open64(C.uintptr_t(_p0), C.int(mode), C.uint(perm))
	fd = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Openat(dirfd int, path string, flags int, mode uint32) (fd int, err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.openat(C.int(dirfd), C.uintptr_t(_p0), C.int(flags), C.uint(mode))
	fd = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func read(fd int, p []byte) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	var _p1 int
	_p1 = len(p)
	r0, er := C.read(C.int(fd), C.uintptr_t(uintptr(unsafe.Pointer(_p0))), C.size_t(_p1))
	n = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Readlink(path string, buf []byte) (n int, err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	var _p1 *byte
	if len(buf) > 0 {
		_p1 = &buf[0]
	}
	var _p2 int
	_p2 = len(buf)
	r0, er := C.readlink(C.uintptr_t(_p0), C.uintptr_t(uintptr(unsafe.Pointer(_p1))), C.size_t(_p2))
	n = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(oldpath)))
	_p1 := uintptr(unsafe.Pointer(C.CString(newpath)))
	r0, er := C.renameat(C.int(olddirfd), C.uintptr_t(_p0), C.int(newdirfd), C.uintptr_t(_p1))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setdomainname(p []byte) (err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	var _p1 int
	_p1 = len(p)
	r0, er := C.setdomainname(C.uintptr_t(uintptr(unsafe.Pointer(_p0))), C.size_t(_p1))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Sethostname(p []byte) (err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	var _p1 int
	_p1 = len(p)
	r0, er := C.sethostname(C.uintptr_t(uintptr(unsafe.Pointer(_p0))), C.size_t(_p1))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setpgid(pid int, pgid int) (err error) {
	r0, er := C.setpgid(C.int(pid), C.int(pgid))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setsid() (pid int, err error) {
	r0, er := C.setsid()
	pid = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Settimeofday(tv *Timeval) (err error) {
	r0, er := C.settimeofday(C.uintptr_t(uintptr(unsafe.Pointer(tv))))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setuid(uid int) (err error) {
	r0, er := C.setuid(C.int(uid))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setgid(uid int) (err error) {
	r0, er := C.setgid(C.int(uid))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setpriority(which int, who int, prio int) (err error) {
	r0, er := C.setpriority(C.int(which), C.int(who), C.int(prio))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Statx(dirfd int, path string, flags int, mask int, stat *Statx_t) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.statx(C.int(dirfd), C.uintptr_t(_p0), C.int(flags), C.int(mask), C.uintptr_t(uintptr(unsafe.Pointer(stat))))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Sync() {
	C.sync()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Times(tms *Tms) (ticks uintptr, err error) {
	r0, er := C.times(C.uintptr_t(uintptr(unsafe.Pointer(tms))))
	ticks = uintptr(r0)
	if uintptr(r0) == ^uintptr(0) && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Umask(mask int) (oldmask int) {
	r0, _ := C.umask(C.int(mask))
	oldmask = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Uname(buf *Utsname) (err error) {
	r0, er := C.uname(C.uintptr_t(uintptr(unsafe.Pointer(buf))))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Unlink(path string) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.unlink(C.uintptr_t(_p0))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Unlinkat(dirfd int, path string, flags int) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.unlinkat(C.int(dirfd), C.uintptr_t(_p0), C.int(flags))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Ustat(dev int, ubuf *Ustat_t) (err error) {
	r0, er := C.ustat(C.int(dev), C.uintptr_t(uintptr(unsafe.Pointer(ubuf))))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func write(fd int, p []byte) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	var _p1 int
	_p1 = len(p)
	r0, er := C.write(C.int(fd), C.uintptr_t(uintptr(unsafe.Pointer(_p0))), C.size_t(_p1))
	n = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func readlen(fd int, p *byte, np int) (n int, err error) {
	r0, er := C.read(C.int(fd), C.uintptr_t(uintptr(unsafe.Pointer(p))), C.size_t(np))
	n = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func writelen(fd int, p *byte, np int) (n int, err error) {
	r0, er := C.write(C.int(fd), C.uintptr_t(uintptr(unsafe.Pointer(p))), C.size_t(np))
	n = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Dup2(oldfd int, newfd int) (err error) {
	r0, er := C.dup2(C.int(oldfd), C.int(newfd))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fadvise(fd int, offset int64, length int64, advice int) (err error) {
	r0, er := C.posix_fadvise64(C.int(fd), C.longlong(offset), C.longlong(length), C.int(advice))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fchown(fd int, uid int, gid int) (err error) {
	r0, er := C.fchown(C.int(fd), C.int(uid), C.int(gid))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func fstat(fd int, stat *Stat_t) (err error) {
	r0, er := C.fstat(C.int(fd), C.uintptr_t(uintptr(unsafe.Pointer(stat))))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func fstatat(dirfd int, path string, stat *Stat_t, flags int) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.fstatat(C.int(dirfd), C.uintptr_t(_p0), C.uintptr_t(uintptr(unsafe.Pointer(stat))), C.int(flags))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fstatfs(fd int, buf *Statfs_t) (err error) {
	r0, er := C.fstatfs(C.int(fd), C.uintptr_t(uintptr(unsafe.Pointer(buf))))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Ftruncate(fd int, length int64) (err error) {
	r0, er := C.ftruncate(C.int(fd), C.longlong(length))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getegid() (egid int) {
	r0, _ := C.getegid()
	egid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Geteuid() (euid int) {
	r0, _ := C.geteuid()
	euid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getgid() (gid int) {
	r0, _ := C.getgid()
	gid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getuid() (uid int) {
	r0, _ := C.getuid()
	uid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Lchown(path string, uid int, gid int) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.lchown(C.uintptr_t(_p0), C.int(uid), C.int(gid))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Listen(s int, n int) (err error) {
	r0, er := C.listen(C.int(s), C.int(n))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func lstat(path string, stat *Stat_t) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.lstat(C.uintptr_t(_p0), C.uintptr_t(uintptr(unsafe.Pointer(stat))))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Pause() (err error) {
	r0, er := C.pause()
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Pread(fd int, p []byte, offset int64) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	var _p1 int
	_p1 = len(p)
	r0, er := C.pread64(C.int(fd), C.uintptr_t(uintptr(unsafe.Pointer(_p0))), C.size_t(_p1), C.longlong(offset))
	n = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Pwrite(fd int, p []byte, offset int64) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	var _p1 int
	_p1 = len(p)
	r0, er := C.pwrite64(C.int(fd), C.uintptr_t(uintptr(unsafe.Pointer(_p0))), C.size_t(_p1), C.longlong(offset))
	n = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Select(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, err error) {
	r0, er := C.c_select(C.int(nfd), C.uintptr_t(uintptr(unsafe.Pointer(r))), C.uintptr_t(uintptr(unsafe.Pointer(w))), C.uintptr_t(uintptr(unsafe.Pointer(e))), C.uintptr_t(uintptr(unsafe.Pointer(timeout))))
	n = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Pselect(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timespec, sigmask *Sigset_t) (n int, err error) {
	r0, er := C.pselect(C.int(nfd), C.uintptr_t(uintptr(unsafe.Pointer(r))), C.uintptr_t(uintptr(unsafe.Pointer(w))), C.uintptr_t(uintptr(unsafe.Pointer(e))), C.uintptr_t(uintptr(unsafe.Pointer(timeout))), C.uintptr_t(uintptr(unsafe.Pointer(sigmask))))
	n = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setregid(rgid int, egid int) (err error) {
	r0, er := C.setregid(C.int(rgid), C.int(egid))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setreuid(ruid int, euid int) (err error) {
	r0, er := C.setreuid(C.int(ruid), C.int(euid))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Shutdown(fd int, how int) (err error) {
	r0, er := C.shutdown(C.int(fd), C.int(how))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Splice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int64, err error) {
	r0, er := C.splice(C.int(rfd), C.uintptr_t(uintptr(unsafe.Pointer(roff))), C.int(wfd), C.uintptr_t(uintptr(unsafe.Pointer(woff))), C.int(len), C.int(flags))
	n = int64(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func stat(path string, statptr *Stat_t) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.stat(C.uintptr_t(_p0), C.uintptr_t(uintptr(unsafe.Pointer(statptr))))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Statfs(path string, buf *Statfs_t) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.statfs(C.uintptr_t(_p0), C.uintptr_t(uintptr(unsafe.Pointer(buf))))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Truncate(path string, length int64) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.truncate(C.uintptr_t(_p0), C.longlong(length))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
	r0, er := C.bind(C.int(s), C.uintptr_t(uintptr(addr)), C.uintptr_t(uintptr(addrlen)))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
	r0, er := C.connect(C.int(s), C.uintptr_t(uintptr(addr)), C.uintptr_t(uintptr(addrlen)))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func getgroups(n int, list *_Gid_t) (nn int, err error) {
	r0, er := C.getgroups(C.int(n), C.uintptr_t(uintptr(unsafe.Pointer(list))))
	nn = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func setgroups(n int, list *_Gid_t) (err error) {
	r0, er := C.setgroups(C.int(n), C.uintptr_t(uintptr(unsafe.Pointer(list))))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error) {
	r0, er := C.getsockopt(C.int(s), C.int(level), C.int(name), C.uintptr_t(uintptr(val)), C.uintptr_t(uintptr(unsafe.Pointer(vallen))))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error) {
	r0, er := C.setsockopt(C.int(s), C.int(level), C.int(name), C.uintptr_t(uintptr(val)), C.uintptr_t(vallen))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func socket(domain int, typ int, proto int) (fd int, err error) {
	r0, er := C.socket(C.int(domain), C.int(typ), C.int(proto))
	fd = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func socketpair(domain int, typ int, proto int, fd *[2]int32) (err error) {
	r0, er := C.socketpair(C.int(domain), C.int(typ), C.int(proto), C.uintptr_t(uintptr(unsafe.Pointer(fd))))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
	r0, er := C.getpeername(C.int(fd), C.uintptr_t(uintptr(unsafe.Pointer(rsa))), C.uintptr_t(uintptr(unsafe.Pointer(addrlen))))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
	r0, er := C.getsockname(C.int(fd), C.uintptr_t(uintptr(unsafe.Pointer(rsa))), C.uintptr_t(uintptr(unsafe.Pointer(addrlen))))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	var _p1 int
	_p1 = len(p)
	r0, er := C.recvfrom(C.int(fd), C.uintptr_t(uintptr(unsafe.Pointer(_p0))), C.size_t(_p1), C.int(flags), C.uintptr_t(uintptr(unsafe.Pointer(from))), C.uintptr_t(uintptr(unsafe.Pointer(fromlen))))
	n = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	var _p1 int
	_p1 = len(buf)
	r0, er := C.sendto(C.int(s), C.uintptr_t(uintptr(unsafe.Pointer(_p0))), C.size_t(_p1), C.int(flags), C.uintptr_t(uintptr(to)), C.uintptr_t(uintptr(addrlen)))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func recvmsg(s int, msg *Msghdr, flags int) (n int, err error) {
	r0, er := C.nrecvmsg(C.int(s), C.uintptr_t(uintptr(unsafe.Pointer(msg))), C.int(flags))
	n = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func sendmsg(s int, msg *Msghdr, flags int) (n int, err error) {
	r0, er := C.nsendmsg(C.int(s), C.uintptr_t(uintptr(unsafe.Pointer(msg))), C.int(flags))
	n = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func munmap(addr uintptr, length uintptr) (err error) {
	r0, er := C.munmap(C.uintptr_t(addr), C.uintptr_t(length))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Madvise(b []byte, advice int) (err error) {
	var _p0 *byte
	if len(b) > 0 {
		_p0 = &b[0]
	}
	var _p1 int
	_p1 = len(b)
	r0, er := C.madvise(C.uintptr_t(uintptr(unsafe.Pointer(_p0))), C.size_t(_p1), C.int(advice))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mprotect(b []byte, prot int) (err error) {
	var _p0 *byte
	if len(b) > 0 {
		_p0 = &b[0]
	}
	var _p1 int
	_p1 = len(b)
	r0, er := C.mprotect(C.uintptr_t(uintptr(unsafe.Pointer(_p0))), C.size_t(_p1), C.int(prot))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mlock(b []byte) (err error) {
	var _p0 *byte
	if len(b) > 0 {
		_p0 = &b[0]
	}
	var _p1 int
	_p1 = len(b)
	r0, er := C.mlock(C.uintptr_t(uintptr(unsafe.Pointer(_p0))), C.size_t(_p1))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mlockall(flags int) (err error) {
	r0, er := C.mlockall(C.int(flags))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Msync(b []byte, flags int) (err error) {
	var _p0 *byte
	if len(b) > 0 {
		_p0 = &b[0]
	}
	var _p1 int
	_p1 = len(b)
	r0, er := C.msync(C.uintptr_t(uintptr(unsafe.Pointer(_p0))), C.size_t(_p1), C.int(flags))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Munlock(b []byte) (err error) {
	var _p0 *byte
	if len(b) > 0 {
		_p0 = &b[0]
	}
	var _p1 int
	_p1 = len(b)
	r0, er := C.munlock(C.uintptr_t(uintptr(unsafe.Pointer(_p0))), C.size_t(_p1))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Munlockall() (err error) {
	r0, er := C.munlockall()
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func pipe(p *[2]_C_int) (err error) {
	r0, er := C.pipe(C.uintptr_t(uintptr(unsafe.Pointer(p))))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func poll(fds *PollFd, nfds int, timeout int) (n int, err error) {
	r0, er := C.poll(C.uintptr_t(uintptr(unsafe.Pointer(fds))), C.int(nfds), C.int(timeout))
	n = int(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func gettimeofday(tv *Timeval, tzp *Timezone) (err error) {
	r0, er := C.gettimeofday(C.uintptr_t(uintptr(unsafe.Pointer(tv))), C.uintptr_t(uintptr(unsafe.Pointer(tzp))))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Time(t *Time_t) (tt Time_t, err error) {
	r0, er := C.time(C.uintptr_t(uintptr(unsafe.Pointer(t))))
	tt = Time_t(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Utime(path string, buf *Utimbuf) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(path)))
	r0, er := C.utime(C.uintptr_t(_p0), C.uintptr_t(uintptr(unsafe.Pointer(buf))))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getsystemcfg(label int) (n uint64) {
	r0, _ := C.getsystemcfg(C.int(label))
	n = uint64(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func umount(target string) (err error) {
	_p0 := uintptr(unsafe.Pointer(C.CString(target)))
	r0, er := C.umount(C.uintptr_t(_p0))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getrlimit(resource int, rlim *Rlimit) (err error) {
	r0, er := C.getrlimit64(C.int(resource), C.uintptr_t(uintptr(unsafe.Pointer(rlim))))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setrlimit(resource int, rlim *Rlimit) (err error) {
	r0, er := C.setrlimit64(C.int(resource), C.uintptr_t(uintptr(unsafe.Pointer(rlim))))
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Seek(fd int, offset int64, whence int) (off int64, err error) {
	r0, er := C.lseek64(C.int(fd), C.longlong(offset), C.int(whence))
	off = int64(r0)
	if r0 == -1 && er != nil {
		err = er
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error) {
	r0, er := C.mmap(C.uintptr_t(addr), C.uintptr_t(length), C.int(prot), C.int(flags), C.int(fd), C.longlong(offset))
	xaddr = uintptr(r0)
	if uintptr(r0) == ^uintptr(0) && er != nil {
		err = er
	}
	return
}

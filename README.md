# Cross-compile Go application for major platforms with Zig and GoReleaser with CGO

Cross-compilation in Go refers to building a Go program on one platform and for another platform. It allows you to create binaries that can be run on systems with different operating systems or architectures from the one on which the program was built. To do that, you just need to specify `GOOS` and `GOARCH` when running go build.

Unfortunately for projects that uses CGO dependencies, things can be harder. Depending on the target architecture it requires the installation of a C compiler like `gcc`, `clang` or `x86_64-w64-mingw64-gcc` and configuring additional environment variables like CC along with the CGO_ENABLED=1 one.

Cross-compiling with cgo involves building a Go program that uses C code and compiling it for a different target platform than the one on which the program is built.

The cgo tool is enabled by default for native builds on systems where it is expected to work.  It is disabled by default when cross-compiling.  You can control this by setting the CGO_ENABLED environment variable when running the go tool: set it to 1 to enable the use of cgo, and to 0 to disable it.  The go tool will set the
build constraint "cgo" if cgo is enabled.

When cross-compiling, you must specify a C cross-compiler for cgo to use. You can do this by setting the CC_FOR_TARGET environment variable when building the toolchain using make.bash, or by setting the CC environment variable any time you run the go tool. The CXX_FOR_TARGET and CXX environment variables work in a similar way for C++ code.
> <https://go-review.googlesource.com/c/go/+/12603/2/src/cmd/cgo/doc.go>

## Zig Cross-compilation

Zig is a programming language that aims to be a better alternative to C, offering a simpler syntax, memory safety, and more expressive error handling. It also includes built-in cross-compilation support.

Zig is a full-fledged C/C++ cross-compiler that leverages LLVM. The crucial detail here is what Zig includes to make cross-compilation possible: Zig bundles standard libraries for all major platforms (GNU libc, musl libc, ...), an advanced artifact caching system, and it has a flag-compatible interface for both clang and gcc.
> <https://dev.to/kristoff/zig-makes-go-cross-compilation-just-work-29ho>

When cross-compiling Zig code, you can use the zig cc and zig c++ commands to compile C and C++ code, respectively. These commands are wrappers around the appropriate compiler for the target platform, and they handle passing the correct flags and options to the underlying compiler.

If you want to cross-compile for x86_64 Linux, for example, all you need to do is add

* CC="zig cc -target x86_64-linux"
* CXX="zig c++ -target x86_64-linux

to the list of env variables when invoking go build. In the case of Hugo, this is the complete command line:

```shell
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC="zig cc -target x86_64-linux" CXX="zig c++ -target x86_64-linux" go build
```

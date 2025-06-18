# pacman

Making the Pac-Man game with Go + Raylib.

Primarily doing this with the intention of becoming familiar with Go, which I have been hoping to try out for a while now. Also, I have heard of Raylib before as a simple graphics library (mainly in the context of C/C++), but I have also not used it earlier, so I thought this might be a good opportunity to use it.

## Setup

### Note for starting from scratch

To initialize the go.mod and go.sum files, just run:

```
go mod init <PACKAGE_NAME>
```

Then, I created the go file (main.go) for the main program source.

(IMPORTANT: For multiple go files, each has to have a package name - but if they are all in the same directory then they must be given the same package name. Alternatively, we can have a subdirectory for other go files, and use a different package name in that case.)

In the terminal run the following to add modules for raylib-go (this is in accordance with the instructions given on https://github.com/gen2brain/raylib-go, for Windows):

```
go get -v -u github.com/gen2brain/raylib-go/raylib
```

We can then import this library in the main.go file using:

```go
import rl "github.com/gen2brain/raylib-go/raylib"
```

### GCC/G++ Resolution for CGO

Since this uses Raylib, there is a dependency on the gcc and g++ compiler (for CGO), and this must be a 64-bit compiler (an issue that took me a while to resolve, since it turns out I had been using a 32 bit MinGW compiler all along till now).

Also, for CGO, the CC and CXX variables need to specifically be set to the paths of the 64-bit gcc and g++ compilers, and these are NOT simply resolved by adding the paths to environment variables PATHS. One way is to just set it manually in the beginning of a terminal session. For example, by running:

```
$env:CC = "C:\msys64\mingw64\bin\gcc.exe"
$env:CXX = "C:\msys64\mingw64\bin\g++.exe"
```

However, instead of doing this each time the terminal window is reopened, it is easier to just add these paths as new User Variables (in the environment variables window).

## Development

To compile and run the project, simply use:

```
go run .
```

(Note: For a single file, one could just do "go run <FILE_NAME>". But that would only compile that file.)

rem Windows requires special linker flags for GUI apps.
rem It's also recommended to use TDM-GCC-64 compiler for CGo.
rem http://tdm-gcc.tdragon.net/download
go build -ldflags="-H windowsgui" -o code-generator.exe
start code-generator.exe
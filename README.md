# Go Native
An example project explaining how to use c/c++ shared library in Golang.

1.  Generate C object,
```bash
gcc -c -Wall -Werror -fpic hello.c
```

2. Generate shared library,
```bash
gcc -shared -o libhello.so hello.o
```

3. Build go,
```bash
go build -v -o main
```

4. Run
```bash
./main
```

# NOTE

### Building `libapi.so`
`cd abi` \
**Linux/MacOS:** `gcc -fPIC -shared -o libapi.so` \
**Windows:** `gcc -fPIC -shared -o libapi.so api.c windows_stubs.c`
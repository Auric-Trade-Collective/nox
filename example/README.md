# NOTE

### Building `libapi.so`
**Linux** \
`cd abi` \
`gcc -fPIC -shared -o libapi.so api.c`

**MacOS** \
`cd abi` \
`gcc -fPIC -dynamiclib -undefined dynamic_lookup -o libapi.dylib api.c`

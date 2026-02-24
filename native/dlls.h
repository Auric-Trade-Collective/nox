
#ifndef DLLS_H
#define DLLS_H

#include <dlfcn.h>
#include <stdio.h>
#include <stdlib.h>

typedef struct {
    void *lib_handle;
} DllManager;


static inline DllManager *LoadDll(char *libname) {
    void *handle = dlopen(libname, RTLD_NOW | RTLD_GLOBAL);
    if(handle == NULL) {
        fprintf(stderr, "dlopen failed: %s\n", dlerror());
        return NULL;
    }

    DllManager *dll = (DllManager*)malloc(sizeof(DllManager));
    dll->lib_handle = handle;
    return dll;
}

static inline void CloseDll(DllManager *dll) {
    dlclose(dll->lib_handle);
    free(dll);
}

#endif

#ifndef WEBAPI_H
#define WEBAPI_H

#include <stdint.h>
#include <string.h>
#include <dlfcn.h>
#include <stdlib.h>
#include <stdio.h>
#include "dlls.h"

typedef struct {

} HttpRequest;

typedef struct {
    uintptr_t gohandle;
} HttpResponse;

typedef void (*apiCallback)(HttpResponse *, HttpRequest *);

typedef struct {
    char *endpoint;
    apiCallback callback;
} NoxEndpoint;


typedef struct {
    DllManager *dll;
    int endpointCount;
    NoxEndpoint *endpoints;
} NoxEndpointCollection;

typedef void (*createEndpoint)(NoxEndpointCollection*, char*, apiCallback);
typedef void (*createNox)(NoxEndpointCollection*, createEndpoint);

static inline char * SanitizePath(char *buff) {
    if(buff == NULL) {
        return NULL;
    }

    int len = 0;
    for(; buff[len] != '\0'; len++);
    len++;

    if(buff[0] != '/') {
        char *newBuff = (char *)malloc(sizeof(char) * (len + 1));
        newBuff[0] = '/';
        for(int i = 1; i < len + 1; i++) {
            newBuff[i] = buff[i - 1];
        }

        free(buff);
        return newBuff;
    }

    return buff;
}

static inline void CreateNoxEndpoint(NoxEndpointCollection *coll, char *endpoint, apiCallback callback) {
    char *sEndp = SanitizePath(strdup(endpoint));
    NoxEndpoint endp = { .endpoint = sEndp, .callback = callback };

    
    NoxEndpoint *ep = (NoxEndpoint *)malloc(sizeof(NoxEndpoint) * (coll->endpointCount + 1));
    memcpy(ep, coll->endpoints, sizeof(NoxEndpoint) * coll->endpointCount);
    ep[coll->endpointCount] = endp;

    free(coll->endpoints);
    coll->endpoints = ep;
    coll->endpointCount++;
}

static inline NoxEndpointCollection *LoadApi(char *location) {
    DllManager *dll = LoadDll(location);
    if(dll == NULL) {
        return NULL;
    }

    NoxEndpointCollection *coll = (NoxEndpointCollection*)malloc(sizeof(NoxEndpointCollection));
    coll->dll = dll;
    coll->endpointCount = 0;
    coll->endpoints = NULL;

    createNox create = (createNox)dlsym(dll->lib_handle, "CreateNoxApi");

    if(create == NULL) {
        printf("Failed to create nox\n");
        return NULL;
    }

    create(coll, CreateNoxEndpoint);

    return coll;
}


static inline void CloseApi(NoxEndpointCollection *coll) {
    if(coll == NULL) return;

    if(coll->dll) {
        CloseDll(coll->dll);
    }

    if(coll->endpoints) {
        for(int i = 0; i < coll->endpointCount; i++) {
            free(coll->endpoints[i].endpoint);
        }
        free(coll->endpoints);
    }

    free(coll);
}

static inline void InvokeApiCallback(apiCallback cb, HttpResponse *resp, HttpRequest *req) {
    cb(resp, req);
}

void WriteStream(HttpResponse *resp, char *buff, int len);

#endif

#ifndef WEBAPI_H
#define WEBAPI_H

#include <string.h>
#include <dlfcn.h>
#include <stdlib.h>
#include "dlls.h"

typedef struct {

} HttpRequest;

typedef struct {

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
typedef NoxEndpoint *(*createNox)(NoxEndpointCollection*, createEndpoint);

void CreateNoxEndpoint(NoxEndpointCollection *coll, char *endpoint, apiCallback callback); 

NoxEndpointCollection *LoadApi(char *location) {
    DllManager *dll = LoadDll(location);
    if(dll == NULL) {
        return NULL;
    }

    NoxEndpointCollection *coll = (NoxEndpointCollection*)malloc(sizeof(NoxEndpointCollection));
    coll->dll = dll;
    coll->endpointCount = 0;
    coll->endpoints = NULL;

    createNox create = (createNox)dlsym(dll->lib_handle, "CreateNoxApi");
    NoxEndpoint *buf = create(coll, CreateNoxEndpoint);

    return coll;
}

void CreateNoxEndpoint(NoxEndpointCollection *coll, char *endpoint, apiCallback callback) {
    NoxEndpoint endp = { .endpoint = strdup(endpoint), .callback = callback };
    
    NoxEndpoint *ep = (NoxEndpoint *)malloc(sizeof(NoxEndpoint) * (coll->endpointCount + 1));
    memcpy(ep, coll->endpoints, sizeof(NoxEndpoint) * coll->endpointCount);
    ep[coll->endpointCount] = endp;

    free(coll->endpoints);
    coll->endpoints = ep;
    coll->endpointCount++;
}

void CloseApi(NoxEndpointCollection *coll) {
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

void InvokeApiCallback(apiCallback cb, HttpResponse *resp, HttpRequest *req) {
    cb(resp, req);
}

#endif

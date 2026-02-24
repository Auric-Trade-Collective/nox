#ifndef WEBAPI_H
#define WEBAPI_H

#include <stdint.h>
#include <string.h>
#include <dlfcn.h> 
#include <stdlib.h>
#include <stdio.h>
#include "dlls.h"

typedef struct {
    char *endpoint;
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

//Here is where the JSON and Stream helpers will exist
//This includes exported function defs, some types, and C logic

typedef enum {
    PLAINTEXT = 0,
    STREAM = 2,
    STREAMFILE = 3, //whilst you can use BYTES/STREAM to stream a file, this will tell
                    //nox what file you want to send, and nox will handle the reading,
                    //writing, parsing, and data types for you
    BYTES= 4,
} NoxDataType;

typedef enum {
    BEGIN = 1,
    PART = 2,
    END = 3,
} NoxStreamSection;

typedef struct {
    NoxDataType type;
    uint8_t *buff; //Does not have to be a Cstring
    size_t length;
    char *filename; // CSTRING, can be NULL
    NoxStreamSection section;
} NoxData;


__attribute__((warning("NoxBuffer: Keep in mind the returned pointer will take ownership of the buffer passed")))
static inline NoxData *NoxBuffer(uint8_t *buff, size_t len) {
    NoxData *dat = (NoxData *)malloc(sizeof(NoxData));
    dat->type = BYTES;
    dat->buff = buff;
    dat->length = len;
    dat->filename = NULL;
    dat->section = (NoxStreamSection)0;

    return dat;
}

static inline NoxData *NoxText(char *buff, size_t len) {
    NoxData *dat = (NoxData *)malloc(sizeof(NoxData));
    dat->type = PLAINTEXT;
    dat->buff = (uint8_t *)strdup(buff);
    dat->length = len;
    dat->filename = NULL;
    dat->section = (NoxStreamSection)0;

    return dat;
}

static inline NoxData *NoxFile(char *filename) {
    NoxData *dat = (NoxData *)malloc(sizeof(NoxData));
    dat->type = STREAMFILE;
    dat->buff = NULL;
    dat->length = -1;
    dat->filename = strdup(filename); //Cstring, can be read without a len
    dat->section = (NoxStreamSection)0;

    return dat;
}

__attribute__((warning("FreeData: This will free any buffers you currently have inside your NoxData pointer")))
static inline void FreeData(NoxData *dat) {
    free(dat->buff);
    free(dat->filename);
    free(dat);
}

//Copies the pointer given
void WriteCopy(HttpResponse *resp, NoxData *dat);
//The pointer given should not be freed by the programmer, but is freed by nox, no copy streaming!
__attribute__((warning("WriteMove: This function takes ownership of all pointer and buffer parameters! Please do not free them!")))
void WriteMove(HttpResponse *resp, NoxData *dat);

void WriteText(HttpResponse *resp, char *buff, int len);
void WriteFile(HttpResponse *resp, NoxData *dat);

#endif

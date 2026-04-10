#ifndef NOX_H
#define NOX_H

#include <stdint.h>
#include <stdbool.h>
#include <string.h>
#include <dlfcn.h> 
#include <stdlib.h>
#include <stdio.h>
#include "dlls.h"
#include<sys/random.h>

// APIS

typedef struct {
    uintptr_t gohandle;

    char *endpoint;
    char *method;
    char *remoteAddr;
    
    // uintptr_t headers;
    // uintptr_t body;
    // uintptr_t url;
} HttpRequest;

typedef struct {
    uintptr_t gohandle;
} HttpResponse;

typedef void (*apiCallback)(HttpResponse *, HttpRequest *);

typedef struct {
    char *endpoint;
    apiCallback callback;
    int method; //0 GET, 1 POST, 2 PUT, 3 DELETE
} NoxEndpoint;

typedef int (*authCallback)(HttpRequest *);

typedef struct {
    DllManager *dll;
    int endpointCount;
    authCallback *auth;
    NoxEndpoint *endpoints;
    char *name;
    char *secret;
} NoxEndpointCollection;

typedef void (*createEndpoint)(NoxEndpointCollection*, char*, apiCallback, int);
typedef void (*createNox)(NoxEndpointCollection*);

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

static inline void CreateNoxEndpoint(NoxEndpointCollection *coll, char *endpoint, apiCallback callback, int method) {
    char *sEndp = SanitizePath(strdup(endpoint));
    NoxEndpoint endp = { .endpoint = sEndp, .callback = callback, .method = method };
    
    NoxEndpoint *ep = (NoxEndpoint *)malloc(sizeof(NoxEndpoint) * (coll->endpointCount + 1));

    if(coll->endpoints != NULL) {
        memcpy(ep, coll->endpoints, sizeof(NoxEndpoint) * coll->endpointCount);
    }

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
    coll->auth = NULL;
    coll->name = NULL;
    coll->secret = NULL;

    createNox create = (createNox)dlsym(dll->lib_handle, "CreateNoxApi");

    if(create == NULL) {
        printf("Failed to create nox\n");
        return NULL;
    }

    create(coll);

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

    free(coll->auth);
    free(coll->name);
    free(coll->secret);
    free(coll);
}


static inline void generate_secure_string(char *buffer, size_t length) {
    const char charset[] = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
    uint8_t random_data[length];

    // Fill the byte array with cryptographically secure bytes
    if (getrandom(random_data, length, 0) == -1) {
        perror("getrandom failed");
        return;
    }

    // Map random bytes to your desired characters
    for (size_t i = 0; i < length; i++) {
        buffer[i] = charset[random_data[i] % (sizeof(charset) - 1)];
    }
    buffer[length] = '\0';
}

static inline char *RegisterName(NoxEndpointCollection *coll, char *name) {
    char *secure = (char *)malloc(16);
    generate_secure_string(secure, 16);

    char *duped = strdup(name);
    
    coll->name = duped;
    coll->secret = secure;

    return secure;
}

static inline void CreateAuth(NoxEndpointCollection *coll, authCallback cb) {
    authCallback *cbp = (authCallback *)malloc(sizeof(authCallback));
    *cbp = cb;
    coll->auth = cbp;
}

static inline void InvokeApiCallback(apiCallback cb, HttpResponse *resp, HttpRequest *req) {
    cb(resp, req);
}

static inline int InvokeAuth(authCallback cb, HttpRequest *req) {
    return cb(req);
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
    char *contentType;
} NoxData;


__attribute__((warning("NoxBuffer: Keep in mind the returned pointer will take ownership of the buffer passed to it")))
static inline NoxData *NoxBuffer(uint8_t *buff, size_t len, char *contentType) {
    NoxData *dat = (NoxData *)malloc(sizeof(NoxData));
    dat->type = BYTES;
    dat->buff = buff;
    dat->length = len;
    dat->filename = NULL;
    dat->section = (NoxStreamSection)0;
    dat->contentType = strdup(contentType);

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
    free(dat->contentType);
    free(dat);
}

//Copies the pointer given
void WriteCopy(HttpResponse *resp, NoxData *dat);
//The pointer given should not be freed by the programmer, but is freed by nox, no copy streaming!
__attribute__((warning("WriteMove: This function takes ownership of all pointer and buffer parameters! Please do not free them!")))
void WriteMove(HttpResponse *resp, NoxData *dat);

void WriteText(HttpResponse *resp, char *buff, int len);
void WriteFile(HttpResponse *resp, NoxData *dat);


void CreateGet(NoxEndpointCollection *collection, char *path, apiCallback callback);
void CreatePost(NoxEndpointCollection *collection, char *path, apiCallback callback);
void CreatePut(NoxEndpointCollection *collection, char *path, apiCallback callback);
void CreateDelete(NoxEndpointCollection *collection, char *path, apiCallback callback);

int TryGetResponseHeader(HttpResponse *resp, char *key, size_t index, char **out);
int TrySetResponseHeader(HttpResponse *resp, char *key, char *val, int add);

int TryGetRequestHeader(HttpRequest *resp, char *key, size_t index, char **out);
int TrySetRequestHeader(HttpRequest *resp, char *key, char *val, int add);

// the returned value is how many bytes are read
size_t ReadBody(HttpRequest *req, uint8_t *buff, size_t bytesToRead);

char *GetUri(HttpRequest *req, size_t *outLength);
int TryGetUriParam(HttpRequest *req, char *key, size_t index, char **out, size_t *outLen);
size_t GetUriParamCount(HttpRequest *req, char *paramName);

char *TryGetCookie(HttpRequest *req, char *key);
void TrySetCookie(HttpResponse *resp, char *key, char *value, char *path, long expires, bool secure, bool httponly);

void LogWrite(char *name_space, char *msg);
void LogWarn(char *name_space, char *msg);
void LogError(char *name_space, char *msg);
void LogPanic(char *name_space, char *msg);
void LogDebug(char *name_space, char *msg);

char *GetEnv(char *secret, char *key);

void TemporaryRedirect(HttpResponse *resp, HttpRequest *req, char *loc);
void PermanentRedirect(HttpResponse *resp, HttpRequest *req, char *loc);

//PLUGINS

typedef struct {
    void *handle;
} PluginCtx;

typedef struct {
    int type;
    char *error;
    HttpRequest *httpRequest;
    HttpResponse *httpResponse;
} EventCtx;

typedef void (*pluginMain)(PluginCtx*);
typedef void (*eventCallback)(EventCtx*);

static inline void InvokePluginMain(PluginCtx *ctx, pluginMain cb) {
    cb(ctx);
}

enum EventType {
    OnLog = 0,
    OnError = 1,
    OnRequest = 2,
    OnResponse = 3,
    OnAny = 4,
};

static inline void RegisterEvent(PluginCtx *plugin, int eventType, eventCallback cb) {

}


#endif

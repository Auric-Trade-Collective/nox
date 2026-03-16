#include "../../native/webapi.h"
#include <stdio.h>
#include <string.h>
#include <time.h>

void SomeCookie(HttpResponse *resp, HttpRequest *req) {
    TrySetCookie(resp, "a", "abcd", "/", time(NULL) + (60 * 60 * 24), false, false);
}

void OtherCookie(HttpResponse *resp, HttpRequest *req) {
    char *val = TryGetCookie(req, "a");
    WriteText(resp, val, 4);

    free(val);
}

void DoSomething(HttpResponse *resp, HttpRequest *req) {
    // if(X() == 5) {
    //     WriteText(resp, "Hello", 5);
    // } else {
    //     WriteText(resp, "Bye", 3);
    // }

    TrySetResponseHeader(resp, "test", "a test", 1);

    char *ptr;
    size_t len;
    if(TryGetUriParam(req, "test", 0, &ptr, &len) == 1) {
        WriteText(resp, ptr, (int)len);
        free(ptr);
    }

}

void DoSomethingPost(HttpResponse *resp, HttpRequest *req) {
    // if(X() == 5) {
    //     WriteText(resp, "Hello", 5);
    // } else {
    //     WriteText(resp, "Bye", 3);
    // }
    WriteText(resp, "Foo Post", 8);
}

void DoSomething2(HttpResponse *resp, HttpRequest *req) {
    WriteText(resp, "Foo", 3);
}

int NoxAuth(HttpRequest *req) {
    //eventually we need to implement cookies!
    char *header;
    if(TryGetRequestHeader(req, "auth", 0, &header) == 1) {
        if(strcmp(header, "test")) return 1;
        free(header);
    }

    return 0;
}

void CreateNoxApi(NoxEndpointCollection *coll) {
    printf("Loading Nox API \n");

    // CreateAuth(coll, NoxAuth);
    CreateGet(coll, "foo", DoSomething);
    CreateGet(coll, "getcookie", SomeCookie);
    CreateGet(coll, "setcookie", OtherCookie);
    CreatePost(coll, "foo", DoSomethingPost);

    LogWrite("Nox API", "Okay we can log from here!");
}

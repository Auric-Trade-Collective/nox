#include "../../native/webapi.h"
#include <stdio.h>

void DoSomething(HttpResponse *resp, HttpRequest *req) {
    // if(X() == 5) {
    //     WriteText(resp, "Hello", 5);
    // } else {
    //     WriteText(resp, "Bye", 3);
    // }
    WriteText(resp, "Foo", 3);
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

void CreateNoxApi(NoxEndpointCollection *coll) {
    printf("Loading Nox API \n");
    CreateGet(coll, "foo", DoSomething);
    CreatePost(coll, "foo", DoSomethingPost);
}

#include "../../native/webapi.h"
#include <stdio.h>

void DoSomething(HttpResponse *resp, HttpRequest *req) {
    // if(X() == 5) {
    //     WriteText(resp, "Hello", 5);
    // } else {
    //     WriteText(resp, "Bye", 3);
    // }
}

void DoSomething2(HttpResponse *resp, HttpRequest *req) {
    WriteText(resp, "Foo", 3);
}

void CreateNoxApi(NoxEndpointCollection *coll, createEndpoint endp) {
    printf("Loading Nox API \n");
    endp(coll, "foo", DoSomething);
    endp(coll, "/bar", DoSomething2);
}

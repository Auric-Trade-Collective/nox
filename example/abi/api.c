#include "../../native/webapi.h"
#include <stdio.h>

void DoSomething(HttpResponse *resp, HttpRequest *req) {
    for(size_t i = 0; i < 1024 * 1024 * 1024; i++) {
        WriteText(resp, "Foo", 3);
    }

    WriteText(resp, "\0", 1);
}

void DoSomething2(HttpResponse *resp, HttpRequest *req) {
    WriteText(resp, "Bar", 3);
}

void CreateNoxApi(NoxEndpointCollection *coll, createEndpoint endp) {
    printf("Loading Nox API \n");
    endp(coll, "foo", DoSomething);
    endp(coll, "/bar", DoSomething2);
}

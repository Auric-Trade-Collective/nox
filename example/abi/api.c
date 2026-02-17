#include "../../native/webapi.h"
#include <stdio.h>

void DoSomething(HttpResponse *resp, HttpRequest *req) {
    WriteStream(resp, "Foo", 3);
}

void DoSomething2(HttpResponse *resp, HttpRequest *req) {
    WriteStream(resp, "Bar", 3);
}

void CreateNoxApi(NoxEndpointCollection *coll, createEndpoint endp) {
    printf("Loading Nox API \n");
    endp(coll, "foo", DoSomething);
    endp(coll, "/bar", DoSomething2);
}

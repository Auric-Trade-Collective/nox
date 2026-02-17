#include "../../native/webapi.h"
#include <stdio.h>

void DoSomething(HttpResponse *resp, HttpRequest *req) {
    printf("Hi!");
}

void CreateNoxApi(NoxEndpointCollection *coll, createEndpoint endp) {
    printf("Loading Nox API \n");
    endp(coll, "/foo", DoSomething);
}

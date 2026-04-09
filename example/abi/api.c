#include "../../include/nox.h"
#include <stdio.h>

void SomeCookie(HttpResponse *resp, HttpRequest *req) {
    TemporaryRedirect(resp, req, "https://google.com/");
}

int NoxAuth(HttpRequest *req) {
    return 0;
}

void CreateNoxApi(NoxEndpointCollection *coll) {
    printf("Loading Nox API \n");

    CreateAuth(coll, NoxAuth);
    CreateGet(coll, "/foo", SomeCookie);

    LogWrite("Nox API", "Okay we can log from here!");
}

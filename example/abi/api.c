#include "../../include/nox.h"
#include <stdio.h>

void SomeCookie(HttpResponse *resp, HttpRequest *req) {
    LogDebug("Nox API", "Redirecting this cool dude!");
    // TemporaryRedirect(resp, req, "https://google.com/");

    ApiStatusPage(resp, 404, 1);
}

int NoxAuth(HttpRequest *req) {
    return 0;
}

void CreateNoxApi(NoxEndpointCollection *coll) {
    printf("Loading Nox API \n");

    CreateAuth(coll, NoxAuth);
    CreateGet(coll, "/foo", SomeCookie);

    RegisterName(coll, "test");

    LogWrite("Nox API", "Okay we can log from here!");
}

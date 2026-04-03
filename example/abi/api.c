#include "../../include/nox.h"
#include <stdio.h>
#include <string.h>
#include <time.h>

void SomeCookie(HttpResponse *resp, HttpRequest *req) {
    TrySetCookie(resp, "a", "abcd", "/", time(NULL) + (60 * 60 * 24), false, false);
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

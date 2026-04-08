#include "../../include/nox.h"
#include <stdio.h>
#include <strings.h>

static char *secret;

void SomeCookie(HttpResponse *resp, HttpRequest *req) {
    char *var = GetEnv(secret, "test");
    WriteText(resp, var, strlen(var));
}

int NoxAuth(HttpRequest *req) {
    return 1;
}

void CreateNoxApi(NoxEndpointCollection *coll) {

    printf("Loading Nox API \n");

    secret = RegisterName(coll, "test");
    CreateAuth(coll, NoxAuth);
    CreateGet(coll, "/foo", SomeCookie);

    LogWrite("Nox API", "Okay we can log from here!");
}

typedef struct {

} HtmlRequest;

typedef struct {

} HtmlResponse;

typedef void (*funcHandler)(HtmlResponse*, HtmlRequest*);
typedef void (*registerFunc)(char *, funcHandler);

void TestEndpoint(HtmlResponse *resp, HtmlRequest *req) {

}

void createNox(registerFunc reg) {
    reg("/test", TestEndpoint);
}


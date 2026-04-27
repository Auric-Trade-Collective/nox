package pages

var Pg401 string = `
<p>401 unauthorized</p>
`

var Pg404 string = `
<p>404 not found</p>
`

var Pg500 string = `
<p>500 interal server error</p>
`

var Pages = map[int]string{
	100: "<p>Continue</p>",
	101: "<p>Switching Protocols</p>",
	102: "<p>Processing</p>",
	103: "<p>Early Hints</p>",
	200: "<p>OK</p>",
	201: "<p>Created</p>",
	202: "<p>Accepted</p>",
	203: "<p>Non-Authoritative Information</p>",
	204: "<p>No Content</p>",
	205: "<p>Reset Content</p>",
	206: "<p>Partial Content</p>",
	207: "<p>Multi-Status</p>",
	208: "<p>Already Reported</p>",
	226: "<p>IM Used</p>",
	300: "<p>Multiple Choices</p>",
	301: "<p>Moved Permanently</p>",
	302: "<p>Found</p>",
	303: "<p>See Other</p>",
	304: "<p>Not Modified</p>",
	305: "<p>Use Proxy</p>",
	307: "<p>Temporary Redirect</p>",
	308: "<p>Permanent Redirect</p>",
	400: "<p>Bad Request</p>",
	401: "<p>Unauthorized</p>",
	402: "<p>Payment Required</p>",
	403: "<p>Forbidden</p>",
	404: "<p>Not Found</p>",
	405: "<p>Method Not Allowed</p>",
	406: "<p>Not Acceptable</p>",
	407: "<p>Proxy Authentication Required</p>",
	408: "<p>Request Timeout</p>",
	409: "<p>Conflict</p>",
	410: "<p>Gone</p>",
	411: "<p>Length Required</p>",
	412: "<p>Precondition Failed</p>",
	413: "<p>Payload Too Large</p>",
	414: "<p>URI Too Long</p>",
	415: "<p>Unsupported Media Type</p>",
	416: "<p>Range Not Satisfiable</p>",
	417: "<p>Expectation Failed</p>",
	418: "<p>I'm a teapot</p>",
	421: "<p>Misdirected Request</p>",
	422: "<p>Unprocessable Entity</p>",
	423: "<p>Locked</p>",
	424: "<p>Failed Dependency</p>",
	425: "<p>Too Early</p>",
	426: "<p>Upgrade Required</p>",
	428: "<p>Precondition Required</p>",
	429: "<p>Too Many Requests</p>",
	431: "<p>Request Header Fields Too Large</p>",
	451: "<p>Unavailable For Legal Reasons</p>",
	500: "<p>Internal Server Error</p>",
	501: "<p>Not Implemented</p>",
	502: "<p>Bad Gateway</p>",
	503: "<p>Service Unavailable</p>",
	504: "<p>Gateway Timeout</p>",
	505: "<p>HTTP Version Not Supported</p>",
	506: "<p>Variant Also Negotiates</p>",
	507: "<p>Insufficient Storage</p>",
	508: "<p>Loop Detected</p>",
	510: "<p>Not Extended</p>",
	511: "<p>Network Authentication Required</p>",
}

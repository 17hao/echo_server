namespace go api

struct EchoRequest {
	1: string message
}

struct EchoResponse {
	1: string message
}

service EchoService {
    EchoResponse echo(1: EchoRequest req)
}

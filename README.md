# GRPC


- **Unary RPC**

    Client sends a single request to the Server and gets a single response back, just like a normal function call.

- **Server Streaming RPC**

  Client sends a single request and server sends a stream of response. The client reads from the returned stream until there are no more messages.

- **Client Streaming RPC**

  Client sends a stream of data and gets a single response back. Once the client has finished writing the messages, it waits for the server to read them and return its response

- **Bidirectional Streaming RPC**

  Both client and server send stream of data. The two streams operate independently, so clients and servers can read and write in whatever order they like.
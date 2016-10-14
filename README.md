# ident

Generating unique identifiers in a distributed context can be a really
challenging problem. This is a simple HTTP server that makes it easy using
minimal state synchronization between processes.

## How it Works

`ident` is similar in nature to Snowflake. At the heart of `ident` is an
algorithm that generates a cluster-wide 96-bit identifier composed a components
that are not unique by themselves but become unique when combined together.

#### Example identifier

```
ACsAAlgBZtfAVYWs
```

#### Generation

Here's an example identifier. It's presented base64 encoded via the API, but
for explanation purposes we present it as a series of bytes.

```
0x0 0x2b 0x0 0x1 0x58 0x1 0x66 0xd0 0xce 0x8b 0xdb 0x26
```
The components of the first 64-bits are as follows

| Start Bit | Stop Bit | Name | Description                                                             |
|-----------|----------|------|-------------------------------------------------------------------------|
| 0         | 15       | ID   | Cluster-wide unique ID of for the  process.                             |
| 16        | 31       | Seq  | Sequence number in the set of identifiers generated within this second. |
| 32        | 63       | Time | The time in seconds since epoch that this identifier was generated.     |
| 64        | 95       | Rand | A random 32 bit number generated in-process.                            |

This is all wrapped up behind an HTTP interface to make it easy to generate new
identifiers with.

```bash
$ curl -i http://localhost:8081
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Server: ident v0.1.0
Date: Fri, 14 Oct 2016 23:15:33 GMT
Content-Length: 33

{"identifier":"ACwAAVgBZxU064V4"}
```

### Generating cluster-unique IDs

The easiest way to generate a cluster-wide unique identifier is to use a
highly-consistent source to generate monotonically-increasing numbers.
Thankfully, DynamoDB has [exactly that
feature](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithItems.html#WorkingWithItems.AtomicCounters).
The default implementation uses Dynamo to accomplish this.

## Deployment

TBD. Use the Cloudformation script provided.

## Development

As with any open source tool, patches are welcome. Please make sure test cases
are provided, where relevant. Just open a PR and we can look at integrating it.

## License

ISC

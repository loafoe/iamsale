# iamsale

An experimental API implementation for integrating with Identity brokers

## Development

This project leverages [goa](https://goa.design/) for designing the API and use, so you'll need the `goa` binary for development. Install:

```shell
go install goa.design/goa/v3/cmd/goa@v3
```

### Generate files

```shell
goa gen github.com/loafoe/iamsale/design
```

### Configuration

The application reads it config from `iamsale.yaml`

```yaml
iam:
  region: us-east
  environment: client-test
  serviceId: "some.app.prosition@someorg.philips-healthsuite.com"
  servicePrivateKey: $KEY_FROM_ENV
  orgId: "460631dc-2ddd-4c38-a18d-7abe72ecb4c9"

api:
  username: "iamsale-client"
  password: $PWD_FROM_ENV
```

## License

License is [Apache 2.0](LICENSE.md)

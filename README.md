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

### Implement

You are now ready to implement the services

## License

License is [Apache 2.0](LICENSE.md)

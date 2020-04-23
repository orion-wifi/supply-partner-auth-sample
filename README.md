# Readme:

This Go binary provides a CLI for generating a JWT in a predeterimined format.

## Running with Docker

By running the CLI in a docker container, you are essentially building & running
the Go binary in a virtulized environment; meaning there is no modification
necessary to your local system. The only dependency for this is to have
[docker](https://www.docker.com/) installed and working on your system. Docker
is available for Mac, Windows and Nix.

A shell script (`run.sh`) is provided that will automatically build and run a
docker container for you. You should open this script (with your favourite text
editor) and replace the value for `SERVICE_ACCOUNT_LOCATION` with the path and
filename to the Service Account JSON file you downnloaed from the GCP console.
You can then run the CLI within a self-contained

```
./ruh.sh

```

## Running from Source

If you would prefer to build and run this CLI from source, you will need the Go
toolchain installed on your system. You can follow your system specific
installation steps [here](https://golang.org/doc/install).

Once installed you can compile and run the Go program by doing:

```
go run main.go --service-account=$HOME/path/to/local/service/account.json
```

## Expected Output:

If the CLI executed successfully you should expect output to your terminal that
looks approximately like the following:

```
$ ./run.sh
sha256:01c1491a84af2b52491ceda1c0aa3fcb2b4ad3a96d99fc299a1778e105dd30fe
Token generated:

=== BEGIN TOKEN ===
<token_redacted_for_readme>
=== END TOKEN ===
```

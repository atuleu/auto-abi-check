# auto-abi-check

A go wrapper arround abi-compliance-checker to automate C/C++ ABI
compliance check

## Why go ?

This project started first as a ruby gem. My concern was to provide
easy tool to add compliance check in my debain package. And ruby gem
and debian package is a bit ... messy. The fact that these tools could
be easily be compiled on debian and packaged in a single binary made
me switch to go.

## Usage

Generate a config file using options
```bash
 $ auto-abi-check generate [OPTIONS] <dest_config>
```

Run a check 

```bash
 $ auto-abi-check run --new [OPTIONS] --old [OPTIONS]
```

### Options for configuration

This script is a wrapper to automate actions for abi-compliance
check. What you would like to do is to insure ABI comaptiniluty
accross verison of your software. Instead to maintain locally
different copy of your software, if you maitain packages of your
software too, you could use them to check ABI compatinility.


## CMake integration

And why not integrate it with CMake, it would be nice no ?

## Contribute  


1. Fork it on github
2. Create a new feature / issue branch
3. Hack 
4. Push
5. Open Pull Request

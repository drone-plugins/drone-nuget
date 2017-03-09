Use this plugin to publish artifacts from the build to a NuGet Repository
You can override the default configuration with the following parameters:

* `source` - NuGet Repository URL
* `api_key` - Api Key used for authentication
* `verbosity` - NuGet output verbosity, default to 'quiet'. Accept: normal, quiet or detailed
* `files` - List of files to upload

All file paths must be relative to current project sources

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
publish:
  nuget:
    image: quay.io/urbit/drone-nuget
    source: http://nuget.company.com
    api_key: <Your Key>
    files: 
      - *.nupkg
```

## .nuspec / .nupkg

If a file to upload does have ```.nuspec``` extension, the __nuget pack__ command is called before a push

If a file to upload does have ```.nupkg``` extension, it is pushed directly

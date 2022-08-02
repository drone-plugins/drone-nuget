A plugin to push packages to nuget.

# Preliminary steps

In order to make best use of this plugin please follow the steps below:

1. Add the package metadata like so:
```xml
<PackageId>AppLogger</PackageId>
<Version>1.0.0</Version>
<Authors>your_name</Authors>
<Company>your_company</Company>
```
2. Automatically generate package on build
   To automatically run dotnet pack when you run dotnet build, add the following line to your project file within <PropertyGroup>:

```xml
<GeneratePackageOnBuild>true</GeneratePackageOnBuild>
```
# Usage

The following settings are required for this plugin

* PLUGIN_NUGET_APIKEY - this is used to authenticate with nuget.

* PLUGIN_NUGET_URI - nuget base url

* PLUGIN_PACKAGE_LOCATION (optional) - the location of the package you wish to publish. Default behaviour will push all packages.


Below is an example `.drone.yml` that uses this plugin.

```yaml
kind: pipeline
name: default
type: docker

steps:
  - name: build
    image: mcr.microsoft.com/dotnet/sdk:5.0
    pull: if-not-exists
    commands:
      - dotnet build
  - name: publish
    image: drone/drone-nuget
    pull: if-not-exists
    settings:
      log_level: debug
      nuget_apikey:
        from_secret: nuget_apikey
      nuget_uri: "https://api.nuget.org/v3/index.json"
      package_location: "SomePackageLocation"
```

# Building

Build the plugin binary:

```text
scripts/build.sh
```

Build the plugin image:

```text
docker build -t plugins/drone-nuget -f docker/Dockerfile .
```

# Testing

Execute the plugin from your current working directory:

```text
docker run --rm -e PLUGIN_NUGET_APIKEY="someKey" \
  -e PLUGIN_NUGET_URI="someUrl" \
  -e PLUGIN_PACKAGE_LOCATION="someLocation" \
  -e DRONE_COMMIT_SHA=8f51ad7884c5eb69c11d260a31da7a745e6b78e2 \
  -e DRONE_COMMIT_BRANCH=master \
  -e DRONE_BUILD_NUMBER=43 \
  -e DRONE_BUILD_STATUS=success \
  -w /drone/src \
  -v $(pwd):/drone/src \
  plugins/drone-nuget
```

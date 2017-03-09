# drone-nuget

[![Build Status](http://ci.urb-it.com/api/badges/urbitassociates/drone-nuget/status.svg)](http://ci.urb-it.com/urbitassociates/drone-nuget)
[![Docker Repository on Quay](https://quay.io/repository/urbit/drone-nuget/status "Docker Repository on Quay")](https://quay.io/repository/urbit/drone-nuget)

Drone plugin to publish files and artifacts to NuGet repository. For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

## Execute

Install the deps using `make`:

```
make install
```

### Example

```sh
npm start <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "source": "http://nuget.company.com",
        "api_key": "SUPER_KEY",
        "files": [
            "*.nupkg"
        ]
    }
}
EOF
```

## Docker

Build the container using `make`:

```
make docker
```

### Example

```sh
docker run -i quay.io/urbit/drone-nuget <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "source": "http://nuget.company.com",
        "api_key": "SUPER_KEY",
        "files": [
            "*.nupkg"
        ]
    }
}
EOF
```

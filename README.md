# drone-nuget

Drone plugin for uploading files to NuGet repository

# Usage

        node index.js <<EOF
        {
            "repo": {
                "clone_url": "git://github.com/drone/drone",
                "full_name": "drone/drone"
            },
            "build": {
                "number": 1,
                "event": "push",
                "branch": "master",
                "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
                "ref": "refs/heads/master",
                "status": "success"
            },
            "workspace": {
                "root": "/drone/src",
                "path": "/drone/src/athieriot/drone-nuget"
            },
            "vargs": {
                "image": "athieriot/drone-nuget",
                "source": "http://nuget.company.com",
                "api_key": "SUPER_KEY",
                "files": [
                    "*.nupkg"
                ]
            }
        }
        EOF

# Docker

Build the Docker container:

    docker build -t athieriot/drone-nuget .

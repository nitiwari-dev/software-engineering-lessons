## Repo for all things on CI-CD

## How to run github runner [locally](https://github.com/actions/runner/releases) ? 

- Install into the machine you want start the runner
```sh
mkdir actions-runner && cd actions-runner
curl -O -L https://github.com/actions/runner/releases/download/v2.320.0/actions-runner-osx-arm64-2.320.0.tar.gz
tar xzf ./actions-runner-osx-arm64-2.320.0.tar.gz
```

- Follow the steps from github - Settings -> Actions -> Runner
  - This configure your system to act as runner and will ask options like runner name, group etc
    - ![Self hosted runner](..%2F..%2F..%2F..%2F..%2F..%2F..%2Fvar%2Ffolders%2F4z%2Fpqdpw4zn705fz3k8z8_gq2nw0000gq%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_IRlKF9%2FScreenshot%202024-11-11%20at%201.35.25%E2%80%AFAM.png)


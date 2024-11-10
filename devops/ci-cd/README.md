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
    <img width="792" alt="Screenshot 2024-11-11 at 1 37 55 AM" src="https://github.com/user-attachments/assets/060b780a-8eb6-4d73-b3f7-be8e142ebdb3">
  - Runner running on the local after the PR push

     <img width="441" alt="Screenshot 2024-11-11 at 1 38 31 AM" src="https://github.com/user-attachments/assets/85cedf35-8c0d-4e06-a528-430819be183c">

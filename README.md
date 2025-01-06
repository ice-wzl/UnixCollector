# UnixCollector
## Overview

## Benifits 
- `UnixCollector` will rebuild the directory structure of the host. For example if you collect `/etc/fstab` you will see that same directory struture reflected in your exfil directory
- See below for a visual example of the rebuilt directory structure
![image](https://github.com/user-attachments/assets/2515b50c-9e22-4dab-854d-94cbb75b2ad8)

## Contributing
- Im always happy to accept PRs!
- Ive attempted to make this code as module as possible to make it as easy as possible to add a new module.
- If there is a super important file you should always collect on a Red Team engagement, add it via a PR.
- If its a file not under `/home` or `/root` then add it in `internals/systemCollector.go`
- If its a file typically found in `/root` or a users home direcotry add it in `internals/userCollector.go`
- For example if you wanted to always collect `/etc/fstab` you could easily find the correct function and add it in
![image](https://github.com/user-attachments/assets/83329069-e2e0-4472-9963-fd00e060ecaf)

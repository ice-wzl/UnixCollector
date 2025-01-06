# UnixCollector
![Gemini_Generated_Image_cgy691cgy691cgy6(1)](https://github.com/user-attachments/assets/aeb9327d-5814-49c5-86d7-9d3c29d4c9e4)

## Overview
- UnixCollector is a simple and lightweight CLI tool designed for rapid collection, and optional exfiltration of sensitive information from Linux systems.
- Its primary purpose is to streamline the process of gathering critical data in red teaming scenarios.

## Building
- Ensure you have golang installed on your host
- To install golang see: https://go.dev/doc/install 
````
git clone https://github.com/ice-wzl/UnixCollector.git
cd UnixCollector
go build
````
- `UnixCollector` takes no arguments or flags, it knows what to do!

## Benifits 
My origional inspiration for this project was https://github.com/R3DRUN3/vermilion
After attempting a PR, it was determined that we had different design ideas, thus I wanted to make my own version
1. `UnixCollector` will rebuild the directory structure of the host. For example if you collect `/etc/fstab` you will see that same directory struture reflected in your exfil directory
  - See below for a visual example of the rebuilt directory structure. This makes it much easier to analyze the files you have collected, as you know the exact position they are on the remote host
![image](https://github.com/user-attachments/assets/2515b50c-9e22-4dab-854d-94cbb75b2ad8)
2. Creates an archive for easy exfiltration
  - When the tool finishes its run, it will `tar/gzip` the contents to make the file transfer process as easy as possible.
 

## Contributing
- Im always happy to accept PRs!
- Ive attempted to make this code as module as possible to make it as easy as possible to add a new module.
- If there is a super important file you should always collect on a Red Team engagement, add it via a PR.
- If its a file not under `/home` or `/root` then add it in `internals/systemCollector.go`
- If its a file typically found in `/root` or a users home direcotry add it in `internals/userCollector.go`
- For example if you wanted to always collect `/etc/fstab` you could easily find the correct function and add it in
![image](https://github.com/user-attachments/assets/83329069-e2e0-4472-9963-fd00e060ecaf)

## FAQ
1. Should you look to automate all of this through `Sliver` or `Havoc`?
  - Yes that would be nice, however sometimes you dont have an oppertunity to use those tools. If you have a way to automate collection of sensitive files through `Sliver` or `Havoc` do that, however for those times you dont, `UnixCollector` is here for you.
2. This program doesnt really do any host enumeration, why?
  - This tool is designed to collect sensitive files from a linux system. If you are looking to priv esc or enumerate the host, use `linpeas`. Its 100x better than anything I could write for host enumeration. Different tools for different jobs.

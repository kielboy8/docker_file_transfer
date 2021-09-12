# Readme Instructions

This Go app deploys 2 docker containers that can transfer files from one container to another. The transfer is one-way, so 1 container is the sender and the other is the receiver.

## Requirements:

- You must have Docker installed on your machine.
- Make sure you have no networks in Docker using subnet range 10.0.0.0/16 (the application uses this)
	- If so you will have to modify the assigned IP addresses to the containers in `docker-compose.yml` and the `SENDER_IP` environment variable in both containers inside the mentioned YAML file.

## Running the app

Make sure that the deploy script has executable permissions (`sudo chmod +x deploy.sh`). To deploy and run the application simply run `./deploy.sh`.

The containers will then be deployed. Once the deployment script deploys the containers, the containers will run the file transfer on start. It will print out the logs in the end showing the result of the file transfer and terminate the containers.

### Output

```
...
sender-app_1    | Showing the initial directory:
sender-app_1    | bin      go.mod   main.go
sender-app_1    | Creating a file file.txt
sender-app_1    | bin       file.txt  go.mod    main.go
sender-app_1    | Listening to address: 10.0.0.2
sender-app_1    | Sending file: file.txt
sender-app_1    | file.txt successfully sent!
receiver-app_1  | Showing the initial directory:
receiver-app_1  | bin      go.mod   main.go
receiver-app_1  | Successfully dialed to IP: 10.0.0.2
receiver-app_1  | File file.txt received!
receiver-app_1  | bin       file.txt  go.mod    main.go
```

As shown in the logs above, the `sender-app` did not have `file.txt` initially so it created the file and ran its executable program to listen to the given IP address. Then, the `receiver-app` started and it did not originally have `file.txt`. It ran the executable program to dial in to the given IP address to receive the file. 

`sender-app` then sent the file `file.txt` over to `receiver-app`. Looking at `receiver-app` again, we can now see `file.txt` appear in its directory.

To down or terminate the containers, run `docker-compose down`.

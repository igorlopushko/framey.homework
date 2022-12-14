# Speed test API tool
## Description
GO library tests the download and upload speeds by using Ookla's https://www.speedtest.net/ and Netflix's https://fast.com/.  
## Configuration
Configuration is done through the environment variables. Supported configuration parameters:  
- ```LOG_LEVEL``` - determines the logs level of the [logrus](https://github.com/sirupsen/logrus) logger. Could have values such as  ```panic```, ```fatal```, ```error```,  ```warn```, ```info```, ```debug```, ```trace```. By default it is set to ```debug```.
## Usage
### Command line
Direct usage could be done with the following command:  

```go run main.go --provider ookla```  

```--provider``` (shot version is ```-p```) parameter could take the following values:
- ```ookla``` (shot version is ```o```) - Ookla's provider.
- ```netflix``` (shot version is ```n```) - Netflix's provider.
### Makefile
Use makefile for a quick run of the program. There are the following rules:
- ```run-all``` - runs tests for both Ookla's and Netflix's providers.  
- ```run-ookla``` - runs test for Ookla's provider.  
- ```run-netflix``` - runs test for Netflix's provider.  
- ```test``` - runs unit-tests.  
- ```bench``` - runs benchmark tests.
- ```godoc``` - generates docs and launches it on http://localhost:6060.  
- ```lint``` - runs linting.  
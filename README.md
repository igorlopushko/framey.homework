# Speed test API tool
## Description
GO library that tests the download and upload speeds by using Ookla's https://www.speedtest.net/ and Netflix's https://fast.com/.

## Usage
### Command line
Direct usage could be done with the following command:  
```go run main.go --provider speedtest.net```  
```--provider``` parameter could take two values:  
- ```speedtest.net```
- ```fast.com```

Short version is of this parameter is ```-p```
### Makefile
Use makefile for quick run of the program. There are following rule:
- ```run-all``` - runs tests for both ```speedtest.net``` and ```fast.com``` providers.  
- ```run-speed``` - runs test for ```speedtest.net``` provider.  
- ```run-fast``` - runs test for ```fast.com``` provider.  
- ```test``` - runs unit-tests.  
- ```bench``` - runs benchmark tests.
- ```godoc``` - generates docs and launches it on http://localhost:6060.  
- ```lint``` - runs linting.  
## Configuration
Configuration is done through the environment variables. Supported configuration parameters:  
- ```LOG_LEVEL``` - determines the logs level of the [logrus](https://github.com/sirupsen/logrus) logger. Could have values as  ```panic```, ```fatal```, ```error```,  ```warn```, ```info```, ```debug```, ```trace```. By default it is set to ```debug```.
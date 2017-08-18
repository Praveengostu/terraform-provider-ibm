# Bluemix SDK for Go

## Installing

1. Install the SDK using the following command

    ```go get github.com/IBM-Bluemix/bluemix-go```

2. Update the SDK to the latest version using the following command

    ```go get -u github.com/IBM-Bluemix/bluemix-go```


## Using the SDK

You must a have working Bluemix account to use the APIs. [Sign up](https://console.ng.bluemix.net/registration/?target=%2Fdashboard%2Fapps) if you don't have one.

The SDK has ```examples``` folder which cites few examples on how to use the SDK.
First you need to create a session.

```
import "github.com/IBM-Bluemix/bluemix-go/session"

   func main(){

    s := session.New()

   }
   .....
```
Creating session in this way creates a default configuration which reads the value from the environment variables.
You must export the following environment variables.
* IBMID - This is the IBM ID
* IBMID_PASSWORD - This is the password for the above ID

OR

* BM_API_KEY/BLUEMIX_API_KEY - This is the Bluemix API Key. Login to [Bluemix](https://console.ng.bluemix.net) to create one if you don't already have one. Follow Manage -> Account -> Users. Click on _Bluemix API Keys_

The default region is _us_south_. You can override it in the Config struct. You can also provide the value via environment variables; either via _BM_REGION_ or _BLUEMIX_REGION_. Valid regions are -
* us-south
* eu-gb
* eu-de
* au-syd



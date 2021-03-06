[![RetailOps SDK](http://cdn2.hubspot.net/hubfs/530512/Image/logo.png)](http://retailops.com)

### RetailOps SDK
----
The RetailOps SDK provides tools to assist you with verifying that your integration service endpoints are returning JSON data that is correctly formatted to meet the requirements of the SDK schema. The verify utility will validate your service's JSON responses by comparing them directly against the [RetailOps Swagger Schema](http://gudtech.github.io/retailops-sdk/v1/channel). The verify utility will also test that your service only responds when sent a valid [Integration Auth Key](#managing-your-integration-auth-key).

The following instructions will aid you in setting up your local environment, installing the SDK verify utility and running the verify utility against your service endpoints.     

 1. Download the `Verify Service` release for your operating system here: [RetailOps SDK Releases Page](https://github.com/gudTECH/retailops-sdk/releases)
 2. Unzip downloaded file
 3. Use terminal and enter unzipped directory (for example: `verify_windows_v0.0.9.0`)
 4. Run the verify utility against your actual endpoints by specifying the '-base-url' flag

    ```
    ./verify -base-url 'http://youractualserver.com/'
    ```

After you have completely developed your channel integration, and have successfully used
the verifier utility to test that your integration is operating correctly, you are ready to attempt certification with RetailOps.

Follow the instructions here: [Certifying Your RetailOps SDK Channel Integration](https://github.com/gudTECH/retailops-sdk/blob/master/verify/CERTIFY_README.md)

Managing Your Integration Auth Key
---

The integration auth key is used to confirm data is coming from RetailOps. This key must be kept secret and must be checked on every request from RetailOps. A bad integration auth key value must result in a HTTP 401. The integration auth key is a long, random string you will generate using the verify utility.

The verify utility helps manage this key. By default it is stored in your `$HOME/.retailops/` folder. When collaborating with other developers you may need to share and install this key.

To generate a new key, required for every service:

```
$ ./verify.exe generate-auth-key
$ ./verify.exe show-auth-key

==== INTEGRATION AUTH KEY BELOW ====
dDkma7W9zA/Nd7r3AUcEI/9MA9KWg+Uc
==== INTEGRATION AUTH KEY DONE ====
```

To install a different auth key, setting it to whatever value you choose:

```
$ ./verify.exe install-auth-key FOOBARBAZ
$ ./verify.exe show-auth-key

==== INTEGRATION AUTH KEY BELOW ====
FOOBARBAZ
==== INTEGRATION AUTH KEY DONE ====
```

> _Note: the integration auth key generated here is only for verification during use of a production service. This is the key that RetailOps passes to your service to verify that it's RetailOps making an API call to your server. The verification utility will use the test integration_auth_key: "RETAILOPS_SDK". During testing and verification your server should allow the test key, once in production it should only allow API calls made with the actual integration key that you've generated._

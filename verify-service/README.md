[![RetailOps SDK](http://cdn2.hubspot.net/hubfs/530512/Image/logo.png)](http://retailops.com)

### RetailOps SDK
----

Self-verify your RetailOps integration with this SDK.

 1. Download the `Verify Service` release for your operating system here: [RetailOps SDK Releases Page](https://github.com/gudTECH/retailops-sdk/releases)
 2. Unzip downloaded file
 3. Use terminal and enter unzipped directory (e.g., `verify-service_windows`)
 4. Start the example web service:

```
cd retailops-asp-dotnet-api-example
dnu restore
dnx web
```

 5. From another terminal run the `verify-service` utility:

```
verify-service-windows-amd64-2016-05-03-303ff.exe -schema-path schema/schemata/
```
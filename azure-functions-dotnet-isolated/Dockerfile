FROM mcr.microsoft.com/dotnet/sdk:7.0 AS installer-env

# Add the line below and it'll build successfully
COPY --from=mcr.microsoft.com/dotnet/sdk:7.0 /usr/share/dotnet /usr/share/dotnet

COPY . /src/dotnet-function-app

RUN cd /src/dotnet-function-app && \
    mkdir -p /home/site/wwwroot && \
    dotnet publish <your AzFunction.DotNet porject>/*.csproj --output /home/site/wwwroot

# To enable ssh & remote debugging on app service change the base image to the one below
FROM mcr.microsoft.com/azure-functions/dotnet-isolated:4-dotnet-isolated7.0

ENV AzureWebJobsScriptRoot=/home/site/wwwroot \
    AzureFunctionsJobHost__Logging__Console__IsEnabled=true \
    AzureWebJobsStorage="AccountName=devstoreaccount1;AccountKey=Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;DefaultEndpointsProtocol=http;BlobEndpoint=http://azurite:10000/devstoreaccount1;QueueEndpoint=http://azurite:10001/devstoreaccount1;TableEndpoint=http://azurite:10002/devstoreaccount1;" \
    AzureWebJobsSecretStorageType=files \
    FUNCTIONS_SECRETS_PATH=/etc/secrets \
    ASPNETCORE_URLS=https://+ \
    ASPNETCORE_Kestrel__Certificates__Default__Path=/certs/cert.pfx \
    FUNCTIONS_LOG_PATH=/azfunc-logs \
    TZ=Asia/Shanghai \
    DEBIAN_FRONTEND=noninteractive

RUN ln -fs /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && dpkg-reconfigure --frontend noninteractive tzdata \
    && rm -rf /var/lib/apt/lists/*

COPY --from=installer-env ["/home/site/wwwroot", "/home/site/wwwroot"]

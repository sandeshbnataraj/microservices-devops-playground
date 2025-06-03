FROM mcr.microsoft.com/dotnet/sdk:8.0

WORKDIR /app

COPY . ./

RUN dotnet publish -c Release

CMD ["dotnet", "run"]

#RUN dotnet publish -c Release -o app
#CMD ["dotnet", "app/cartservice.dll"]
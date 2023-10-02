FROM mcr.microsoft.com/dotnet/aspnet:7.0 AS base
WORKDIR /app
EXPOSE 80

FROM mcr.microsoft.com/dotnet/sdk:7.0 AS build
WORKDIR /src
COPY ["ApiServer/ApiServer.csproj", "ApiServer/"]
RUN dotnet restore "ApiServer/ApiServer.csproj"
COPY . .
WORKDIR "/src/ApiServer"
RUN dotnet build "ApiServer.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "ApiServer.csproj" -c Release -o /app/publish /p:UseAppHost=false

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "ApiServer.dll"]
# См. статью по ссылке https://aka.ms/customizecontainer, чтобы узнать как настроить контейнер отладки и как Visual Studio использует этот Dockerfile для создания образов для ускорения отладки.


FROM mcr.microsoft.com/dotnet/aspnet:8.0 AS base
USER $APP_UID
WORKDIR /app
EXPOSE 5555



FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
ARG BUILD_CONFIGURATION=Release
WORKDIR /src
COPY ["BookServiceCS/BookServiceCS.Api.csproj", "BookServiceCS/"]
COPY ["BookServiceCS.Application/BookServiceCS.Application.csproj", "BookServiceCS.Application/"]
COPY ["BookServiceCS.Contracts/BookServiceCS.Contracts.csproj", "BookServiceCS.Contracts/"]
COPY ["BookServiceCS.DataAccess/BookServiceCS.DataAccess.csproj", "BookServiceCS.DataAccess/"]
COPY ["BookServiceCS.Domain/BookServiceCS.Domain.csproj", "BookServiceCS.Domain/"]
RUN dotnet restore "./BookServiceCS/BookServiceCS.Api.csproj"
COPY . .
WORKDIR "/src/BookServiceCS"
RUN dotnet build "./BookServiceCS.Api.csproj" -c $BUILD_CONFIGURATION -o /app/build


FROM build AS publish
ARG BUILD_CONFIGURATION=Release
RUN dotnet publish "./BookServiceCS.Api.csproj" -c $BUILD_CONFIGURATION -o /app/publish /p:UseAppHost=false


FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "BookServiceCS.Api.dll"]
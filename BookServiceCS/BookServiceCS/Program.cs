using BookServiceCS.Application.Extensions;
using BookServiceCS.DataAccess.Extensions;
using BookServiceCS.Services;
using BookServiceCS.Filters;

var builder = WebApplication.CreateBuilder(args);


builder.Services.AddGrpc(options =>
{
    options.Interceptors.Add<GrpcExceptionInterceptor>();
});
builder.Services.AddScoped<GrpcExceptionInterceptor>();


builder.Services
    .AddServices()
    .AddDapper(builder.Configuration)
    .AddRepositories()
    .AddFluentMigrator(builder.Configuration)
    .MigrateUp();

var app = builder.Build();

app.MapGrpcService<BookGrpcService>();
app.MapGet("/", () => "gRPC Book Service. Use a gRPC client to communicate.");

app.Run();

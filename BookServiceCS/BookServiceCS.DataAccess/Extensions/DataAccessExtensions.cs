using BookServiceCS.Contracts.Dapper;
using BookServiceCS.Contracts.Settings;
using BookServiceCS.DataAccess.Dapper;
using BookServiceCS.DataAccess.Repositories;
using BookServiceCS.DataAccess.Migrations;
using BookServiceCS.Domain.Repositories;
using FluentMigrator.Runner;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;

namespace BookServiceCS.DataAccess.Extensions
{
    public static class DataAccessExtensions
    {
        public static IServiceCollection AddDapper(this IServiceCollection services, IConfiguration configuration)
        {
            services.AddScoped<IDapperContext, DapperContext>();

            services.Configure<DapperSettings>(o =>
            {
                o.ConnectionString = configuration.GetConnectionString("DefaultConnection")
                    ?? throw new InvalidOperationException("Connection string 'DefaultConnection' not found.");
            });

            return services;
        }

        public static IServiceCollection AddRepositories(this IServiceCollection services)
        {
            services.AddScoped<IBookRepository, BookRepository>();
            return services;
        }

        public static IServiceCollection AddFluentMigrator(this IServiceCollection services, IConfiguration configuration)
        {
            services
                .AddFluentMigratorCore()
                .ConfigureRunner(builder => builder
                    .AddPostgres()
                    .WithGlobalConnectionString(configuration.GetConnectionString("DefaultConnection")
                        ?? throw new InvalidOperationException("Connection string 'DefaultConnection' not found."))
                    .ScanIn(typeof(M001_Initial).Assembly).For.Migrations())
                .AddLogging(lb => lb.AddFluentMigratorConsole())
                .BuildServiceProvider(false);

            return services;
        }

        public static IServiceCollection MigrateUp(this IServiceCollection services)
        {
            using (var scope = services.BuildServiceProvider().CreateScope())
            {
                var runner = scope.ServiceProvider.GetRequiredService<IMigrationRunner>();
                runner.ListMigrations();
                runner.MigrateUp();
            }

            return services;
        }
    }
}



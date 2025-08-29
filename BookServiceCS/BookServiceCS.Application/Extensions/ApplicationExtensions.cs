using Microsoft.Extensions.DependencyInjection;
using BookServiceCS.Application.Services;
using BookServiceCS.Contracts.Services;

namespace BookServiceCS.Application.Extensions;

public static class ApplicationExtensions
{
    public static IServiceCollection AddServices(this IServiceCollection services)
    {
        services.AddScoped<IBookService, BookService>();
        return services;
    }
}



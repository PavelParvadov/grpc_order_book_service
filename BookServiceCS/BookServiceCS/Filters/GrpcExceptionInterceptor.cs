using BookServiceCS.Domain.Exceptions.Common;
using Grpc.Core;
using Grpc.Core.Interceptors;

namespace BookServiceCS.Filters;

public class GrpcExceptionInterceptor : Interceptor
{
    private readonly ILogger<GrpcExceptionInterceptor> _logger;

    public GrpcExceptionInterceptor(ILogger<GrpcExceptionInterceptor> logger)
    {
        _logger = logger;
    }

    public override async Task<TResponse> UnaryServerHandler<TRequest, TResponse>(
        TRequest request,
        ServerCallContext context,
        UnaryServerMethod<TRequest, TResponse> continuation)
    {
        try
        {
            return await continuation(request, context);
        }
        catch (Exception ex)
        {
            _logger.LogError(ex, "Exception occured: {Message}\n\tEndpoint: {Path}", ex.Message, context.Method);

            var status = ex switch
            {
                EntityNotFoundException => new Status(StatusCode.NotFound, ex.Message),
                EntityAlreadyExistsException => new Status(StatusCode.AlreadyExists, ex.Message),
                _ => new Status(StatusCode.Internal, "Internal server error")
            };

            throw new RpcException(status);
        }
    }
}



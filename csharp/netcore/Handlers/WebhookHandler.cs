namespace csharp.netcore.Handlers;

public class PostProcessWebhook : ResponseHandler<ConcealRequest, ConcealResponse>
{
    public async override Task<ConcealResponse> Handle(HelloRequest request)
    {
        return new HelloResponse {
            Greeting = $"Hello, {request.Name}"
        };
    }
}
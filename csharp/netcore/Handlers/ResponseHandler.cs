using Microsoft.AspNetCore.Mvc;

namespace csharp.netcore.Handlers;

[ApiController, Route("{prefix:webhookRoutePrefix}/[controller]")]
public abstract class ResponseHandler<TRequest, TResponse>
{
    [HttpPost, Route("api-key-signature-protected")]
    public abstract Task<TResponse> Handle([FromBody] TRequest request);
}
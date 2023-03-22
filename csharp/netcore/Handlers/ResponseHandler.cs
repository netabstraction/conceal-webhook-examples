using Microsoft.AspNetCore.Mvc;

namespace csharp.netcore.Handlers;

[ApiController, Route("[controller]")]
public abstract class ResponseHandler<TRequest, TResponse>
{
    [HttpPost, Route("")]
    public abstract Task<TResponse> Handle([FromBody] TRequest request);
}
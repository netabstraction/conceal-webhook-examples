using Microsoft.AspNetCore.Mvc;

namespace csharp.netcore.Handlers;

[ApiController, Route("[controller]")]
public abstract class ResponseHandler<TRequest, TResponse>
{
    [HttpPost, Route("")]
    public abstract Task<TResponse> Handle([FromBody] TRequest request, [FromHeader(Name = "X-Api-Key")] string apiKey, [FromHeader(Name = "X-Signature")] string signature, [FromHeader(Name = "X-TimeStamp")] long timeStamp);
}
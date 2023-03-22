using Microsoft.AspNetCore.Mvc;

namespace csharp.netcore.Handlers;

[ApiController, Route("[controller]")]
public abstract class ResponseHandler<TRequest, TResponse>
{
    [HttpPost, Route("")]
    public abstract Task<TResponse> Handle([FromBody] TRequest request,
                                           [FromHeader(Name = "x-api-key")] string apiKey,
                                           [FromHeader(Name = "conceal-signature")] string signature,
                                           [FromHeader(Name = "conceal-timeStamp")] long timeStamp);
}
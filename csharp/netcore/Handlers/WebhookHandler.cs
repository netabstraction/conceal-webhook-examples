using csharp.netcore.Models;
using csharp.netcore.Validators;
using Microsoft.AspNetCore.Mvc;

namespace csharp.netcore.Handlers;

[Route("api-key-signature-protected")]
public class PostProcessWebhook : ResponseHandler<ConcealRequest, ConcealResponse>
{
    private static SignatureValidator signatureValidator = new SignatureValidator();
    private static TimeStampValidator timeStampValidator = new TimeStampValidator();
    private static ApiKeyValidator apiKeyValidator = new ApiKeyValidator();
    
    [HttpPost]
    public override Task<ConcealResponse> Handle(
        [FromBody] ConcealRequest request,
        [FromHeader(Name = "X-Api-Key")] string apiKey, 
        [FromHeader(Name = "X-Signature")] string signature, 
        [FromHeader(Name = "X-TimeStamp")] long timeStamp)
    {
        throw new NotImplementedException();
    }
}
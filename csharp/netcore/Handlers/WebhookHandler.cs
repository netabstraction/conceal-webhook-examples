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
    public override async Task<ConcealResponse> Handle(
        [FromBody] ConcealRequest request,
        [FromHeader(Name = "x-api-key")] string apiKey,
        [FromHeader(Name = "conceal_signature")] string signature,
        [FromHeader(Name = "conceal_timestamp")] long timeStamp)
    {


        if (!apiKeyValidator.Validate(apiKey, signature, timeStamp))
        {
            return await Task.FromResult(new ConcealResponse("API Key missing/API Key doesnot match", 401));
        }

        if (!timeStampValidator.Validate(apiKey, signature, timeStamp))
        {
            return await Task.FromResult(new ConcealResponse("Invalid Timestamp. Timestamp not in range", 400));
        }

        if (!signatureValidator.Validate(apiKey, signature, timeStamp))
        {
            return await Task.FromResult(new ConcealResponse("Invalid Signature", 400));
        }

        // DEMO: Print the payload
        return await Task.FromResult(new ConcealResponse(request.ToString(), 200));

    }
}
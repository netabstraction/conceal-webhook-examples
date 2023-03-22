using csharp.netcore.Models;
using csharp.netcore.Validators;
using Microsoft.AspNetCore.Mvc;

namespace csharp.netcore.Handlers;

[Route("api-key-signature-protected")]
public class PostProcessWebhook : ResponseHandler<ConcealRequest, ConcealResponse>
{


    [HttpPost]
    public override async Task<ConcealResponse> Handle(
        [FromBody] ConcealRequest request)
    {
        // DEMO: Print the payload
        return await Task.FromResult(new ConcealResponse(request.ToString(), 200));

    }
}
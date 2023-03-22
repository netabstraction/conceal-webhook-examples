using csharp.netcore.Models;
using csharp.netcore.Validators;

namespace csharp.netcore.Middleware;

public class Middleware
{
    private static SignatureValidator signatureValidator = new SignatureValidator();
    private static TimeStampValidator timeStampValidator = new TimeStampValidator();
    private static ApiKeyValidator apiKeyValidator = new ApiKeyValidator();

    private readonly RequestDelegate _next;

    public Middleware(RequestDelegate next)
    {
        _next = next;
    }

    public async Task Invoke(HttpContext context)
    {
        String apiKey = context.Request.Headers["x-api-key"];
        String signature = context.Request.Headers["conceal-signature"];
        long timeStamp = long.Parse(context.Request.Headers["conceal-timestamp"]);

        if (!apiKeyValidator.Validate(apiKey, signature, timeStamp))
        {
            context.Response.StatusCode = 401;
            await context.Response.WriteAsync("API Key missing/API Key doesnot match");
            return;
        }

        if (!timeStampValidator.Validate(apiKey, signature, timeStamp))
        {
            context.Response.StatusCode = 400;
            await context.Response.WriteAsync("Invalid Timestamp. Timestamp not in range");
            return;
        }

        if (!signatureValidator.Validate(apiKey, signature, timeStamp))
        {
            context.Response.StatusCode = 400;
            await context.Response.WriteAsync("Invalid Signature");
            return;
        }

        await _next(context);
    }
}
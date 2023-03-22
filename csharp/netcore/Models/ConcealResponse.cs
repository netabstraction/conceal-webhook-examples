namespace csharp.netcore.Models;

public class ConcealResponse
{
    public ConcealResponse()
    {
    }

    public ConcealResponse(string? message, int? status)
    {
        this.message = message;
        this.status = status;
    }

    public string? message { get; set; }
    public int? status { get; set; }

    public override String ToString()
    {
        return $"{{\"message\": \"{message}\", \"status\": {status}}}";
    }
}
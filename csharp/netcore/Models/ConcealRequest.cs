namespace csharp.netcore.Models;

public class ConcealRequest
{
    public ConcealRequest()
    {
    }

    public string Event { get; set; }
    public string host { get; set; }
    public string sourceType { get; set; }
    public string compantId { get; set; }
    public string companyName { get; set; }
    public string userEmail { get; set; }
    public string userId { get; set; }
    public string url { get; set; }
    public int count { get; set; }
    public Decision decision { get; set; }
    public bool finalDecision { get; set; }
    public long timeStamp { get; set; }

    public override bool Equals(object? obj)
    {
        return obj is ConcealRequest request &&
               Event == request.Event &&
               host == request.host &&
               sourceType == request.sourceType &&
               compantId == request.compantId &&
               companyName == request.companyName &&
               userEmail == request.userEmail &&
               userId == request.userId &&
               url == request.url &&
               count == request.count &&
               EqualityComparer<Decision>.Default.Equals(decision, request.decision) &&
               finalDecision == request.finalDecision &&
               timeStamp == request.timeStamp;
    }

    public override int GetHashCode()
    {
        HashCode hash = new HashCode();
        hash.Add(Event);
        hash.Add(host);
        hash.Add(sourceType);
        hash.Add(compantId);
        hash.Add(companyName);
        hash.Add(userEmail);
        hash.Add(userId);
        hash.Add(url);
        hash.Add(count);
        hash.Add(decision);
        hash.Add(finalDecision);
        hash.Add(timeStamp);
        return hash.ToHashCode();
    }
}
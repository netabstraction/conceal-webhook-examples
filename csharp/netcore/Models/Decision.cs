namespace csharp.netcore.Models;

public class Decision
{
    public Decision()
    {
    }

    public bool enforceTls { get; set; }
    public bool noIp { get; set; }

    public override bool Equals(object? obj)
    {
        return obj is Decision decision &&
               enforceTls == decision.enforceTls &&
               noIp == decision.noIp;
    }

    public override int GetHashCode()
    {
        return HashCode.Combine(enforceTls, noIp);
    }
}
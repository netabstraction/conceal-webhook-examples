namespace csharp.netcore.Validators;

public class TimeStampValidator : IValidator
{
    private static long TIME_STAMP_TOLERANCE = 300000;

    public bool Validate(String apiKey, String signature, long timeStamp)
    {
        long currentTime = DateTimeOffset.Now.ToUnixTimeMilliseconds() / 1000;

        if (timeStamp - currentTime < -60000 ||
            timeStamp - currentTime > 120000)
        {
            return false;
        }

        return true;
    }
}
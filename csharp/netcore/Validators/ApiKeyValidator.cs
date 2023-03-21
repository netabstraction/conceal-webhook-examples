namespace csharp.netcore.Validators;

public class ApiKeyValidator : IValidator
{

    private static String API_VALUE_CONSTANT = "sample-key";

    public bool Validate(String apiKey, String signature, long timeStamp)
    {
        if (apiKey == null || !apiKey.Equals(API_VALUE_CONSTANT))
        {
            return false;
        }

        return true;
    }
}
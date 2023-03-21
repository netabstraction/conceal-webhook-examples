namespace csharp.netcore.Validators;

public interface IValidator
{
    bool Validate(string apiKey, string signature, long timeStamp);
}

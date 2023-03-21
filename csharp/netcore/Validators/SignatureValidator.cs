using System.Security.Cryptography;
using System.Text;

namespace csharp.netcore.Validators;

public class SignatureValidator : IValidator
{
    private static String SIGNATURE_KEY_CONSTANT = "signature-key";
    private static String WEB_HOOK_URL = "http://localhost:8080/api-key-signature-protected";

    public bool Validate(String apiKey, String signature, long timeStamp)
    {
        String messasge = timeStamp + WEB_HOOK_URL;
        byte[] key = Encoding.ASCII.GetBytes(SIGNATURE_KEY_CONSTANT);


        using (HMACSHA256 hmac = new HMACSHA256(key))
        {
            String expextedSignature = HmacSHA256(SIGNATURE_KEY_CONSTANT, messasge);

            if (signature != null && signature.Equals(expextedSignature))
            {
                return false;
            }
        }

        return true;
    }

    private string HmacSHA256(string key, string data)
    {
        string hash;
        ASCIIEncoding encoder = new ASCIIEncoding();
        Byte[] code = encoder.GetBytes(key);
        using (HMACSHA256 hmac = new HMACSHA256(code))
        {
            Byte[] hmBytes = hmac.ComputeHash(encoder.GetBytes(data));
            hash = ToHexString(hmBytes);
        }
        return hash;
    }

    private static string ToHexString(byte[] array)
    {
        StringBuilder hex = new StringBuilder(array.Length * 2);
        foreach (byte b in array)
        {
            hex.AppendFormat("{0:x2}", b);
        }
        return hex.ToString();
    }
}
using System.Text.RegularExpressions;

namespace csharp;

public static class Toolbox
{
    private static Boolean IsAlphanumeric(char character)
    {
        Regex rg = new Regex(@"^[a-zA-Z0-9\s,]*$");
        return rg.IsMatch(character.ToString());
    }

    public static int ConvertToPriorityValue(char character)
    {
        if (!(IsAlphanumeric(character)))
        {
            throw new Exception($"Character {character} is not alphanumeric");
        }
        // Implicitly casts character to UTF-8 integer code value
        int characterValue;
        if (Char.IsUpper(character))
        {
            characterValue = character;
            return characterValue += 27;
        }
        characterValue = character;
        return characterValue += 1;
    }
}


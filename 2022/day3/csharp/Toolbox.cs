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

    public static List<List<char>> SplitIntoCompartments(String RucksackContents)
    {
        if (RucksackContents.Length < 1)
        {
            throw new Exception($"Rucksack contents to short: {RucksackContents}");
        }



        return new List<List<char>>
        {
          RucksackContents.Substring(0, (int)(RucksackContents.Length / 2)).ToList(),
          RucksackContents.Substring((int)(RucksackContents.Length /2), (int)(RucksackContents.Length / 2)).ToList(),
        };
    }

    public static char GetIntersectionCharacter(List<char> firstCompartment, List<char> secondCompartment)
    {
        return firstCompartment.Intersect(secondCompartment).First();
    }
}


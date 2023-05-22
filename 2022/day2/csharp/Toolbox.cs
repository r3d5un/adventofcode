using csharp.Choices;

namespace csharp;

public static class Toolbox
{
    private static IChoice SubstituteValue(String value)
    {
        if (value.Equals("A") | value.Equals("X"))
        {
            return new Rock();
        }
        if (value.Equals("B") | value.Equals("Y"))
        {
            return new Paper();
        }
        if (value.Equals("C") | value.Equals("Z"))
        {
            return new Scissors();
        }

        throw new InvalidDataException($"Given value not valid: {value}");
    }

    public static PlayerChoices GetPlayerChoices(String line)
    {
        IList<String> values = line.Split(" ");
        return new PlayerChoices(
            playerOneChoice: SubstituteValue(values[0]),
            playerTwoChoice: SubstituteValue(values[1])
        );
    }

    public static IChoice EnforceStrategyGuide(IChoice opponentsChoice, String wantedOutcome)
    {
        switch (wantedOutcome)
        {
            case "X" when opponentsChoice.GetType() == typeof(Rock):
                return new Scissors();
            case "X" when opponentsChoice.GetType() == typeof(Paper):
                return new Rock();
            case "X" when opponentsChoice.GetType() == typeof(Scissors):
                return new Paper();
            case "Y":
                return opponentsChoice;
            case "Z" when opponentsChoice.GetType() == typeof(Rock):
                return new Paper();
            case "Z" when opponentsChoice.GetType() == typeof(Paper):
                return new Scissors();
            case "Z" when opponentsChoice.GetType() == typeof(Scissors):
                return new Rock();
            default:
                throw new InvalidDataException("Parameters outside expectations");
        }
    }
}

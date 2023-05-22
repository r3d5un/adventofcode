using System.Runtime.CompilerServices;
using csharp.Choices;

namespace csharp;

public static class Toolbox
{
    public static IChoice SubstituteValue(String value)
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
        if (wantedOutcome == "X")
        {
            if (opponentsChoice.GetType().IsInstanceOfType(new Rock()))
            {
                return new Scissors();
            }
            if (opponentsChoice.GetType().IsInstanceOfType(new Paper()))
            {
                return new Rock();
            }
            if (opponentsChoice.GetType().IsInstanceOfType(new Scissors()))
            {
                return new Paper();
            }
        }

        if (wantedOutcome == "Y")
        {
            return opponentsChoice;
        }

        if (wantedOutcome == "Z")
        {
            if (opponentsChoice.GetType().IsInstanceOfType(new Rock()))
            {
                return new Paper();
            }
            if (opponentsChoice.GetType().IsInstanceOfType(new Paper()))
            {
                return new Scissors();
            }

            if (opponentsChoice.GetType().IsInstanceOfType(new Scissors()))
            {
                return new Rock();
            }
        }

        throw new InvalidDataException("Parameters outside expectations");
    }
}

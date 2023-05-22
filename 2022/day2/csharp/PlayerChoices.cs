using csharp.Choices;

namespace csharp;

public class PlayerChoices : IPlayerChoices
{
    public IChoice PlayerOneChoice { get; set; }
    public IChoice PlayerTwoChoice { get; set; }

    public PlayerChoices(IChoice playerOneChoice, IChoice playerTwoChoice)
    {
        PlayerOneChoice = playerOneChoice;
        PlayerTwoChoice = playerTwoChoice;
    }
}

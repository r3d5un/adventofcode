using System.Reflection;
using csharp.Choices;

namespace csharp;

public class RockPaperScissorsRound : IRockPaperScissorsRound
{
    private IChoice PlayerOneChoice { get; set; }
    private IChoice PlayerTwoChoice { get; set; }
    public RoundScore Result { get; set; }

    public RockPaperScissorsRound(IChoice playerOneChoice, IChoice playerTwoChoice)
    {
        PlayerOneChoice = playerOneChoice;
        PlayerTwoChoice = playerTwoChoice;
        Result = PlayRound();
    }

    public RoundScore PlayRound()
    {
        if (PlayerOneChoice.GetType().IsInstanceOfType(new Rock()))
        {
            if (PlayerTwoChoice.GetType().IsInstanceOfType(new Paper()))
            {
                return new RoundScore(
                    playerOnePoints: (0 + PlayerOneChoice.Value),
                    playerTwoPoints: (6 + PlayerTwoChoice.Value)
                );
            }
            if (PlayerTwoChoice.GetType().IsInstanceOfType(new Scissors()))
            {
                return new RoundScore(
                    playerOnePoints: (6 + PlayerOneChoice.Value),
                    playerTwoPoints: (0 + PlayerTwoChoice.Value)
                );
            }
        }
        if (PlayerOneChoice.GetType().IsInstanceOfType(new Paper()))
        {
            if (PlayerTwoChoice.GetType().IsInstanceOfType(new Scissors()))
            {
                return new RoundScore(
                    playerOnePoints: (0 + PlayerOneChoice.Value),
                    playerTwoPoints: (6 + PlayerTwoChoice.Value)
                );
            }
            if (PlayerTwoChoice.GetType().IsInstanceOfType(new Rock()))
            {
                return new RoundScore(
                    playerOnePoints: (6 + PlayerOneChoice.Value),
                    playerTwoPoints: (0 + PlayerTwoChoice.Value)
                );
            }
        }
        if (PlayerOneChoice.GetType().IsInstanceOfType(new Scissors()))
        {
            if (PlayerTwoChoice.GetType().IsInstanceOfType(new Rock()))
            {
                return new RoundScore(
                    playerOnePoints: (0 + PlayerOneChoice.Value),
                    playerTwoPoints: (6 + PlayerTwoChoice.Value)
                );
            }
            if (PlayerTwoChoice.GetType().IsInstanceOfType(new Paper()))
            {
                return new RoundScore(
                    playerOnePoints: (6 + PlayerOneChoice.Value),
                    playerTwoPoints: (0 + PlayerTwoChoice.Value)
                );
            }
        }
        if (PlayerOneChoice.GetType().IsInstanceOfType(PlayerTwoChoice))
        {
            return new RoundScore(
                playerOnePoints: (3 + PlayerOneChoice.Value),
                playerTwoPoints: (3 + PlayerTwoChoice.Value)
            );
        }

        throw new Exception("No branches hit");
    }
}

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
        if (PlayerOneChoice.GetType() == typeof(Rock))
        {
            if (PlayerTwoChoice.GetType() == typeof(Paper))
            {
                return new RoundScore(
                    playerOnePoints: (0 + PlayerOneChoice.Value),
                    playerTwoPoints: (6 + PlayerTwoChoice.Value)
                );
            }
            if (PlayerTwoChoice.GetType() == typeof(Scissors))
            {
                return new RoundScore(
                    playerOnePoints: (6 + PlayerOneChoice.Value),
                    playerTwoPoints: (0 + PlayerTwoChoice.Value)
                );
            }
        }
        if (PlayerOneChoice.GetType() == typeof(Paper))
        {
            if (PlayerTwoChoice.GetType() == typeof(Scissors))
            {
                return new RoundScore(
                    playerOnePoints: (0 + PlayerOneChoice.Value),
                    playerTwoPoints: (6 + PlayerTwoChoice.Value)
                );
            }
            if (PlayerTwoChoice.GetType() == typeof(Rock))
            {
                return new RoundScore(
                    playerOnePoints: (6 + PlayerOneChoice.Value),
                    playerTwoPoints: (0 + PlayerTwoChoice.Value)
                );
            }
        }
        if (PlayerOneChoice.GetType() == typeof(Scissors))
        {
            if (PlayerTwoChoice.GetType() == typeof(Rock))
            {
                return new RoundScore(
                    playerOnePoints: (0 + PlayerOneChoice.Value),
                    playerTwoPoints: (6 + PlayerTwoChoice.Value)
                );
            }
            if (PlayerTwoChoice.GetType() == typeof(Paper))
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

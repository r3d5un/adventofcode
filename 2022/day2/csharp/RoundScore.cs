namespace csharp;

public class RoundScore : IRoundScore
{
    public int PlayerOnePoints { get; set; }
    public int PlayerTwoPoints { get; set; }

    public RoundScore(int playerOnePoints, int playerTwoPoints)
    {
        PlayerOnePoints = playerOnePoints;
        PlayerTwoPoints = playerTwoPoints;
    }
}

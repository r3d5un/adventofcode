namespace csharp.Choices;

public class Scissors : IChoice
{
    public int Value { get; }

    public Scissors()
    {
        Value = 3;
    }
}

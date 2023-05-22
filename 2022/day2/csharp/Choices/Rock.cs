namespace csharp.Choices;

public class Rock : IChoice
{
    public int Value { get; }

    public Rock()
    {
        Value = 1;
    }
}

using csharp.Choices;

namespace csharp;

public interface IPlayerChoices
{
    IChoice PlayerOneChoice { set; get; }
    IChoice PlayerTwoChoice { set; get; }
}

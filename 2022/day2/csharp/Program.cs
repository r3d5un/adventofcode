using static csharp.Toolbox;

using csharp;

IEnumerable<String> inputText = File.ReadLines(
    "/home/r3d5un/Development/Learning/adventofcode/2022/day2/input.txt"
);

int partOneSum = 0;
foreach (String line in inputText)
{
    if (!(line.Trim() == "\n" | line.Trim() == ""))
    {
        PlayerChoices playerChoices = GetPlayerChoices(line);
        RockPaperScissorsRound round = new RockPaperScissorsRound(
            playerChoices.PlayerOneChoice,
            playerChoices.PlayerTwoChoice
        );
        partOneSum += round.Result.PlayerTwoPoints;
    }
}
Console.WriteLine($"Part 1: {partOneSum}");

int partTwoSum = 0;
foreach (String line in inputText)
{
    if (!(line.Trim() == "\n" | line.Trim() == ""))
    {
        PlayerChoices playerChoices = GetPlayerChoices(line);
        playerChoices.PlayerTwoChoice = EnforceStrategyGuide(
            opponentsChoice: playerChoices.PlayerOneChoice,
            wantedOutcome: line.Trim().Split(" ")[1]
        );
        RockPaperScissorsRound round = new RockPaperScissorsRound(
            playerChoices.PlayerOneChoice,
            playerChoices.PlayerTwoChoice
        );
        partTwoSum += round.Result.PlayerTwoPoints;
    }
}
Console.WriteLine($"Part 2: {partTwoSum}");

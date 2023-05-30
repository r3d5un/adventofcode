using static csharp.Toolbox;

IEnumerable<String> inputText = File.ReadLines(
    "/home/r3d5un/Development/Learning/adventofcode/2022/day3/input.txt"
);

int sumPartOne = 0;
foreach (String line in inputText)
{
    List<List<char>> compartmentLists = SplitIntoCompartments(line);
    char item = GetIntersectionCharacter(compartmentLists[0], compartmentLists[1]);
    sumPartOne += ConvertToPriorityValue(item);
}
Console.WriteLine($"Part one sum: {sumPartOne}");

String inputText = File.ReadAllText("/home/r3d5un/Development/Learning/adventofcode/2022/day1/input.txt");

string[] calorieGroupArray = inputText.Split("\n\n");

List<int> caloriesPerElfList = new List<int> { };

foreach (String group in calorieGroupArray) {
    if (string.IsNullOrWhiteSpace(group))
        continue;

    List<String> calorieNumberList = new List<string>(group.Split("\n").ToList());

    int currentElfCalories = 0;
    foreach (String calorieString in calorieNumberList)
    {
        currentElfCalories += Convert.ToInt32(calorieString);
    }

    caloriesPerElfList.Add(currentElfCalories);
}

// Answer to Part 1
Console.WriteLine(caloriesPerElfList.Max());

// Answert to Part 2
int topThreeCalories = (from i in caloriesPerElfList
    orderby i descending
    select i).Take(3).Sum();
Console.WriteLine(topThreeCalories);

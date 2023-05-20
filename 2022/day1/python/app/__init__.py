from pathlib import Path

def main() -> None:
    input_file_path = Path("../input.txt")

    with open(input_file_path, "r") as f:
        input_text: str = f.read()

    calorie_group_list: list[str] = input_text.split("\n\n")

    calories_per_elf_list: list[int] = []

    for group in calorie_group_list:
        if group == '':
            continue

        calorie_number_list: list[str] = group.strip().split("\n")

        current_elf: int = 0
        for calorie in calorie_number_list:
            current_elf += int(calorie)
        calories_per_elf_list.append(current_elf)

    # Answer to Part 1
    print(max(calories_per_elf_list))

    calories_per_elf_list.sort(reverse=True)

    # Answer to Part 2
    print(sum(calories_per_elf_list[0:3]))

if __name__ == "__main__":
    main()


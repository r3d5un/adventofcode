def convert_to_priority_value(character: str) -> int:
    """
    Converts a character to a priority using it's unicode code.
    """
    if not character.isalpha():
        raise ValueError(f"Invalid character: {character}")

    if character.isupper():
        return ord(character) - ord("A") + 27
    return ord(character) - ord("a") + 1


def split_into_compartments(rucksack_contents: str) -> tuple:
    """
    Takes a string representing the contents of a rucksack and splits it
    into compartments as a tuple.
    """
    if len(rucksack_contents) < 1:
        raise ValueError("String to short")

    return set(rucksack_contents[: len(rucksack_contents) // 2]), set(
        rucksack_contents[len(rucksack_contents) // 2 :]
    )


def get_interserction_character(first_compartment: str, second_compartment: str) -> str:
    """
    Takes in two strings representing compartments, and returns any intersecting
    character.
    """
    return set(first_compartment).intersection(set(second_compartment)).pop()


def get_rucksack_groups(rucksacks: list[str], group_size: int):
    return (
        rucksacks[pos : pos + group_size]
        for pos in range(0, len(rucksacks), group_size)
    )


def get_badge_for_group(rucksack_group: list[str]) -> str:
    return (
        set(rucksack_group[0].strip())
        .intersection(set(rucksack_group[1].strip()))
        .intersection(set(rucksack_group[2].strip()))
        .pop()
    )


def main():
    with open("../input.txt", "r") as f:
        rucksacks: list = f.readlines()

    part_one_sum: int = 0

    for rucksack in rucksacks:
        compartment_one, compartment_two = split_into_compartments(
            rucksack_contents=rucksack.strip()
        )

        item = get_interserction_character(
            first_compartment=compartment_one, second_compartment=compartment_two
        )

        part_one_sum += convert_to_priority_value(item)

    print(f"Part one sum: {part_one_sum}")

    badge_sum: int = 0
    for rucksack_group in get_rucksack_groups(rucksacks=rucksacks, group_size=3):
        badge_sum += convert_to_priority_value(get_badge_for_group(rucksack_group))

    print(f"Part two sum: {badge_sum}")


if __name__ == "__main__":
    main()

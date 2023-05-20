use std::fs;

fn main() {
    let input_text = fs::read_to_string("../input.txt").expect("Failed to read the file.");
    let calorie_groups: Vec<&str> = input_text.split("\n\n").collect();
    // dbg!(calorie_groups);

    let mut calories_per_elf: Vec<i32> = Vec::new();
    for group in calorie_groups {
        if group.is_empty() {
            continue;
        }
        let calorie_numbers_for_elf: Vec<&str> = group.split("\n").collect();

        let mut current_elf_sum: i32 = 0;
        for number in calorie_numbers_for_elf {
            let parsed_number = number.parse::<i32>().unwrap();
            current_elf_sum += parsed_number;
        }
        calories_per_elf.push(current_elf_sum);
    }
    let highest_number_of_calories = calories_per_elf.iter().max().unwrap();

    // Answer to Part 1
    println!("{}", highest_number_of_calories);

    calories_per_elf.sort();
    let sum_of_top_three: i32 = calories_per_elf.iter().rev().take(3).sum();

    // Answer to Part 2
    println!("{}", sum_of_top_three);
}

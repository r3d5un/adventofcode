fn main() {
    println!("Hello, world!");
struct Choice {
    value: String,
    points: u8,
}

struct PlayerChoices {
    player_one_choice: Choice,
    player_two_choice: Choice,
}

fn get_player_choices(line: &str) -> PlayerChoices {
    let values: Vec<&str> = line.split(" ").collect();
    PlayerChoices {
        player_one_choice: substitue_value(values[0]),
        player_two_choice: substitue_value(values[1]),
    }
}

fn substitue_value(value: &str) -> Choice {
    match value {
        "A" | "X" => {
            return Choice {
                value: String::from("rock"),
                points: 1,
            }
        }
        "B" | "Y" => {
            return Choice {
                value: String::from("paper"),
                points: 2,
            }
        }
        "C" | "Z" => {
            return Choice {
                value: String::from("scissors"),
                points: 3,
            }
        }
        _ => panic!("No valid substitutions"),
    }
}
const DRAW: u8 = 3;
const WIN: u8 = 6;
const LOSE: u8 = 0;
}

use std::{fs, u8};

fn main() {
    let input_text = fs::read_to_string("../input.txt").expect("Failed to read the file");
    let lines: Vec<&str> = input_text.split("\n").collect();
    let mut part_one_sum: u32 = 0;
    for line in lines {
        if !line.is_empty() {
            dbg!(line);
            let player_choices: PlayerChoices = get_player_choices(line);
            let mut game = RockPaperScissors {
                player_one_choice: player_choices.player_one_choice,
                player_two_choice: player_choices.player_two_choice,
                result: RoundScores {
                    player_one_points: 0,
                    player_two_points: 0,
                },
            };
            game.play();
            dbg!(&game);
            part_one_sum += u32::try_from(game.result.player_two_points).unwrap();
        }
    }
    println!("Part 1 Sum: {}", part_one_sum);
}

#[derive(Debug)]
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

#[derive(Debug)]
struct RoundScores {
    player_one_points: u8,
    player_two_points: u8,
}

const DRAW: u8 = 3;
const WIN: u8 = 6;
const LOSE: u8 = 0;

#[derive(Debug)]
struct RockPaperScissors {
    player_one_choice: Choice,
    player_two_choice: Choice,
    result: RoundScores,
}

impl RockPaperScissors {
    fn play(&mut self) {
        if self.player_one_choice.value == self.player_two_choice.value {
            self.result = RoundScores {
                player_one_points: DRAW + self.player_one_choice.points,
                player_two_points: DRAW + self.player_two_choice.points,
            }
        } else if self.player_one_choice.value == "rock" {
            if self.player_two_choice.value == "paper" {
                self.result = RoundScores {
                    player_one_points: LOSE + self.player_one_choice.points,
                    player_two_points: WIN + self.player_two_choice.points,
                };
                return;
            }
            self.result = RoundScores {
                player_one_points: WIN + self.player_one_choice.points,
                player_two_points: LOSE + self.player_two_choice.points,
            };
            return;
        } else if self.player_one_choice.value == "paper" {
            if self.player_two_choice.value == "scissors" {
                self.result = RoundScores {
                    player_one_points: LOSE + self.player_one_choice.points,
                    player_two_points: WIN + self.player_two_choice.points,
                };
                return;
            }
            self.result = RoundScores {
                player_one_points: WIN + self.player_one_choice.points,
                player_two_points: LOSE + self.player_two_choice.points,
            };
            return;
        } else if self.player_one_choice.value == "scissors" {
            if self.player_two_choice.value == "rock" {
                self.result = RoundScores {
                    player_one_points: LOSE + self.player_one_choice.points,
                    player_two_points: WIN + self.player_two_choice.points,
                };
                return;
            }
            self.result = RoundScores {
                player_one_points: WIN + self.player_one_choice.points,
                player_two_points: LOSE + self.player_two_choice.points,
            };
            return;
        } else {
            panic!("Invalid choices");
        }
    }
}

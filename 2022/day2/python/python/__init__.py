from pathlib import Path
from dataclasses import dataclass


class Rock:
    def __init__(self) -> None:
        self.value = 1

    def __repr__(self) -> str:
        return f"<Rock(value={self.value})>"


class Paper:
    def __init__(self) -> None:
        self.value = 2

    def __repr__(self) -> str:
        return f"<Paper(value={self.value})>"


class Scissors:
    def __init__(self) -> None:
        self.value = 3

    def __repr__(self) -> str:
        return f"<Scissors(value={self.value})>"


class RoundScore:
    def __init__(self, player_1_points: int, player_2_points: int) -> None:
        self.player_1_points: int = player_1_points
        self.player_2_points: int = player_2_points

    def __repr__(self) -> str:
        return (
            f"<RoundScore(player_1_points={self.player_1_points}, "
            f"player_2_points={self.player_2_points})>"
        )


class RockPaperScissorsRound:
    def __init__(
        self,
        player_1_choice: Rock | Paper | Scissors,
        player_2_choice: Rock | Paper | Scissors,
    ) -> None:
        self.player_1_choice: Rock | Paper | Scissors = player_1_choice
        self.player_2_choice: Rock | Paper | Scissors = player_2_choice
        self.result: RoundScore = self.play_round()

    def __repr__(self) -> str:
        return (
            f"<RockPaperScissorsRound(player_1_choice={self.player_1_choice}, "
            f"player_2_choice={self.player_2_choice}, "
            f"result={self.result})>"
        )

    def play_round(self) -> RoundScore:
        if type(self.player_1_choice) == type(self.player_2_choice):
            return RoundScore(
                player_1_points=(3 + self.player_1_choice.value),
                player_2_points=(3 + self.player_2_choice.value),
            )
        elif isinstance(self.player_1_choice, Rock):
            if isinstance(self.player_2_choice, Paper):
                return RoundScore(
                    player_1_points=(0 + self.player_1_choice.value),
                    player_2_points=(6 + self.player_2_choice.value),
                )
            else:
                return RoundScore(
                    player_1_points=(6 + self.player_1_choice.value),
                    player_2_points=(0 + self.player_2_choice.value),
                )
        elif isinstance(self.player_1_choice, Paper):
            if isinstance(self.player_2_choice, Scissors):
                return RoundScore(
                    player_1_points=(0 + self.player_1_choice.value),
                    player_2_points=(6 + self.player_2_choice.value),
                )
            else:
                return RoundScore(
                    player_1_points=(6 + self.player_1_choice.value),
                    player_2_points=(0 + self.player_2_choice.value),
                )
        elif isinstance(self.player_1_choice, Scissors):
            if isinstance(self.player_2_choice, Rock):
                return RoundScore(
                    player_1_points=(0 + self.player_1_choice.value),
                    player_2_points=(6 + self.player_2_choice.value),
                )
            else:
                return RoundScore(
                    player_1_points=(6 + self.player_1_choice.value),
                    player_2_points=(0 + self.player_2_choice.value),
                )
        return RoundScore(0, 0)


@dataclass
class PlayerChoices:
    player_1_choice: Rock | Paper | Scissors
    player_2_choice: Rock | Paper | Scissors

    def __repr__(self) -> str:
        return f"<PlayerChoices(player_1_choice={self.player_1_choice}, player_2_choice={self.player_2_choice})>"


def subsitute_value(value: str) -> Rock | Paper | Scissors:
    if value in ("A", "X"):
        return Rock()
    elif value in ("B", "Y"):
        return Paper()
    else:
        return Scissors()


def get_player_choices(line: str) -> PlayerChoices:
    game_values = line.strip().split(" ")
    return PlayerChoices(
        player_1_choice=subsitute_value(game_values[0]),
        player_2_choice=subsitute_value(game_values[1]),
    )


def main():
    input_file_path = Path("../input.txt")
    with open(input_file_path, "r") as f:
        input_lines: list[str] = f.readlines()

    player_2_sum = 0
    for line in input_lines:
        if line != "\n":
            choices = get_player_choices(line)
            round = RockPaperScissorsRound(
                choices.player_1_choice, choices.player_2_choice
            )
            player_2_sum += round.result.player_2_points

    print(f"Total score after all rounds for player 2: {player_2_sum}")


if __name__ == "__main__":
    main()

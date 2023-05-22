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
        return (
            f"<PlayerChoices(player_1_choice={self.player_1_choice}"
            f"player_2_choice={self.player_2_choice})>"
        )


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


def enforce_strategy_guide(
    opponents_choice: Rock | Paper | Scissors, wanted_outcome: str
) -> Rock | Paper | Scissors:
    """
    Accepts the rock, paper, scissors choice of an opponent, along with the desired
    outcome of the round in `wanted_outcome`.

    `X` should cause an intentional loss for, `Y` should force a draw and `Z` should
    cause a win.
    """
    if wanted_outcome == "X":
        # Intentionally lose
        if isinstance(opponents_choice, Rock):
            return Scissors()
        elif isinstance(opponents_choice, Paper):
            return Rock()
        else:
            return Paper()
    elif wanted_outcome == "Y":
        # Force a draw
        return opponents_choice
    else:
        # Win
        if isinstance(opponents_choice, Rock):
            return Paper()
        elif isinstance(opponents_choice, Paper):
            return Scissors()
        else:
            return Rock()


def main():
    input_file_path = Path("../input.txt")
    with open(input_file_path, "r") as f:
        input_lines: list[str] = f.readlines()

    part_1_sum = 0
    for line in input_lines:
        if line != "\n":
            choices = get_player_choices(line)
            round = RockPaperScissorsRound(
                choices.player_1_choice, choices.player_2_choice
            )
            part_1_sum += round.result.player_2_points

    print(f"Part 1: {part_1_sum}")

    part_2_sum = 0
    for line in input_lines:
        if line != "\n":
            choices = get_player_choices(line)
            choices.player_2_choice = enforce_strategy_guide(
                choices.player_1_choice, wanted_outcome=line.strip().split()[1]
            )
            round = RockPaperScissorsRound(
                choices.player_1_choice, choices.player_2_choice
            )
            part_2_sum += round.result.player_2_points

    print(f"Part 2: {part_2_sum}")

if __name__ == "__main__":
    main()

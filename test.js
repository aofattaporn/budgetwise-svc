const rock = "rock";
const paper = "paper";
const scissors = "scissors";
const draw = "draw";

function play(arg1, arg2) {
  const validArgs = [rock, paper, scissors];

  if (!validArgs.includes(arg1) || !validArgs.includes(arg2)) {
    return undefined;
  }

  if (arg1 === arg2) {
    return draw;
  }

  if (arg1 === rock) {
    if (arg2 === paper) {
      return paper;
    }
    if (arg2 === scissors) {
      return rock;
    }
  }

  if (arg1 === paper) {
    if (arg2 === rock) {
      return paper;
    }
    if (arg2 === scissors) {
      return scissors;
    }
  }

  if (arg1 === scissors) {
    if (arg2 === rock) {
      return rock;
    }
    if (arg2 === paper) {
      return scissors;
    }
  }
}

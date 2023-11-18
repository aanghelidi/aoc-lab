const file = Bun.file("input.txt");
const input: string = await file.text();

type Position = {
  x: number;
  y: number;
};

// Start at 0,0 facing north
let position: Position = { x: 0, y: 0 };
let direction: string = "N";

const directionsTurn: Map<string, Map<string, string>> = new Map([
  [
    "N",
    new Map([
      ["L", "W"],
      ["R", "E"],
    ]),
  ],
  [
    "E",
    new Map([
      ["L", "N"],
      ["R", "S"],
    ]),
  ],
  [
    "S",
    new Map([
      ["L", "E"],
      ["R", "W"],
    ]),
  ],
  [
    "W",
    new Map([
      ["L", "S"],
      ["R", "N"],
    ]),
  ],
]);
const visited: Set<string> = new Set();
let visitTwice = 0;
visited.add(JSON.stringify(position));

function manhattanDistance(start: Position): number {
  return Math.abs(start.x) + Math.abs(start.y);
}

input
  .trim()
  .split(", ")
  .map((instruction) => {
    const turn = instruction[0];
    const steps = Number(instruction.slice(1));
    const newDirection = directionsTurn.get(direction);
    if (!newDirection) {
      throw new Error("Invalid direction");
    }
    const turnDirection = newDirection.get(turn);
    if (!turnDirection) {
      throw new Error("Invalid turn");
    }
    direction = turnDirection;
    switch (direction) {
      case "N":
        for (let i = 0; i < steps; i++) {
          position.y += 1;
          if (visited.has(JSON.stringify(position)) && visitTwice === 0) {
            console.log(`Part 2: ${manhattanDistance(position)}`);
            visitTwice += 1;
          }
          visited.add(JSON.stringify(position));
        }
        break;
      case "E":
        for (let i = 0; i < steps; i++) {
          position.x += 1;
          if (visited.has(JSON.stringify(position)) && visitTwice === 0) {
            console.log(`Part 2: ${manhattanDistance(position)}`);
            visitTwice += 1;
          }
          visited.add(JSON.stringify(position));
        }
        break;
      case "S":
        for (let i = 0; i < steps; i++) {
          position.y -= 1;
          if (visited.has(JSON.stringify(position)) && visitTwice === 0) {
            console.log(`Part 2: ${manhattanDistance(position)}`);
            visitTwice += 1;
          }
          visited.add(JSON.stringify(position));
        }
        break;
      case "W":
        for (let i = 0; i < steps; i++) {
          position.x -= 1;
          if (visited.has(JSON.stringify(position)) && visitTwice === 0) {
            console.log(`Part 2: ${manhattanDistance(position)}`);
            visitTwice += 1;
          }
          visited.add(JSON.stringify(position));
        }
        break;
    }
  });

console.log(`Part 1: ${manhattanDistance(position)}`);

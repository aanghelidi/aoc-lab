const file = Bun.file("input.txt");
const raw = await file.text();
const lines = raw.trim().split("\n");

type Grid = number[][];
const grid: Grid = Array.from({ length: 1000 }, () => Array(1000).fill(0));

type PuzzlePart = "1" | "2";
interface Action {
  readonly name: string;

  start: number[];
  end: number[];
  execute: (grid: Grid, part: PuzzlePart) => void;
}

class TurnOffAction implements Action {
  readonly name: string = "turn off";

  start: number[];
  end: number[];

  constructor(start: number[], end: number[]) {
    this.start = start;
    this.end = end;
  }

  execute(grid: Grid, part: PuzzlePart): void {
    const [x1, y1] = this.start;
    const [x2, y2] = this.end;

    switch (part) {
      case "1":
        for (let x = x1; x <= x2; x++) {
          for (let y = y1; y <= y2; y++) {
            grid[x][y] = 0;
          }
        }
        break;

      case "2":
        for (let x = x1; x <= x2; x++) {
          for (let y = y1; y <= y2; y++) {
            grid[x][y] = Math.max(0, grid[x][y] - 1);
          }
        }
        break;
    }
  }
}

class TurnOnAction implements Action {
  readonly name: string = "turn on";

  start: number[];
  end: number[];

  constructor(start: number[], end: number[]) {
    this.start = start;
    this.end = end;
  }

  execute(grid: Grid, part: PuzzlePart): void {
    const [x1, y1] = this.start;
    const [x2, y2] = this.end;

    switch (part) {
      case "1":
        for (let x = x1; x <= x2; x++) {
          for (let y = y1; y <= y2; y++) {
            grid[x][y] = 1;
          }
        }
        break;
      case "2":
        for (let x = x1; x <= x2; x++) {
          for (let y = y1; y <= y2; y++) {
            grid[x][y] += 1;
          }
        }
    }
  }
}

class ToggleAction implements Action {
  readonly name: string = "toggle";

  start: number[];
  end: number[];

  constructor(start: number[], end: number[]) {
    this.start = start;
    this.end = end;
  }

  execute(grid: Grid, part: PuzzlePart): void {
    const [x1, y1] = this.start;
    const [x2, y2] = this.end;

    switch (part) {
      case "1":
        for (let x = x1; x <= x2; x++) {
          for (let y = y1; y <= y2; y++) {
            grid[x][y] === 0 ? (grid[x][y] = 1) : (grid[x][y] = 0);
          }
        }
        break;
      case "2":
        for (let x = x1; x <= x2; x++) {
          for (let y = y1; y <= y2; y++) {
            grid[x][y] += 2;
          }
        }
        break;
    }
  }
}

const parseTurn = (line: string): Action => {
  const parts = line.split(" ");
  const start = parts[2].split(",").map(Number);
  const end = parts[4].split(",").map(Number);
  return parts[1] === "on"
    ? new TurnOnAction(start, end)
    : new TurnOffAction(start, end);
};

const parseToggle = (line: string): Action => {
  const parts = line.split(" ");
  const start = parts[1].split(",").map(Number);
  const end = parts[3].split(",").map(Number);
  return new ToggleAction(start, end);
};

const parseAction = (line: string): Action => {
  return line.startsWith("turn") ? parseTurn(line) : parseToggle(line);
};

const countLit = (grid: Grid): number => {
  return grid.reduce((acc: number, row: number[]) => {
    return (
      acc +
      row.reduce((acc: number, c: number) => {
        if (c) acc += c;
        return acc;
      }, 0)
    );
  }, 0);
};

lines.forEach((l: string) => {
  const action = parseAction(l);
  //action.execute(grid, "1");
  action.execute(grid, "2");
});

console.log(`Answer: ${countLit(grid)}`);

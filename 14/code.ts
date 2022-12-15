import * as fs from "fs";
import { range_ } from "@arrows/array";

type Point = [number, number];

const START: Point = [500, 0];

const parseInput = (input: string) =>
  input
    .split("\n")
    .map(line => line.split(" -> ").map(item => item.split(",").map(Number)));

const containers = (input: number[][][]) => {
  const blocked = new Set<string>();
  let maxY = 0;

  for (const containerData of input) {
    for (let i = 0; i < containerData.length - 1; i++) {
      const [fromX, fromY] = containerData[i];
      const [toX, toY] = containerData[i + 1];

      maxY = Math.max(fromY, toY, maxY);

      for (const y of range_(fromY, toY + (fromY > toY ? -1 : 1))) {
        for (const x of range_(fromX, toX + (fromX > toX ? -1 : 1))) {
          blocked.add(`${x}:${y}`);
        }
      }
    }
  }

  return { blocked, maxY };
}

const simulateSandUnit = (
  blocked: Set<string>,
  stopWhen?: (y: number) => boolean
): Point | null => {
  let [x, y] = START;

  while (true) {
    if (stopWhen?.(y)) {
      return null;
    }

    if (!blocked.has(`${x}:${y+1}`)) {
      y++;
    } else if (!blocked.has(`${x-1}:${y+1}`)) {
      x--;
      y++;
    } else if (!blocked.has(`${x+1}:${y+1}`)) {
      x++;
      y++;
    } else {
      return [x, y];
    }
  }
}

const part1 = (contents: string) => {
  const input = parseInput(contents);
  const { blocked, maxY } = containers(input);
  const containersCount = blocked.size;

  const stopWhen = (y: number) => y === maxY;

  while (true) {
    const result = simulateSandUnit(blocked, stopWhen);

    if (result === null) {
      break;
    }

    const [x, y] = result;
    blocked.add(`${x}:${y}`);
  }

  return blocked.size - containersCount;
}

const part2 = (contents: string) => {
  const input = parseInput(contents);
  const { blocked, maxY } = containers(input);
  const floorY = maxY + 2;
  const minReqFloorX = range_(START[0] - floorY, START[0] + floorY + 1);

  for (const x of minReqFloorX) {
    blocked.add(`${x}:${floorY}`);
  }

  const containersCount = blocked.size;

  while (true) {
    const [x, y] = simulateSandUnit(blocked) as Point;
    blocked.add(`${x}:${y}`);

    if (y === 0) {
      break;
    }
  }

  return blocked.size - containersCount;
}

["./test.txt", "./input.txt"].forEach((filename) => {
  console.log(filename);
  const input = fs.readFileSync(filename, "utf-8");
  console.log(part1(input));
  console.log(part2(input));
})


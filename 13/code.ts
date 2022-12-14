import * as fs from "fs";
import * as util from "util";

type Data = number | Data[];
type Result = "equal" | "right" | "wrong"

const a: Data = [[1], [2, 3, 4]];
const b: Data = [1, [2, [3, [4, [5, 6, 7]]]], 8, 9];

interface Pair {
  first: Data;
  second: Data;
}

const parsePairs = (input: string[]): Pair[] => {
  const result: Pair[] = [];
  for (let i = 0; i < input.length; i += 3) {
    const [firstStr, secondStr] = input.slice(i, i+2);
    const first = JSON.parse(firstStr) as Data;
    const second = JSON.parse(secondStr) as Data;
    result.push({first, second});
  }
  return result;
}

const comparePairs = (first: Data | undefined, second: Data | undefined): Result => {
  // const spaces = Array(level).join(" ");
  // console.log(spaces, `- Compare ${util.inspect(first)} vs ${util.inspect(second)}`);
  if (typeof first === "number" && typeof second === "number") {
    if (first === second) {
      return "equal";
    } else if (first < second) {
      // console.log(spaces, "- Left side is smaller, so inputs are in the right order");
      return "right";
    } else {
      // console.log(spaces, "- Right side is smaller, so inputs are not in the right order")
      return "wrong";
    }
  }
  if (typeof first === "undefined" && (typeof second === "number" || Array.isArray(second))) {
    // console.log(spaces, "- Left side ran out of items, so inputs are in the right order");
    return "right";
  }
  if ((typeof first === "number" || Array.isArray(first)) && typeof second === "undefined") {
    // console.log(spaces, "- Right side ran out of items, so inputs are not in the right order");
    return "wrong";
  }
  if ((Array.isArray(first) && Array.isArray(second)) ||
      (Array.isArray(first) && typeof second === "number") ||
      (typeof first === "number" && Array.isArray(second))) {
    if (typeof first === "number") {
      // console.log(spaces, `- Mixed types; convert left to [${first}] and retry`);
    } else if (typeof second === "number") {
      // console.log(spaces, `- Mixed types; convert right to [${second}] and retry`);
    }
    first = typeof first === "number" ? [first] : first;
    second = typeof second === "number" ? [second] : second;
    for (let i = 0; i < Math.max(first.length, second.length); i++) {
      const result = comparePairs(first[i], second[i], level + 1);
      if (result === "equal") {
        continue
      } else {
        return result;
      }
    }
  }
  return "equal";
}

["./test.txt", "./input.txt"].forEach((filename) => {
  const contents = fs.readFileSync(filename, "utf-8").split("\n");
  const pairs = parsePairs(contents);
  console.log("Part1");
  const result = pairs
    .map((pair, index) => ({result: comparePairs(pair.first, pair.second), index: index + 1}))
    .filter((intermediate) => intermediate.result === "right")
    .map(({index}) => index)
    .reduce((acc, cur) => acc + cur, 0);
  console.log(result);
  console.log("Part2");
  const allPackets: Data[] = contents
    .filter(line => line.length > 0)
    .map(line => JSON.parse(line) as Data);
  const dividerPackets: Data[] = [[[2]], [[6]]];
  const indices = dividerPackets.map((dividerPacket, index) => {
    // [[6]] will come after [[2]], add 1 to its count
    let count = index + 1;
    allPackets.forEach(packet => {
      const order = comparePairs(dividerPacket, packet);
      if (order === "wrong") {
        count++;
      }
    });
    return count;
  });
  console.log(indices.reduce((acc, cur) => acc * cur, 1));
})


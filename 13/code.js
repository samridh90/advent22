"use strict";
exports.__esModule = true;
var fs = require("fs");
var a = [[1], [2, 3, 4]];
var b = [1, [2, [3, [4, [5, 6, 7]]]], 8, 9];
var parsePairs = function (input) {
    var result = [];
    for (var i = 0; i < input.length; i += 3) {
        var _a = input.slice(i, i + 2), firstStr = _a[0], secondStr = _a[1];
        var first = JSON.parse(firstStr);
        var second = JSON.parse(secondStr);
        result.push({ first: first, second: second });
    }
    return result;
};
var comparePairs = function (first, second, level) {
    // const spaces = Array(level).join(" ");
    // console.log(spaces, `- Compare ${util.inspect(first)} vs ${util.inspect(second)}`);
    if (typeof first === "number" && typeof second === "number") {
        if (first === second) {
            return "equal";
        }
        else if (first < second) {
            // console.log(spaces, "- Left side is smaller, so inputs are in the right order");
            return "right";
        }
        else {
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
        }
        else if (typeof second === "number") {
            // console.log(spaces, `- Mixed types; convert right to [${second}] and retry`);
        }
        first = typeof first === "number" ? [first] : first;
        second = typeof second === "number" ? [second] : second;
        for (var i = 0; i < Math.max(first.length, second.length); i++) {
            var result = comparePairs(first[i], second[i], level + 1);
            if (result === "equal") {
                continue;
            }
            else {
                return result;
            }
        }
    }
    return "equal";
};
["./test.txt", "./input.txt"].forEach(function (filename) {
    var contents = fs.readFileSync(filename, "utf-8").split("\n");
    var pairs = parsePairs(contents);
    console.log("Part1");
    var result = pairs
        .map(function (pair, index) { return ({ result: comparePairs(pair.first, pair.second, 0), index: index + 1 }); })
        .filter(function (intermediate) { return intermediate.result === "right"; })
        .map(function (_a) {
        var index = _a.index;
        return index;
    })
        .reduce(function (acc, cur) { return acc + cur; }, 0);
    console.log(result);
    console.log("Part2");
    var allPackets = contents
        .filter(function (line) { return line.length > 0; })
        .map(function (line) { return JSON.parse(line); });
    var dividerPackets = [[[2]], [[6]]];
    var indices = dividerPackets.map(function (dividerPacket, index) {
        // [[6]] will come after [[2]], add 1 to its count
        var count = index + 1;
        allPackets.forEach(function (packet) {
            var order = comparePairs(dividerPacket, packet, 0);
            if (order === "wrong") {
                count++;
            }
        });
        return count;
    });
    console.log(indices.reduce(function (acc, cur) { return acc * cur; }, 1));
});

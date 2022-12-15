"use strict";
exports.__esModule = true;
var fs = require("fs");
var array_1 = require("@arrows/array");
var START = [500, 0];
var parseInput = function (input) {
    return input
        .split("\n")
        .map(function (line) { return line.split(" -> ").map(function (item) { return item.split(",").map(Number); }); });
};
var containers = function (input) {
    var blocked = new Set();
    var maxY = 0;
    for (var _i = 0, input_1 = input; _i < input_1.length; _i++) {
        var containerData = input_1[_i];
        for (var i = 0; i < containerData.length - 1; i++) {
            var _a = containerData[i], fromX = _a[0], fromY = _a[1];
            var _b = containerData[i + 1], toX = _b[0], toY = _b[1];
            maxY = Math.max(fromY, toY, maxY);
            for (var _c = 0, _d = (0, array_1.range_)(fromY, toY + (fromY > toY ? -1 : 1)); _c < _d.length; _c++) {
                var y = _d[_c];
                for (var _e = 0, _f = (0, array_1.range_)(fromX, toX + (fromX > toX ? -1 : 1)); _e < _f.length; _e++) {
                    var x = _f[_e];
                    blocked.add("".concat(x, ":").concat(y));
                }
            }
        }
    }
    return { blocked: blocked, maxY: maxY };
};
var simulateSandUnit = function (blocked, stopWhen) {
    var x = START[0], y = START[1];
    while (true) {
        if (stopWhen === null || stopWhen === void 0 ? void 0 : stopWhen(y)) {
            return null;
        }
        if (!blocked.has("".concat(x, ":").concat(y + 1))) {
            y++;
        }
        else if (!blocked.has("".concat(x - 1, ":").concat(y + 1))) {
            x--;
            y++;
        }
        else if (!blocked.has("".concat(x + 1, ":").concat(y + 1))) {
            x++;
            y++;
        }
        else {
            return [x, y];
        }
    }
};
var part1 = function (contents) {
    var input = parseInput(contents);
    var _a = containers(input), blocked = _a.blocked, maxY = _a.maxY;
    var containersCount = blocked.size;
    var stopWhen = function (y) { return y === maxY; };
    while (true) {
        var result = simulateSandUnit(blocked, stopWhen);
        if (result === null) {
            break;
        }
        var x = result[0], y = result[1];
        blocked.add("".concat(x, ":").concat(y));
    }
    return blocked.size - containersCount;
};
var part2 = function (contents) {
    var input = parseInput(contents);
    var _a = containers(input), blocked = _a.blocked, maxY = _a.maxY;
    var floorY = maxY + 2;
    var minReqFloorX = (0, array_1.range_)(START[0] - floorY, START[0] + floorY + 1);
    for (var _i = 0, minReqFloorX_1 = minReqFloorX; _i < minReqFloorX_1.length; _i++) {
        var x = minReqFloorX_1[_i];
        blocked.add("".concat(x, ":").concat(floorY));
    }
    var containersCount = blocked.size;
    while (true) {
        var _b = simulateSandUnit(blocked), x = _b[0], y = _b[1];
        blocked.add("".concat(x, ":").concat(y));
        if (y === 0) {
            break;
        }
    }
    return blocked.size - containersCount;
};
["./test.txt", "./input.txt"].forEach(function (filename) {
    console.log(filename);
    var input = fs.readFileSync(filename, "utf-8");
    console.log(part1(input));
    console.log(part2(input));
});

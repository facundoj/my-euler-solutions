var fs = require('fs'),
    input;

var data = fs.readFileSync('./triangle.txt');
input = data.toString('utf8').trim().split('\n').map(function(row) {
    return row.trim().split(' ').map(function(num) {
        return +num;
    });
});

var references = [],
    sumMax = 0;

input.forEach(function(list, i) {
    var max = 0;
    references[i] = [];

    list.forEach(function(num, j) {
        max = Math.max(max, num);
    });

    sumMax += max;

    list.forEach(function(num, j) {
        references[i][j] = num - max;
    });
});

var costsTriangle = [];
references.forEach(function(list, i) {
    costsTriangle[i] = [];
    list.forEach(function(num, j) {
        var left, right;

        if (i == 0) {
            costsTriangle[i][j] = 0;
            return;
        }

        if (j == 0) {
            right = references[i][j] + costsTriangle[i - 1][j];
            left = null;
        } else if (j + 1 == list.length) {
            right = null;
            left = references[i][j] + costsTriangle[i - 1][j - 1];
        } else {
            right = references[i][j] + costsTriangle[i - 1][j];
            left = references[i][j] + costsTriangle[i - 1][j - 1];
        }

        costsTriangle[i][j] = Math.max(left === null ? -Infinity : left, right === null ? -Infinity : right);
    });
});

var aux = -Infinity;
costsTriangle[costsTriangle.length - 1].forEach(function(num, i) {
    if (num > aux) {
        aux = num;
    }
});

console.log(sumMax + aux);

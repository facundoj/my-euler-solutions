// Problem 17
// If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
//
// If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?
//
//
// NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.

'use strict';
var numbers = {
        1: 'one',
        2: 'two',
        3: 'three',
        4: 'four',
        5: 'five',
        6: 'six',
        7: 'seven',
        8: 'eight',
        9: 'nine',
        10: 'ten',
        11: 'eleven',
        12: 'twelve',
        13: 'thirteen',
        14: 'fourteen',
        15: 'fifteen',
        16: 'sixteen',
        17: 'seventeen',
        18: 'eighteen',
        teens: '%un%teen',
        20: 'twenty',
        30: 'thirty',
        40: 'forty',
        50: 'fifty',
        60: 'sixty',
        70: 'seventy',
        80: 'eighty',
        90: 'ninety',
        twoDecs: '%dec%%un%',
        roundHundred: '%hun%hundred',
        hundreds: '%hun%hundredand%twoDecs%',
        1000: 'onethousand'
    },
    l = 0;

for (let i = 1; i <= 1000; i++ ) {
    let word,
        twoDecs = i % 100, twoDecsWord;

    // Generating word for the last two nums
    if (twoDecs > 0 && twoDecs <= 18) {
        twoDecsWord = numbers[twoDecs];
    } else if (twoDecs > 18 && twoDecs < 20) {
        twoDecsWord = numbers.teens.replace('%un%', numbers[twoDecs % 10]);
    } else if (twoDecs % 10 === 0) {
        twoDecsWord = numbers[twoDecs];
    } else if (twoDecs > 0) {
        twoDecsWord = numbers.twoDecs
            .replace('%dec%', numbers[twoDecs - (twoDecs % 10)])
            .replace('%un%', numbers[twoDecs % 10]);
    }

    if (i < 1000 && twoDecs === 0) {
        word = numbers.roundHundred
            .replace('%hun%', numbers[i / 100]);
    } else if (i < 1000 && i > 100) {
        word = numbers.hundreds
            .replace('%hun%', numbers[(i - twoDecs) / 100])
            .replace('%twoDecs%', twoDecsWord);
    } else if (i < 100) {
        word = twoDecsWord;
    } else {
        word = numbers[i];
    }

    l += word.length;
    console.log(i, word, word.length);
}

console.log('Solution', l);

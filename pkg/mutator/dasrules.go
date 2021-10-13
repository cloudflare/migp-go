// Copyright (c) 2021 Cloudflare, Inc. All rights reserved.
// SPDX-License-Identifier: BSD-3-Clause

package mutator

// dasRulesJsonString is just a cut and paste of the ordered Das mangling rules
const dasRulesJSONString = `[
    {
        "ruletype": "c",
        "string1": "",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "d",
        "string1": "",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "d",
        "string1": "",
        "position": -2,
        "string2": ""
    },
    {
        "ruletype": "d",
        "string1": "",
        "position": -3,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "0",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "1",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "a",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "q",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "0",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "d",
        "string1": "",
        "position": 1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "a",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "5",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "123",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "d",
        "string1": "",
        "position": 2,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "2",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "7",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "z",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "9",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "d",
        "string1": "",
        "position": 3,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "00",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "6",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "3",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "8",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "11",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "4",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "w",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "b",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "q",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "00",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "99",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "12",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "s",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "98",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "1",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "s",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "78",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "l",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "p",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "89",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "000",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "90",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "z",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "10",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "23",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "01",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "m",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "r",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "x",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "000",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "k",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "e",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "22",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "f",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "n",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "88",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "w",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "19",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "c",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "13",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "m",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "p",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "456",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "21",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "y",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "777",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "d",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "b",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "10",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "s",
        "string1": "0",
        "position": 0,
        "string2": "o"
    },
    {
        "ruletype": "i",
        "string1": "08",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "07",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "g",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "007",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "d",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "y",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "h",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "77",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "t",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "k",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "55",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "f",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "09",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "c",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "A",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "56",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "j",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "v",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "45",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "67",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "i",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "s",
        "string1": "o",
        "position": 0,
        "string2": "0"
    },
    {
        "ruletype": "i",
        "string1": "001",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "A",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "t",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "101",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "l",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "01",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "92",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "r",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "g",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "02",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "n",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "91",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "o",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "3",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "h",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "321",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "95",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "123",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "33",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "x",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "85",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "17",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "06",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "66",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "j",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "19",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "05",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "03",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "666",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "24",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "69",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "u",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "16",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "80",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "Q",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "87",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "79",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "86",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "93",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "96",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "e",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "65",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "15",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "25",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "7",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "44",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "4",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "97",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "05",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "09",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "8",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "2",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "83",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "57",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "50",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "94",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "54",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "11",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "08",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "61",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "20",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "14",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "v",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "81",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "74",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "5",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "76",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "34",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "18",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "82",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "41",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "04",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "K",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "06",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "S",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "V",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "007",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "75",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "02",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "9",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "28",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "53",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "s",
        "string1": "asdf",
        "position": 0,
        "string2": "1234"
    },
    {
        "ruletype": "s",
        "string1": "1234",
        "position": 0,
        "string2": "asdf"
    },
    {
        "ruletype": "i",
        "string1": "N",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "84",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "20",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "27",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "Q",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "13",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "D",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "J",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "26",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "52",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "29",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "30",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "49",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "58",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "07",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "38",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "Z",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "31",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "o",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "72",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "63",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "420",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "s",
        "string1": "3",
        "position": 0,
        "string2": "e"
    },
    {
        "ruletype": "i",
        "string1": "P",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "04",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "32",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "85",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "R",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "47",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "16",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "i",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "71",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "51",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "12",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "22",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "87",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "man",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "73",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "46",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "35",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "L",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "59",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "98",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "68",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "32",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "60",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "K",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "42",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "27",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "36",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "H",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "64",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "G",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "s",
        "string1": "e",
        "position": 0,
        "string2": "3"
    },
    {
        "ruletype": "i",
        "string1": "101",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "u",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "Z",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "B",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "39",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "boy",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "P",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "E",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "79",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "28",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "O",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "30",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "L",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "69",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "D",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "31",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "52",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "456",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "s",
        "string1": "qwer",
        "position": 0,
        "string2": "1234"
    },
    {
        "ruletype": "s",
        "string1": "5678",
        "position": 0,
        "string2": "qwer"
    },
    {
        "ruletype": "i",
        "string1": "70",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "s",
        "string1": "qwer",
        "position": 0,
        "string2": "5678"
    },
    {
        "ruletype": "i",
        "string1": "41",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "S",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "F",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "C",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "15",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "s",
        "string1": "qwe",
        "position": 0,
        "string2": "qaz"
    },
    {
        "ruletype": "s",
        "string1": "qaz",
        "position": 0,
        "string2": "qwe"
    },
    {
        "ruletype": "i",
        "string1": "18",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "36",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "66",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "53",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "W",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "89",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "48",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "I",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "17",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "40",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "O",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "82",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "77",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "X",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "94",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "45",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "51",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "50",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "ita",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "W",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "65",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "14",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "93",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "21",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "s",
        "string1": "1234",
        "position": 0,
        "string2": "qwer"
    },
    {
        "ruletype": "i",
        "string1": "R",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "T",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "25",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "143",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "s",
        "string1": "5678",
        "position": 0,
        "string2": "1234"
    },
    {
        "ruletype": "i",
        "string1": "60",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "V",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "man",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "38",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "86",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "46",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "92",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "G",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "54",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "M",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "70",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "089",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "68",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "62",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "78",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "Y",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "62",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "40",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "C",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "H",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "B",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "X",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "26",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "T",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "s",
        "string1": "wsx",
        "position": 0,
        "string2": "2wsx"
    },
    {
        "ruletype": "s",
        "string1": "2wsx",
        "position": 0,
        "string2": "wsx"
    },
    {
        "ruletype": "i",
        "string1": "37",
        "position": 0,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "03",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "J",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "55",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "49",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "i",
        "string1": "34",
        "position": -1,
        "string2": ""
    },
    {
        "ruletype": "s",
        "string1": "zxcv",
        "position": 0,
        "string2": "asdf"
    },
    {
        "ruletype": "s",
        "string1": "asdf",
        "position": 0,
        "string2": "zxcv"
    },
    {
        "ruletype": "i",
        "string1": "N",
        "position": -1,
        "string2": ""
    }
]`

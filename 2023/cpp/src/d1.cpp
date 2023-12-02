//
// Created by jacob on 12/1/2023.
//

#include "d1.h"
#include <string>
#include <iostream>
#include <vector>
#include <limits.h>
#include <ctype.h>
#include <cstdlib>
#include <fstream>

void aoc::d1::day_one_solution() {

    // Test cases
    std::string t1{"1abc2"};
    std::string t2{"pqr3stu8vwx"};
    std::string t3{"a1b2c3d4e5f"};
    std::string t4{"treb7uchet"};
    std::vector<std::string> inputs{t1, t2, t3, t4};

    int sum{0};
    std::string file{"../inputs/d1.txt"};
    std::ifstream fs;
    fs.open(file);

    if (!fs.is_open()) {
        std::cout << "fstream failed to parse input for d1\n";
        fs.clear();
        fs.close();
        return;
    }

    std::string line; // line storage buffer
    while (std::getline(fs, line)) {
        sum += process_line(line);
    }
    fs.close();

    // Print soln and clear buffer
    std::cout << "D1P1 Soln: " << sum << std::endl;
}


int aoc::d1::process_line(const std::string &line_from_file) {
    int first{INT_MIN};
    int second{INT_MIN};
    bool first_flip{false};

    for (char c: line_from_file) {
        if (isdigit(c)) {
            if (!first_flip) {
                first = c - '0';
                first_flip = true;
            } else {
                second = c - '0';
            }
        }
    }
    // second value was never encountered... store the same value first encountered
    if (second == INT_MIN) {
        second = first;
    }
    std::string combined{std::to_string(first)};
    combined += std::to_string(second);
    int f = std::stoi(combined);

    return f;
}
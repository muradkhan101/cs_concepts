#include <stdio.h>
#include <string.h>

void copyChars(char* copyFrom, char* copyTo, int start, int end) {
    int copyToIndex = 0;
    for (int i = start; i < end; i++) {
        copyTo[copyToIndex] = copyFrom[i];
        copyToIndex++;
    }
}

char** split(char* string, char delimiter) {
    int timesToSplit = 0;
    for (int i = 0; string[i] != 0; i++) {
        if (string[i] == delimiter) {
            timesToSplit++;
        }
    }
    if (timesToSplit == 0) {
        timesToSplit = 1;
    }
    char ** finalArray = malloc(sizeof(char*) * timesToSplit);
    int lastMatchAt = 0;
    int timesSplit = 0;
    for (int i = 0; string[i] != 0; i++) {
        if (string[i] == delimiter) {
            char* part = malloc(sizeof(char) * (i - lastMatchAt));
            copyChars(string, part, lastMatchAt, i);
            finalArray[timesSplit] = part;
            lastMatchAt = i + 1;
            timesSplit++;
        }
        if (timesSplit == 0 && string[i + 1] == 0) {
            char *part = malloc(sizeof(char) * (i - lastMatchAt));
            copyChars(string, part, lastMatchAt, i + 1);
            finalArray[timesSplit] = part;
        }
    }
    return finalArray;
}

int main() {
    printf("SPLITTING FIRST ONE\n");
    char** splitOne = split("this.is.the.first.test\0", '.');
    printf(splitOne[0]);
    printf(splitOne[3]);
    char **splitTwo = split("secondtest\0", ".");
    printf(splitTwo[0]);
}
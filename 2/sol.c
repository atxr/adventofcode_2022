#!/usr/bin/tcc -run

#include <stdio.h>
#define BUF_SIZE 5

unsigned int get_score1(char* line) {
  char opponent = line[0] - 'A';
  char you = line[2] - 'X';
  return you + 1 + 3 * ((you - opponent + 4) % 3);
}

unsigned int get_score2(char* line) {
  char opponent = line[0] - 'A';
  char res = line[2] - 'X';
  char you = (opponent + res + 2) % 3;
  return res * 3 + you + 1;
}

int main (int argc, char *argv[])
{
  FILE* input;
  char buf[BUF_SIZE];
  unsigned int score1;
  unsigned int score2;

  if (argc != 2) {
    perror("One arg required\n");
    return -1;
  }

  if (!(input  = fopen(argv[1], "r"))) {
    perror("Cannot open file.");
    return -1;
  }

  while (fgets(buf, BUF_SIZE, input) != NULL) {
    score1 += get_score1(buf);
    score2 += get_score2(buf);
  }

  printf("With this input, you score with method 1 is %u\n", score1);
  printf("With this input, you score with method 2 is %u\n", score2);
  fclose(input);
  return 0;
}

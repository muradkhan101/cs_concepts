#include <stdio.h>
#include <fcntl.h> // open, O_RDONLY
#include <stdlib.h> // exit
#include <unistd.h> // read
#include <string.h>

int read_line(int fd, char **line_ptr) {
  char next_byte;
  int line_length = 0;
  // int some_number = read(fd, &next_byte, 1);
  // printf("The value of the number is %d\n", some_number);
  while (read(fd, &next_byte, 1) > 0) {
    if (next_byte != 10) {
      line_length++;
    } else {
      line_length++;
      lseek(fd, (line_length) * -1, SEEK_CUR);
      *line_ptr = malloc(line_length);
      read(fd, *line_ptr, line_length);
      return line_length;
    }
  }
  if (line_length > 0) {
    lseek(fd, (line_length) * -1, SEEK_CUR);
    *line_ptr = malloc(line_length);
    read(fd, *line_ptr, line_length);
    return line_length;
  } else {
    // At end of file already and no new line
    return -1;
  }
}

void print_line(char *line_ptr, int length) {
  for (int i = 0; i < length; i++) {
    printf("%c", line_ptr[i]);
  }
}

int main() {
  int file_d;
  char *line;

  file_d = open("text.txt", O_RDONLY);
  if (file_d == -1) {
    printf("[ERROR] Couldnt open file\n");
    exit(-1);
  }
  printf("[DEBUG] Successfully opened file with id: %d\n", file_d);
  int line_length = read_line(file_d, &line);
  while (line_length != -1) {
    print_line(line, line_length);
    free(line);
    line_length = read_line(file_d, &line);
  }
  close(file_d);
  return 0;
}

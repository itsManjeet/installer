#ifndef TEMP_HH
#define TEMP_HH

#include <unistd.h>

#include <ctime>
#include <filesystem>
#include <iostream>
#include <string>
static std::string randomString(int const len) {
  static const char alphanum[] =
      "0123456789"
      "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
      "abcdefghijklmnopqrstuvwxyz";
  std::string temp_s;

  for (int i = 0; i < len; i++) {
    temp_s += alphanum[rand() % (sizeof(alphanum) - 1)];
  }

  return temp_s;
}

static std::string tempdir(std::string prefix) {
  auto dir = prefix + "-" + randomString(10);
  std::filesystem::create_directories(dir);
  return dir;
}

#endif
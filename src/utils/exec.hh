#ifndef EXEC_HH
#define EXEC_HH

#include <string>
#include <tuple>

class Exec {
 public:
  static std::tuple<int, std::string> output(const char* cmd) {
    std::array<char, 128> buffer;
    std::string result;

    auto pipe = popen(cmd, "r");  // get rid of shared_ptr

    if (!pipe) throw std::runtime_error("popen() failed!");

    while (!feof(pipe)) {
      if (fgets(buffer.data(), 128, pipe) != nullptr) result += buffer.data();
    }

    auto rc = pclose(pipe);

    result = result.substr(0, result.length() - 1);
    return {rc, result};
  }
};

#endif
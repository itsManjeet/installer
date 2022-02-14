#ifndef HUMANIZE_HH
#define HUMANIZE_HH

#include <string>

static inline std::string humanize(size_t bytes) {
  std::string r;
  if (bytes <= 0)
    r = "0 Bytes";
  else if (bytes >= 1073741824)
    r = std::to_string(bytes / 1073741824) + " GBytes";
  else if (bytes >= 1048576)
    r = std::to_string(bytes / 1048576) + " MBytes";
  else if (bytes >= 1024)
    r = std::to_string(bytes / 1024) + " KBytes";
  return r;
}

#endif
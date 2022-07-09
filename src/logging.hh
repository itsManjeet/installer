#ifndef INSTALLER_LOGGING_HH
#define INSTALLER_LOGGING_HH

#include <time.h>

#include <fstream>
#include <string>

inline std::string getCurrentDateTime(std::string s = "now") {
  time_t now = time(0);
  struct tm tstruct;
  char buf[80];
  tstruct = *localtime(&now);
  if (s == "now")
    strftime(buf, sizeof(buf), "%Y-%m-%d %X", &tstruct);
  else if (s == "date")
    strftime(buf, sizeof(buf), "%Y-%m-%d", &tstruct);
  return std::string(buf);
};

extern std::ofstream logfile;

#define LOG logfile << "[DEBUG " << getCurrentDateTime() << "]: "
#define ERROR logfile << "[ERROR " << getCurrentDateTime() << "]: "

#endif
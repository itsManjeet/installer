#include <fstream>
#include <iostream>

#include "config.h"
#include "interface/application.hh"

std::ofstream logfile;

int main(int argc, char** argv) {
#ifdef DEBUG
  Glib::setenv("GSETTINGS_SCHEMA_DIR", DATA_DIR, true);
#endif
  auto app = Application::create();

  logfile = std::ofstream(LOGFILE);
  return app->run(argc, argv);
}
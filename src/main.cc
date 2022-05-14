#include <fstream>
#include <iostream>

#include "config.h"
#include "interface/application.hh"

int main(int argc, char** argv) {
#ifdef DEBUG
  Glib::setenv("GSETTINGS_SCHEMA_DIR", DATA_DIR, true);
#endif
  auto app = Application::create();
  return app->run(argc, argv);
}
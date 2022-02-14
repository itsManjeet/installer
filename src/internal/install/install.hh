#ifndef INSTALL_HH
#define INSTALL_HH
#include <memory>
#include <string>
#include <vector>

#include "../../locale.hh"
#include "../../worker/worker.hh"
#include "../data/data.hh"

class Install {
 public:
  static std::vector<std::shared_ptr<Process>> list(InstallData* data);
};

#endif